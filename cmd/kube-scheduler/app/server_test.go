/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	componentbaseconfig "k8s.io/component-base/config"
	"k8s.io/kube-scheduler/config/v1beta3"
	"k8s.io/kubernetes/cmd/kube-scheduler/app/options"
	"k8s.io/kubernetes/pkg/scheduler/apis/config"
	"k8s.io/kubernetes/pkg/scheduler/apis/config/testing/defaults"
)

func TestSetup(t *testing.T) {
	// temp dir
	tmpDir, err := ioutil.TempDir("", "scheduler-options")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// https server
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`ok`))
	}))
	defer server.Close()

	configKubeconfig := filepath.Join(tmpDir, "config.kubeconfig")
	if err := ioutil.WriteFile(configKubeconfig, []byte(fmt.Sprintf(`
apiVersion: v1
kind: Config
clusters:
- cluster:
    insecure-skip-tls-verify: true
    server: %s
  name: default
contexts:
- context:
    cluster: default
    user: default
  name: default
current-context: default
users:
- name: default
  user:
    username: config
`, server.URL)), os.FileMode(0600)); err != nil {
		t.Fatal(err)
	}

	// plugin config
	pluginConfigFile := filepath.Join(tmpDir, "plugin.yaml")
	if err := ioutil.WriteFile(pluginConfigFile, []byte(fmt.Sprintf(`
apiVersion: kubescheduler.config.k8s.io/v1beta2
kind: KubeSchedulerConfiguration
clientConnection:
  kubeconfig: "%s"
profiles:
- plugins:
    preFilter:
      enabled:
      - name: NodeResourcesFit
      - name: NodePorts
      disabled:
      - name: "*"
    filter:
      enabled:
      - name: NodeResourcesFit
      - name: NodePorts
      disabled:
      - name: "*"
    preScore:
      enabled:
      - name: InterPodAffinity
      - name: TaintToleration
      disabled:
      - name: "*"
    score:
      enabled:
      - name: InterPodAffinity
      - name: TaintToleration
      disabled:
      - name: "*"
`, configKubeconfig)), os.FileMode(0600)); err != nil {
		t.Fatal(err)
	}

	// multiple profiles config
	multiProfilesConfig := filepath.Join(tmpDir, "multi-profiles.yaml")
	if err := ioutil.WriteFile(multiProfilesConfig, []byte(fmt.Sprintf(`
apiVersion: kubescheduler.config.k8s.io/v1beta2
kind: KubeSchedulerConfiguration
clientConnection:
  kubeconfig: "%s"
profiles:
- schedulerName: "profile-default-plugins"
- schedulerName: "profile-disable-all-filter-and-score-plugins"
  plugins:
    preFilter:
      disabled:
      - name: "*"
    filter:
      disabled:
      - name: "*"
    postFilter:
      disabled:
      - name: "*"
    preScore:
      disabled:
      - name: "*"
    score:
      disabled:
      - name: "*"
`, configKubeconfig)), os.FileMode(0600)); err != nil {
		t.Fatal(err)
	}

	// empty leader-election config
	emptyLeaderElectionConfig := filepath.Join(tmpDir, "empty-leader-election-config.yaml")
	if err := ioutil.WriteFile(emptyLeaderElectionConfig, []byte(fmt.Sprintf(`
apiVersion: kubescheduler.config.k8s.io/v1beta3
kind: KubeSchedulerConfiguration
clientConnection:
  kubeconfig: "%s"
`, configKubeconfig)), os.FileMode(0600)); err != nil {
		t.Fatal(err)
	}

	// leader-election config
	leaderElectionConfig := filepath.Join(tmpDir, "leader-election-config.yaml")
	if err := ioutil.WriteFile(leaderElectionConfig, []byte(fmt.Sprintf(`
apiVersion: kubescheduler.config.k8s.io/v1beta3
kind: KubeSchedulerConfiguration
clientConnection:
  kubeconfig: "%s"
leaderElection:
  leaseDuration: 1h
`, configKubeconfig)), os.FileMode(0600)); err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		name               string
		flags              []string
		wantPlugins        map[string]*config.Plugins
		wantLeaderElection *componentbaseconfig.LeaderElectionConfiguration
	}{
		{
			name: "default config",
			flags: []string{
				"--kubeconfig", configKubeconfig,
			},
			wantPlugins: map[string]*config.Plugins{
				"default-scheduler": defaults.PluginsV1beta3,
			},
		},
		{
			name: "component configuration",
			flags: []string{
				"--config", pluginConfigFile,
				"--kubeconfig", configKubeconfig,
			},
			wantPlugins: map[string]*config.Plugins{
				"default-scheduler": {
					Bind: config.PluginSet{Enabled: []config.Plugin{{Name: "DefaultBinder"}}},
					Filter: config.PluginSet{
						Enabled: []config.Plugin{
							{Name: "NodeResourcesFit"},
							{Name: "NodePorts"},
						},
					},
					PreFilter: config.PluginSet{
						Enabled: []config.Plugin{
							{Name: "NodeResourcesFit"},
							{Name: "NodePorts"},
						},
					},
					PostFilter: config.PluginSet{Enabled: []config.Plugin{{Name: "DefaultPreemption"}}},
					PreScore: config.PluginSet{
						Enabled: []config.Plugin{
							{Name: "InterPodAffinity"},
							{Name: "TaintToleration"},
						},
					},
					QueueSort: config.PluginSet{Enabled: []config.Plugin{{Name: "PrioritySort"}}},
					Score: config.PluginSet{
						Enabled: []config.Plugin{
							{Name: "InterPodAffinity", Weight: 1},
							{Name: "TaintToleration", Weight: 1},
						},
					},
					Reserve: config.PluginSet{Enabled: []config.Plugin{{Name: "VolumeBinding"}}},
					PreBind: config.PluginSet{Enabled: []config.Plugin{{Name: "VolumeBinding"}}},
				},
			},
		},
		{
			name: "leader election CLI args, along with --config arg",
			flags: []string{
				"--leader-elect=false",
				"--leader-elect-lease-duration=2h", // CLI args are favored over the fields in ComponentConfig
				"--lock-object-namespace=default",  // deprecated CLI arg will be ignored if --config is specified
				"--config", emptyLeaderElectionConfig,
			},
			wantLeaderElection: &componentbaseconfig.LeaderElectionConfiguration{
				LeaderElect:       false,                                    // from CLI args
				LeaseDuration:     metav1.Duration{Duration: 2 * time.Hour}, // from CLI args
				RenewDeadline:     metav1.Duration{Duration: 10 * time.Second},
				RetryPeriod:       metav1.Duration{Duration: 2 * time.Second},
				ResourceLock:      "leases",
				ResourceName:      v1beta3.SchedulerDefaultLockObjectName,
				ResourceNamespace: v1beta3.SchedulerDefaultLockObjectNamespace,
			},
		},
		{
			name: "leader election CLI args, without --config arg",
			flags: []string{
				"--leader-elect=false",
				"--leader-elect-lease-duration=2h",
				"--lock-object-namespace=default", // deprecated CLI arg is honored if --config is not specified
				"--kubeconfig", configKubeconfig,
			},
			wantLeaderElection: &componentbaseconfig.LeaderElectionConfiguration{
				LeaderElect:       false,                                    // from CLI args
				LeaseDuration:     metav1.Duration{Duration: 2 * time.Hour}, // from CLI args
				RenewDeadline:     metav1.Duration{Duration: 10 * time.Second},
				RetryPeriod:       metav1.Duration{Duration: 2 * time.Second},
				ResourceLock:      "leases",
				ResourceName:      v1beta3.SchedulerDefaultLockObjectName,
				ResourceNamespace: "default", // from deprecated CLI args
			},
		},
		{
			name: "leader election settings specified by ComponentConfig only",
			flags: []string{
				"--config", leaderElectionConfig,
			},
			wantLeaderElection: &componentbaseconfig.LeaderElectionConfiguration{
				LeaderElect:       true,
				LeaseDuration:     metav1.Duration{Duration: 1 * time.Hour}, // from CC
				RenewDeadline:     metav1.Duration{Duration: 10 * time.Second},
				RetryPeriod:       metav1.Duration{Duration: 2 * time.Second},
				ResourceLock:      "leases",
				ResourceName:      v1beta3.SchedulerDefaultLockObjectName,
				ResourceNamespace: v1beta3.SchedulerDefaultLockObjectNamespace,
			},
		},
		{
			name: "leader election settings specified by CLI args and ComponentConfig",
			flags: []string{
				"--leader-elect=true",
				"--leader-elect-renew-deadline=5s",
				"--leader-elect-retry-period=1s",
				"--config", leaderElectionConfig,
			},
			wantLeaderElection: &componentbaseconfig.LeaderElectionConfiguration{
				LeaderElect:       true,
				LeaseDuration:     metav1.Duration{Duration: 1 * time.Hour},   // from CC
				RenewDeadline:     metav1.Duration{Duration: 5 * time.Second}, // from CLI args
				RetryPeriod:       metav1.Duration{Duration: 1 * time.Second}, // from CLI args
				ResourceLock:      "leases",
				ResourceName:      v1beta3.SchedulerDefaultLockObjectName,
				ResourceNamespace: v1beta3.SchedulerDefaultLockObjectNamespace,
			},
		},
	}

	makeListener := func(t *testing.T) net.Listener {
		t.Helper()
		l, err := net.Listen("tcp", ":0")
		if err != nil {
			t.Fatal(err)
		}
		return l
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			fs := pflag.NewFlagSet("test", pflag.PanicOnError)
			opts, err := options.NewOptions()
			if err != nil {
				t.Fatal(err)
			}
			// use listeners instead of static ports so parallel test runs don't conflict
			opts.SecureServing.Listener = makeListener(t)
			defer opts.SecureServing.Listener.Close()

			nfs := opts.Flags()
			for _, f := range nfs.FlagSets {
				fs.AddFlagSet(f)
			}
			if err := fs.Parse(tc.flags); err != nil {
				t.Fatal(err)
			}

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			_, sched, err := Setup(ctx, opts)
			if err != nil {
				t.Fatal(err)
			}

			if tc.wantPlugins != nil {
				gotPlugins := make(map[string]*config.Plugins)
				for n, p := range sched.Profiles {
					gotPlugins[n] = p.ListPlugins()
				}

				if diff := cmp.Diff(tc.wantPlugins, gotPlugins); diff != "" {
					t.Errorf("Unexpected plugins diff (-want, +got): %s", diff)
				}
			}

			if tc.wantLeaderElection != nil {
				gotLeaderElection := opts.ComponentConfig.LeaderElection
				if diff := cmp.Diff(*tc.wantLeaderElection, gotLeaderElection); diff != "" {
					t.Errorf("Unexpected leaderElection diff (-want, +got): %s", diff)
				}
			}
		})
	}
}
