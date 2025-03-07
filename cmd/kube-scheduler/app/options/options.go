/*
Copyright 2018 The Kubernetes Authors.

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

package options

import (
	"fmt"
	"net"
	"os"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	apiserveroptions "k8s.io/apiserver/pkg/server/options"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/events"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/tools/record"
	cliflag "k8s.io/component-base/cli/flag"
	componentbaseconfig "k8s.io/component-base/config"
	"k8s.io/component-base/config/options"
	"k8s.io/component-base/logs"
	"k8s.io/component-base/metrics"
	schedulerappconfig "k8s.io/kubernetes/cmd/kube-scheduler/app/config"
	"k8s.io/kubernetes/pkg/scheduler"
	kubeschedulerconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"
	"k8s.io/kubernetes/pkg/scheduler/apis/config/latest"
	"k8s.io/kubernetes/pkg/scheduler/apis/config/validation"
	netutils "k8s.io/utils/net"
)

// Options has all the params needed to run a Scheduler
type Options struct {
	// The default values.
	ComponentConfig kubeschedulerconfig.KubeSchedulerConfiguration

	SecureServing  *apiserveroptions.SecureServingOptionsWithLoopback
	Authentication *apiserveroptions.DelegatingAuthenticationOptions
	Authorization  *apiserveroptions.DelegatingAuthorizationOptions
	Metrics        *metrics.Options
	Logs           *logs.Options
	Deprecated     *DeprecatedOptions

	// ConfigFile is the location of the scheduler server's configuration file.
	ConfigFile string

	// WriteConfigTo is the path where the default configuration will be written.
	WriteConfigTo string

	Master string

	// flags hold the parsed CLI flags.
	flags *cliflag.NamedFlagSets
}

// NewOptions returns default scheduler app options.
func NewOptions() (*Options, error) {
	cfg, err := latest.Default()
	if err != nil {
		return nil, err
	}

	o := &Options{
		ComponentConfig: *cfg,
		SecureServing:   apiserveroptions.NewSecureServingOptions().WithLoopback(),
		Authentication:  apiserveroptions.NewDelegatingAuthenticationOptions(),
		Authorization:   apiserveroptions.NewDelegatingAuthorizationOptions(),
		Deprecated:      &DeprecatedOptions{},
		Metrics:         metrics.NewOptions(),
		Logs:            logs.NewOptions(),
	}

	o.Authentication.TolerateInClusterLookupFailure = true
	o.Authentication.RemoteKubeConfigFileOptional = true
	o.Authorization.RemoteKubeConfigFileOptional = true

	// Set the PairName but leave certificate directory blank to generate in-memory by default
	o.SecureServing.ServerCert.CertDirectory = ""
	o.SecureServing.ServerCert.PairName = "kube-scheduler"
	o.SecureServing.BindPort = kubeschedulerconfig.DefaultKubeSchedulerPort

	return o, nil
}

// Complete obtains the CLI args related with leaderelection, and override the values in `cfg`.
// Then the `cfg` object is injected into the `options` object.
func (o *Options) Complete(cfg *kubeschedulerconfig.KubeSchedulerConfiguration) {
	// Obtain CLI args related with leaderelection. Set them to `cfg` if specified in command line.
	leaderelection := o.Flags().FlagSet("leader election")
	if leaderelection.Changed("leader-elect") {
		cfg.LeaderElection.LeaderElect = o.ComponentConfig.LeaderElection.LeaderElect
	}
	if leaderelection.Changed("leader-elect-lease-duration") {
		cfg.LeaderElection.LeaseDuration = o.ComponentConfig.LeaderElection.LeaseDuration
	}
	if leaderelection.Changed("leader-elect-renew-deadline") {
		cfg.LeaderElection.RenewDeadline = o.ComponentConfig.LeaderElection.RenewDeadline
	}
	if leaderelection.Changed("leader-elect-retry-period") {
		cfg.LeaderElection.RetryPeriod = o.ComponentConfig.LeaderElection.RetryPeriod
	}
	if leaderelection.Changed("leader-elect-resource-lock") {
		cfg.LeaderElection.ResourceLock = o.ComponentConfig.LeaderElection.ResourceLock
	}
	if leaderelection.Changed("leader-elect-resource-name") {
		cfg.LeaderElection.ResourceName = o.ComponentConfig.LeaderElection.ResourceName
	}
	if leaderelection.Changed("leader-elect-resource-namespace") {
		cfg.LeaderElection.ResourceNamespace = o.ComponentConfig.LeaderElection.ResourceNamespace
	}

	o.ComponentConfig = *cfg
}

// Flags returns flags for a specific scheduler by section name
func (o *Options) Flags() *cliflag.NamedFlagSets {
	if o.flags != nil {
		return o.flags
	}

	nfs := cliflag.NamedFlagSets{}
	fs := nfs.FlagSet("misc")
	fs.StringVar(&o.ConfigFile, "config", o.ConfigFile, "The path to the configuration file.")
	fs.StringVar(&o.WriteConfigTo, "write-config-to", o.WriteConfigTo, "If set, write the configuration values to this file and exit.")
	fs.StringVar(&o.Master, "master", o.Master, "The address of the Kubernetes API server (overrides any value in kubeconfig)")

	// In some tests, o.SecureServing can be nil.
	if o.SecureServing != nil {
		o.SecureServing.AddFlags(nfs.FlagSet("secure serving"))
	}
	o.Authentication.AddFlags(nfs.FlagSet("authentication"))
	o.Authorization.AddFlags(nfs.FlagSet("authorization"))
	o.Deprecated.AddFlags(nfs.FlagSet("deprecated"), &o.ComponentConfig)

	options.BindLeaderElectionFlags(&o.ComponentConfig.LeaderElection, nfs.FlagSet("leader election"))
	utilfeature.DefaultMutableFeatureGate.AddFlag(nfs.FlagSet("feature gate"))
	o.Metrics.AddFlags(nfs.FlagSet("metrics"))
	o.Logs.AddFlags(nfs.FlagSet("logs"))

	o.flags = &nfs
	return o.flags
}

// ApplyTo applies the scheduler options to the given scheduler app configuration.
func (o *Options) ApplyTo(c *schedulerappconfig.Config) error {
	if len(o.ConfigFile) == 0 {
		c.ComponentConfig = o.ComponentConfig
	} else {
		cfg, err := loadConfigFromFile(o.ConfigFile)
		if err != nil {
			return err
		}
		// Honor the CLI args before assigning `cfg` to `c.ComponentConfig`.
		o.Complete(cfg)

		if err := validation.ValidateKubeSchedulerConfiguration(cfg); err != nil {
			return err
		}

		c.ComponentConfig = *cfg
	}

	if err := o.SecureServing.ApplyTo(&c.SecureServing, &c.LoopbackClientConfig); err != nil {
		return err
	}
	if o.SecureServing != nil && (o.SecureServing.BindPort != 0 || o.SecureServing.Listener != nil) {
		if err := o.Authentication.ApplyTo(&c.Authentication, c.SecureServing, nil); err != nil {
			return err
		}
		if err := o.Authorization.ApplyTo(&c.Authorization); err != nil {
			return err
		}
	}
	o.Metrics.Apply()
	o.Logs.Apply()
	return nil
}

// Validate validates all the required options.
func (o *Options) Validate() []error {
	var errs []error

	if err := validation.ValidateKubeSchedulerConfiguration(&o.ComponentConfig); err != nil {
		errs = append(errs, err.Errors()...)
	}
	errs = append(errs, o.SecureServing.Validate()...)
	errs = append(errs, o.Authentication.Validate()...)
	errs = append(errs, o.Authorization.Validate()...)
	errs = append(errs, o.Deprecated.Validate()...)
	errs = append(errs, o.Metrics.Validate()...)
	errs = append(errs, o.Logs.Validate()...)

	return errs
}

// Config return a scheduler config object
func (o *Options) Config() (*schedulerappconfig.Config, error) {
	if o.SecureServing != nil {
		if err := o.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{netutils.ParseIPSloppy("127.0.0.1")}); err != nil {
			return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
		}
	}

	c := &schedulerappconfig.Config{}
	if err := o.ApplyTo(c); err != nil {
		return nil, err
	}

	// Prepare kube config.
	kubeConfig, err := createKubeConfig(c.ComponentConfig.ClientConnection, o.Master)
	if err != nil {
		return nil, err
	}

	// Prepare kube clients.
	client, eventClient, err := createClients(kubeConfig)
	if err != nil {
		return nil, err
	}

	c.EventBroadcaster = events.NewEventBroadcasterAdapter(eventClient)

	// Set up leader election if enabled.
	var leaderElectionConfig *leaderelection.LeaderElectionConfig
	if c.ComponentConfig.LeaderElection.LeaderElect {
		// Use the scheduler name in the first profile to record leader election.
		schedulerName := corev1.DefaultSchedulerName
		if len(c.ComponentConfig.Profiles) != 0 {
			schedulerName = c.ComponentConfig.Profiles[0].SchedulerName
		}
		coreRecorder := c.EventBroadcaster.DeprecatedNewLegacyRecorder(schedulerName)
		leaderElectionConfig, err = makeLeaderElectionConfig(c.ComponentConfig.LeaderElection, kubeConfig, coreRecorder)
		if err != nil {
			return nil, err
		}
	}

	c.Client = client
	c.KubeConfig = kubeConfig
	c.InformerFactory = scheduler.NewInformerFactory(client, 0)
	dynClient := dynamic.NewForConfigOrDie(kubeConfig)
	c.DynInformerFactory = dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynClient, 0, corev1.NamespaceAll, nil)
	c.LeaderElection = leaderElectionConfig

	return c, nil
}

// makeLeaderElectionConfig builds a leader election configuration. It will
// create a new resource lock associated with the configuration.
func makeLeaderElectionConfig(config componentbaseconfig.LeaderElectionConfiguration, kubeConfig *restclient.Config, recorder record.EventRecorder) (*leaderelection.LeaderElectionConfig, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("unable to get hostname: %v", err)
	}
	// add a uniquifier so that two processes on the same host don't accidentally both become active
	id := hostname + "_" + string(uuid.NewUUID())

	rl, err := resourcelock.NewFromKubeconfig(config.ResourceLock,
		config.ResourceNamespace,
		config.ResourceName,
		resourcelock.ResourceLockConfig{
			Identity:      id,
			EventRecorder: recorder,
		},
		kubeConfig,
		config.RenewDeadline.Duration)
	if err != nil {
		return nil, fmt.Errorf("couldn't create resource lock: %v", err)
	}

	return &leaderelection.LeaderElectionConfig{
		Lock:            rl,
		LeaseDuration:   config.LeaseDuration.Duration,
		RenewDeadline:   config.RenewDeadline.Duration,
		RetryPeriod:     config.RetryPeriod.Duration,
		WatchDog:        leaderelection.NewLeaderHealthzAdaptor(time.Second * 20),
		Name:            "kube-scheduler",
		ReleaseOnCancel: true,
	}, nil
}

// createKubeConfig creates a kubeConfig from the given config and masterOverride.
// TODO remove masterOverride when CLI flags are removed.
func createKubeConfig(config componentbaseconfig.ClientConnectionConfiguration, masterOverride string) (*restclient.Config, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags(masterOverride, config.Kubeconfig)
	if err != nil {
		return nil, err
	}

	kubeConfig.DisableCompression = true
	kubeConfig.AcceptContentTypes = config.AcceptContentTypes
	kubeConfig.ContentType = config.ContentType
	kubeConfig.QPS = config.QPS
	kubeConfig.Burst = int(config.Burst)

	return kubeConfig, nil
}

// createClients creates a kube client and an event client from the given kubeConfig
func createClients(kubeConfig *restclient.Config) (clientset.Interface, clientset.Interface, error) {
	client, err := clientset.NewForConfig(restclient.AddUserAgent(kubeConfig, "scheduler"))
	if err != nil {
		return nil, nil, err
	}

	eventClient, err := clientset.NewForConfig(kubeConfig)
	if err != nil {
		return nil, nil, err
	}

	return client, eventClient, nil
}
