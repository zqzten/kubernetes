<!-- BEGIN MUNGE: GENERATED_TOC -->

- [v1.22.2](#v1222)
  - [Downloads for v1.22.2](#downloads-for-v1222)
    - [Source Code](#source-code)
    - [Client Binaries](#client-binaries)
    - [Server Binaries](#server-binaries)
    - [Node Binaries](#node-binaries)
  - [Changelog since v1.22.1](#changelog-since-v1221)
  - [Important Security Information](#important-security-information)
    - [CVE-2021-25741: Symlink Exchange Can Allow Host Filesystem Access](#cve-2021-25741-symlink-exchange-can-allow-host-filesystem-access)
  - [Changes by Kind](#changes-by-kind)
    - [Feature](#feature)
    - [Bug or Regression](#bug-or-regression)
  - [Dependencies](#dependencies)
    - [Added](#added)
    - [Changed](#changed)
    - [Removed](#removed)
- [v1.22.1](#v1221)
  - [Downloads for v1.22.1](#downloads-for-v1221)
    - [Source Code](#source-code-1)
    - [Client Binaries](#client-binaries-1)
    - [Server Binaries](#server-binaries-1)
    - [Node Binaries](#node-binaries-1)
  - [Changelog since v1.22.0](#changelog-since-v1220)
  - [Changes by Kind](#changes-by-kind-1)
    - [Feature](#feature-1)
    - [Bug or Regression](#bug-or-regression-1)
    - [Other (Cleanup or Flake)](#other-cleanup-or-flake)
  - [Dependencies](#dependencies-1)
    - [Added](#added-1)
    - [Changed](#changed-1)
    - [Removed](#removed-1)
- [v1.22.0](#v1220)
  - [Downloads for v1.22.0](#downloads-for-v1220)
    - [Source Code](#source-code-2)
    - [Client Binaries](#client-binaries-2)
    - [Server Binaries](#server-binaries-2)
    - [Node Binaries](#node-binaries-2)
  - [Changelog since v1.21.0](#changelog-since-v1210)
  - [What's New (Major Themes)](#whats-new-major-themes)
    - [Removal of several beta Kubernetes APIs](#removal-of-several-beta-kubernetes-apis)
    - [Kubernetes release cadence change](#kubernetes-release-cadence-change)
    - [External credential providers](#external-credential-providers)
    - [Server-side Apply graduates to GA](#server-side-apply-graduates-to-ga)
    - [Cluster Storage Interface graduations](#cluster-storage-interface-graduations)
    - [SIG Windows development tools](#sig-windows-development-tools)
    - [Deploy a more secure control plane with kubeadm](#deploy-a-more-secure-control-plane-with-kubeadm)
    - [etcd moves to version 3.5.0](#etcd-moves-to-version-350)
    - [Kubernetes Node system swap support](#kubernetes-node-system-swap-support)
    - [Cluster-wide seccomp defaults](#cluster-wide-seccomp-defaults)
    - [Quality of Service for memory resources](#quality-of-service-for-memory-resources)
    - [API changes and improvements for ephemeral containers](#api-changes-and-improvements-for-ephemeral-containers)
  - [Known Issues](#known-issues)
    - [CPU and Memory manager are not working correctly for Guaranteed Pods with multiple containers](#cpu-and-memory-manager-are-not-working-correctly-for-guaranteed-pods-with-multiple-containers)
    - [<code>CSIMigrationvSphere</code> feature gate has not migrated to new CRD APIs](#-feature-gate-has-not-migrated-to-new-crd-apis)
  - [Urgent Upgrade Notes](#urgent-upgrade-notes)
    - [(No, really, you MUST read this before you upgrade)](#no-really-you-must-read-this-before-you-upgrade)
  - [Changes by Kind](#changes-by-kind-2)
    - [Deprecation](#deprecation)
    - [API Change](#api-change)
    - [Feature](#feature-2)
    - [Documentation](#documentation)
    - [Failing Test](#failing-test)
    - [Bug or Regression](#bug-or-regression-2)
    - [Other (Cleanup or Flake)](#other-cleanup-or-flake-1)
  - [Dependencies](#dependencies-2)
    - [Added](#added-2)
    - [Changed](#changed-2)
    - [Removed](#removed-2)
- [v1.22.0-rc.0](#v1220-rc0)
  - [Downloads for v1.22.0-rc.0](#downloads-for-v1220-rc0)
    - [Source Code](#source-code-3)
    - [Client Binaries](#client-binaries-3)
    - [Server Binaries](#server-binaries-3)
    - [Node Binaries](#node-binaries-3)
  - [Changelog since v1.22.0-beta.2](#changelog-since-v1220-beta2)
  - [Changes by Kind](#changes-by-kind-3)
    - [API Change](#api-change-1)
    - [Bug or Regression](#bug-or-regression-3)
  - [Dependencies](#dependencies-3)
    - [Added](#added-3)
    - [Changed](#changed-3)
    - [Removed](#removed-3)
- [v1.22.0-beta.2](#v1220-beta2)
  - [Downloads for v1.22.0-beta.2](#downloads-for-v1220-beta2)
    - [Source Code](#source-code-4)
    - [Client Binaries](#client-binaries-4)
    - [Server Binaries](#server-binaries-4)
    - [Node Binaries](#node-binaries-4)
  - [Changelog since v1.22.0-beta.1](#changelog-since-v1220-beta1)
  - [Changes by Kind](#changes-by-kind-4)
    - [API Change](#api-change-2)
    - [Feature](#feature-3)
    - [Bug or Regression](#bug-or-regression-4)
  - [Dependencies](#dependencies-4)
    - [Added](#added-4)
    - [Changed](#changed-4)
    - [Removed](#removed-4)
- [v1.22.0-beta.1](#v1220-beta1)
  - [Downloads for v1.22.0-beta.1](#downloads-for-v1220-beta1)
    - [Source Code](#source-code-5)
    - [Client Binaries](#client-binaries-5)
    - [Server Binaries](#server-binaries-5)
    - [Node Binaries](#node-binaries-5)
  - [Changelog since v1.22.0-beta.0](#changelog-since-v1220-beta0)
  - [Urgent Upgrade Notes](#urgent-upgrade-notes-1)
    - [(No, really, you MUST read this before you upgrade)](#no-really-you-must-read-this-before-you-upgrade-1)
  - [Changes by Kind](#changes-by-kind-5)
    - [Deprecation](#deprecation-1)
    - [API Change](#api-change-3)
    - [Feature](#feature-4)
    - [Documentation](#documentation-1)
    - [Bug or Regression](#bug-or-regression-5)
    - [Other (Cleanup or Flake)](#other-cleanup-or-flake-2)
  - [Dependencies](#dependencies-5)
    - [Added](#added-5)
    - [Changed](#changed-5)
    - [Removed](#removed-5)
- [v1.22.0-beta.0](#v1220-beta0)
  - [Downloads for v1.22.0-beta.0](#downloads-for-v1220-beta0)
    - [Source Code](#source-code-6)
    - [Client Binaries](#client-binaries-6)
    - [Server Binaries](#server-binaries-6)
    - [Node Binaries](#node-binaries-6)
  - [Changelog since v1.22.0-alpha.3](#changelog-since-v1220-alpha3)
  - [Urgent Upgrade Notes](#urgent-upgrade-notes-2)
    - [(No, really, you MUST read this before you upgrade)](#no-really-you-must-read-this-before-you-upgrade-2)
  - [Changes by Kind](#changes-by-kind-6)
    - [Deprecation](#deprecation-2)
    - [API Change](#api-change-4)
    - [Feature](#feature-5)
    - [Bug or Regression](#bug-or-regression-6)
    - [Other (Cleanup or Flake)](#other-cleanup-or-flake-3)
  - [Dependencies](#dependencies-6)
    - [Added](#added-6)
    - [Changed](#changed-6)
    - [Removed](#removed-6)
- [v1.22.0-alpha.3](#v1220-alpha3)
  - [Downloads for v1.22.0-alpha.3](#downloads-for-v1220-alpha3)
    - [Source Code](#source-code-7)
    - [Client Binaries](#client-binaries-7)
    - [Server Binaries](#server-binaries-7)
    - [Node Binaries](#node-binaries-7)
  - [Changelog since v1.22.0-alpha.2](#changelog-since-v1220-alpha2)
  - [Urgent Upgrade Notes](#urgent-upgrade-notes-3)
    - [(No, really, you MUST read this before you upgrade)](#no-really-you-must-read-this-before-you-upgrade-3)
  - [Changes by Kind](#changes-by-kind-7)
    - [Deprecation](#deprecation-3)
    - [API Change](#api-change-5)
    - [Feature](#feature-6)
    - [Bug or Regression](#bug-or-regression-7)
    - [Other (Cleanup or Flake)](#other-cleanup-or-flake-4)
  - [Dependencies](#dependencies-7)
    - [Added](#added-7)
    - [Changed](#changed-7)
    - [Removed](#removed-7)
- [v1.22.0-alpha.2](#v1220-alpha2)
  - [Downloads for v1.22.0-alpha.2](#downloads-for-v1220-alpha2)
    - [Source Code](#source-code-8)
    - [Client Binaries](#client-binaries-8)
    - [Server Binaries](#server-binaries-8)
    - [Node Binaries](#node-binaries-8)
  - [Changelog since v1.22.0-alpha.1](#changelog-since-v1220-alpha1)
  - [Urgent Upgrade Notes](#urgent-upgrade-notes-4)
    - [(No, really, you MUST read this before you upgrade)](#no-really-you-must-read-this-before-you-upgrade-4)
  - [Changes by Kind](#changes-by-kind-8)
    - [Deprecation](#deprecation-4)
    - [API Change](#api-change-6)
    - [Feature](#feature-7)
    - [Failing Test](#failing-test-1)
    - [Bug or Regression](#bug-or-regression-8)
    - [Other (Cleanup or Flake)](#other-cleanup-or-flake-5)
  - [Dependencies](#dependencies-8)
    - [Added](#added-8)
    - [Changed](#changed-8)
    - [Removed](#removed-8)
- [v1.22.0-alpha.1](#v1220-alpha1)
  - [Downloads for v1.22.0-alpha.1](#downloads-for-v1220-alpha1)
    - [Source Code](#source-code-9)
    - [Client binaries](#client-binaries-9)
    - [Server binaries](#server-binaries-9)
    - [Node binaries](#node-binaries-9)
  - [Changelog since v1.21.0](#changelog-since-v1210-1)
  - [Urgent Upgrade Notes](#urgent-upgrade-notes-5)
    - [(No, really, you MUST read this before you upgrade)](#no-really-you-must-read-this-before-you-upgrade-5)
  - [Changes by Kind](#changes-by-kind-9)
    - [Deprecation](#deprecation-5)
    - [API Change](#api-change-7)
    - [Feature](#feature-8)
    - [Failing Test](#failing-test-2)
    - [Bug or Regression](#bug-or-regression-9)
    - [Other (Cleanup or Flake)](#other-cleanup-or-flake-6)
  - [Dependencies](#dependencies-9)
    - [Added](#added-9)
    - [Changed](#changed-9)
    - [Removed](#removed-9)

<!-- END MUNGE: GENERATED_TOC -->

# v1.22.2


## Downloads for v1.22.2

### Source Code

filename | sha512 hash
-------- | -----------
[kubernetes.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes.tar.gz) | 5a567e5133a04da82a656b072151fa2b8a6b680db85a4faf69f6e727d2ec9889fd2bae9c5be8562074352c67eace863c894b734650e07e7fce91dcf5986f9357
[kubernetes-src.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-src.tar.gz) | eebf175f20ebee1ac03a067b66a680e67bc2f9a209c699ff9954d53a6f37e35ae832397ead051c993b14aa12d627b16e5d89518cf2974a1a26a8ed28277458e0

### Client Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-client-darwin-amd64.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-client-darwin-amd64.tar.gz) | 264ccda40e275da89d66284002ab7818880bb37056053d0d2e8d535cd835f2f56fc43e995c76ea9041e25a59f8f6b4cff8710130c285fcf383c5417926e371d1
[kubernetes-client-darwin-arm64.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-client-darwin-arm64.tar.gz) | 04aa25b204d289e5abd704055a05adee331dac9d075791737b64033142d1c43238773ab2357b308fdde5ff1cc390a686ed5efa3be424b4324dae361418d898d9
[kubernetes-client-linux-386.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-client-linux-386.tar.gz) | c3b8f0b54fc9fcf64f87af93ca74d3e2eb87fbfdc346e09ed36406827fab29ac63e0feefe6f47612d26caacb9ed2d56189ae29c996d98c35b4984a87a51c1507
[kubernetes-client-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-client-linux-amd64.tar.gz) | 09694e377b5104c47d291626cdb9c199519119b0ae27c1d9ed61b6dd544f462032026188a298f533494ad04ec6e0366ed3e3eac89122f658c2efee433b25090f
[kubernetes-client-linux-arm.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-client-linux-arm.tar.gz) | 39b9a64d6f58ec42455985c9feb785ea6f9354612f16bc9e4c2e1ae6a74bf5e252d609907ad6c9b7ba4ba84763cc2382531115ed1c5265eb91d457ec6cb8a31d
[kubernetes-client-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-client-linux-arm64.tar.gz) | 832a8b023f3b4f54ba21ac2e9e74aa686522441619c53f73b3cf34efe9ee4df447447526215b68f759564f3be9929436247b23250d29202eef87c7c1f58bb26b
[kubernetes-client-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-client-linux-ppc64le.tar.gz) | b078e73d0dabb6cb4409aa2a01f4139cda82228573f6702a2890b84417750bfcf01c868d2914f7231361a0782e40082bdef9460cd59e4d04a210b7278495d495
[kubernetes-client-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-client-linux-s390x.tar.gz) | c7f4776df0613e0c76e9e7e42d3020d2391daba0f50530df246cd65767dc394346d3acc0be91b1fc3eabcb930e73afcf4ad1068110d859b4fdd2f6a0cca4b69b
[kubernetes-client-windows-386.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-client-windows-386.tar.gz) | 27f770302a66b922b04fe810687fd1b36cfe9a96e12d67ed8e3873d2b27dc3bb2df55ba5e305c64184d38f20c8ea55627419f37cd8c1822b7716637fe09aff0b
[kubernetes-client-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-client-windows-amd64.tar.gz) | dc504970bb4d7555e3d41a0b4762d1d94fbae4118e4399bd49bea8766b41fbcf6e25276efb7dada941d14447e51fc00372535fbf2f6045ea4eca17a8dcd2978d

### Server Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-server-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-server-linux-amd64.tar.gz) | 4187b6d62c57412820394c051c14fe00be3c02ecae0f3a9d2dfd62a5e74cc89f2c94e094d3c23e9e5ec130cc2d0f68d7d2f7d6f484cbfbae4939493f44332d76
[kubernetes-server-linux-arm.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-server-linux-arm.tar.gz) | 17103d3e88c7542a742fcd758ee75cbab2d54e9c1f5eb9b998b2a278a1401c819fae2234959971a4f9fe8e6f4a491b0bf7cccc3e98efc31a14704282ff74e835
[kubernetes-server-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-server-linux-arm64.tar.gz) | c97be69e10f19862f61e898baff43475a1ef59336d9ced5f6688feebece476b6cf78a2f51ee6a3940405ce0d6162e85465515e533fff1716af5bc8266fd94f50
[kubernetes-server-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-server-linux-ppc64le.tar.gz) | 1b8e65e010c9885ba5dd2b7b9cbcd4caaca2c9edc5af77b92276d47f3e7b5c2af49eeb1436300682596d0fca15793ada21160a8c58bdc74ef81dead8d391c246
[kubernetes-server-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-server-linux-s390x.tar.gz) | a571c3ce7d295628dd7a979765594ba2e6dfe96f494002edd0df4e704dc002eef497655cb31a98991c2bbdd09bd432797e0c778e9bbd8168ded2867b012285d5

### Node Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-node-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-node-linux-amd64.tar.gz) | 40e97ee26566af810a2881441572b798c8feb010ac5727507ec379e897dc3cfefc9e6c06f09de8ef0b52d941dbb929003590006cc8cf76403f6c64e231e6759e
[kubernetes-node-linux-arm.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-node-linux-arm.tar.gz) | d852803b6cc264f83328e6ecb7d509d083cc4ed9a84dafc2d08c5fc1ef38fc2bf395d9bf5e4c87b11699c047be07a957177157396f61d7e957c7cc08881d3fb8
[kubernetes-node-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-node-linux-arm64.tar.gz) | 9b35d16a1d2760a017564f6d09ec17b0aa2ed2b4cd62065d9aa0fe760d080fa8d64feb4ab7fb54eda1b8a1a760b4ee58ac36ad1c134b4adefe3e45872209816e
[kubernetes-node-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-node-linux-ppc64le.tar.gz) | dfc779984b3c2859b103ffa446131befe455531b1eedc8dd796871d027d17abf2d33c1ea46f0cdb0c1a349418980e8986dcd10a8fae779966922fe543b070bd0
[kubernetes-node-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-node-linux-s390x.tar.gz) | 508b47667233f467418d24a4b55af0d6360ab82744568b207465d035035ff8106dae0ef992cb4af60884be8c27b7370f90c1db161444a9adc90a61814133b361
[kubernetes-node-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.2/kubernetes-node-windows-amd64.tar.gz) | 1d9472eef1be3c640d86e3299c96f55de975311ef3589473024fbcc9b9ff4d1111af993cbafe526ff45ecf7ff01146a5c65224b23f4f2ed289137bcfb6374ab9

## Changelog since v1.22.1

## Important Security Information

This release contains changes that address the following vulnerabilities:

### CVE-2021-25741: Symlink Exchange Can Allow Host Filesystem Access

A security issue was discovered in Kubernetes where a user may be able to
create a container with subpath volume mounts to access files &
directories outside of the volume, including on the host filesystem.

**Affected Versions**:
  - kubelet v1.22.0 - v1.22.1
  - kubelet v1.21.0 - v1.21.4
  - kubelet v1.20.0 - v1.20.10
  - kubelet <= v1.19.14

**Fixed Versions**:
  - kubelet v1.22.2
  - kubelet v1.21.5
  - kubelet v1.20.11
  - kubelet v1.19.15

This vulnerability was reported by Fabricio Voznika and Mark Wolters of Google.

**CVSS Rating:** High (8.8) [CVSS:3.0/AV:N/AC:L/PR:L/UI:N/S:U/C:H/I:H/A:H](https://www.first.org/cvss/calculator/3.0#CVSS:3.0/AV:N/AC:L/PR:L/UI:N/S:U/C:H/I:H/A:H)

## Changes by Kind

### Feature

- Kubernetes is now built with Golang 1.16.8 ([#104905](https://github.com/kubernetes/kubernetes/pull/104905), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]

### Bug or Regression

- Fix Job tracking with finalizers for more than 500 pods, ensuring all finalizers are removed before counting the Pod. ([#104876](https://github.com/kubernetes/kubernetes/pull/104876), [@alculquicondor](https://github.com/alculquicondor)) [SIG Apps]
- Fix: skip case sensitivity when checking Azure NSG rules
  fix: ensure InstanceShutdownByProviderID return false for creating Azure VMs ([#104446](https://github.com/kubernetes/kubernetes/pull/104446), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fixed occasional pod cgroup freeze when using cgroup v1 and systemd driver. ([#104529](https://github.com/kubernetes/kubernetes/pull/104529), [@kolyshkin](https://github.com/kolyshkin)) [SIG Node]
- Fixes a regression that could cause panics in LRU caches in controller-manager, kubelet, kube-apiserver, or client-go EventSourceObjectSpamFilter ([#104469](https://github.com/kubernetes/kubernetes/pull/104469), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Storage]
- When using `kubectl replace` (or the equivalent API call) on a Service, the caller no longer needs to do a read-modify-write cycle to fetch the allocated values for `.spec.clusterIP` and `.spec.ports[].nodePort`.  Instead the API server will automatically carry these forward from the original object when the new object does not specify them. ([#104672](https://github.com/kubernetes/kubernetes/pull/104672), [@thockin](https://github.com/thockin)) [SIG Network]

## Dependencies

### Added
_Nothing has changed._

### Changed
- github.com/opencontainers/runc: [v1.0.1 → v1.0.2](https://github.com/opencontainers/runc/compare/v1.0.1...v1.0.2)
- k8s.io/utils: 4b05e18 → bdf08cb

### Removed
_Nothing has changed._



# v1.22.1


## Downloads for v1.22.1

### Source Code

filename | sha512 hash
-------- | -----------
[kubernetes.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes.tar.gz) | 4ba0d63665c5959cec560264a260bb1341d09a28f6651e9814982af1ec47aff144e9ad7e6c4273867c864cdf3d12f1e7cbb4bfa44301cf5b9e81f98a345acdfb
[kubernetes-src.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-src.tar.gz) | eac27ada62bf719f8d31655fadbb74c2c13793c57fe866e768c86dfbbe90c8ca1c2c07d093f1874c198e49a125968de5951ec8f07d90f8b236a1a1bcb1640b03

### Client Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-client-darwin-amd64.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-client-darwin-amd64.tar.gz) | 41dcfff40498335c92d872eb6aa6215d9d84908f79bdf61777d52243a446fffb114313fb586ec6a58908684eaf86fae6b2fbaa70a2a3f06f78470a20ca4d9727
[kubernetes-client-darwin-arm64.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-client-darwin-arm64.tar.gz) | 810b6db31fa9c277044b1b0e39a72430f3ff728722f1fd9cc85143835147314b3fd7a6374d963eeff96cf572d2a0b034e3cab4bc0294a219832d58c40f0c1302
[kubernetes-client-linux-386.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-client-linux-386.tar.gz) | d215c81c9a5274ced217fe09d2111f4073ed225c24c3daa71f91a4f75e38817dea504ff9b9558dcee7bd1a0c26d936c2e2bbe66e30736879fd0fba4aa0438f03
[kubernetes-client-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-client-linux-amd64.tar.gz) | 064bd1eaf468c9b4a00e31bec3f9c80850c52cd1e06edfd86f307236acbdba7c89dbe663cbfefdc5ead72f1cb93ba9d21d828558508f0b29f8c814d9085846ca
[kubernetes-client-linux-arm.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-client-linux-arm.tar.gz) | 07754877a5f486efb97ddb4a5d0edc8389a59bc7dba2927c73829f1aa77e36593dec6d0ba5386582221c9cf33497ebc68181d8cef4144185cf2817a5d85064c8
[kubernetes-client-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-client-linux-arm64.tar.gz) | 0b2265f27baad42425f6d815182261542efba2e5defccdb92668a50ff0f49c831f3af4ea6c38f351004ab1eafc270871e443e673188f37463ee3012da20cfc27
[kubernetes-client-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-client-linux-ppc64le.tar.gz) | a689eaef1c0788eb734fd5e70899f4cd0bc4ad6dd9f56827e7d55cab4acb2b1b0035226e69e1230fe74ac45accbbafbcc616c00fabe8dc55b9d8f94519cde253
[kubernetes-client-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-client-linux-s390x.tar.gz) | a76300dbd915b232826d022e1fd5b8a336e7de2089d897aed594381ad571396619609498687eba75f2dcb7fe2fb2bfe50fe6a199eda53af8513d9b3b2967b7ec
[kubernetes-client-windows-386.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-client-windows-386.tar.gz) | 46c62a9dd3d5753bf8111a32b19f41babf7c06a10f8931fede339da1605d45f020850bc568cdc9b78300d0d7a5deb4c8602092f8b6667b73bea18eadf758a14c
[kubernetes-client-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-client-windows-amd64.tar.gz) | 88e30aee1103ab6b4f3c2bd5570b4f739cbc01b4755bfaee915dcb717ae4d8b0b99457e2bb44cc60f47b77822ce6119fbe77bdfbaaf4afa4220b78b40b1c9b2c

### Server Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-server-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-server-linux-amd64.tar.gz) | 66c596577d19c1a609b554e1e61b0bc2c8ba691b2101f481dfc616799e37347edc62956b7684370dec5b16388d48a665ba78c942de4aaf5df1f78cf0667df545
[kubernetes-server-linux-arm.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-server-linux-arm.tar.gz) | b2f6d2ef240a2cf829823ef599aad5590f531c6873ace7a682d9f74954f90b9c48d3f5f33460bd050bee87736e37c331a99ed599146cebd5ade36459d41a0842
[kubernetes-server-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-server-linux-arm64.tar.gz) | bf7d8978b72409f02c1b7ad5f44a908f31287ee715966e47726c3f9f85c349a14f457cd84160cfe78d4e39233a119afd152528ed79db5eef6558a25098825438
[kubernetes-server-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-server-linux-ppc64le.tar.gz) | 9546df4488665af35a7289e77f8b194f12e5939199411fb8f4b2f15db5e669eef9dfba950315205d39f4d89c135fbccd37e8ad9c554a482bb09977234d7d7185
[kubernetes-server-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-server-linux-s390x.tar.gz) | e1eb0868e613e27ea7f71df5875c5d0b107a7fc38525151dffd22826c7a47111ece7d3948badb28b1a1c9d8cabca70af54cdc1501e5aadf5c842bf09bd26f1c2

### Node Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-node-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-node-linux-amd64.tar.gz) | 53b44fbf58308c18b478ced5629a82d01c67dd6ae5be9559a2d65edae339983bff4d143b1bde062f5d8202ec8241e45da3b087af1fb0aea35afb1812a26450d1
[kubernetes-node-linux-arm.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-node-linux-arm.tar.gz) | c3855f27d4001dde9f89f41dc15ddab66a9bd5fa2cc058b65fc47e6c487ac12ef3e2d94ed72e9560a88696a86ef893a548eef3ef78d4528b2903f0f3d487e9df
[kubernetes-node-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-node-linux-arm64.tar.gz) | eb6a5cb5270a662a79302472a347a5486b3e0689c0c8ea0ceb62e83f72c043dc3b15ee3d5ec40790a1754748950208d8b5818b3f1d64e44a881720a78f50cc7e
[kubernetes-node-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-node-linux-ppc64le.tar.gz) | 45038ed6ee110d04ea7fb18e8f3c5f33aff2cc291a169d7a2f5f37ae1780b534ff5cd2895b7c2a9fb675a8490c9da7518e09f4ca9e2af4d57238557e1021d255
[kubernetes-node-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-node-linux-s390x.tar.gz) | 14a286eb32ae450e1afb9bfcf4fd84be259613b1be4d81a4433000c507427f137ff79a6da9f457ade4b668a01c1de76488b134a895d36673531cde372f3ee884
[kubernetes-node-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.1/kubernetes-node-windows-amd64.tar.gz) | a38016acc743eb87d4803d2d06c24fc3315a4db6eacbd3f2060f28355ca96f1b3294b30483b74ecea18fe220a9d1319509f26fe02615fbaa7f1ecd11f76b1f6b

## Changelog since v1.22.0

## Changes by Kind

### Feature

- Kubernetes is now built with Golang 1.16.7 ([#104200](https://github.com/kubernetes/kubernetes/pull/104200), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]

### Bug or Regression

- Fix kube-apiserver metric reporting for the deprecated watch path of /api/<version>/watch/... ([#104188](https://github.com/kubernetes/kubernetes/pull/104188), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery and Instrumentation]
- Kube-proxy: delete stale conntrack UDP entries for loadbalancer ingress IP. ([#104009](https://github.com/kubernetes/kubernetes/pull/104009), [@aojea](https://github.com/aojea)) [SIG Network]
- Pass additional flags to subpath mount to avoid flakes in certain conditions ([#104346](https://github.com/kubernetes/kubernetes/pull/104346), [@mauriciopoppe](https://github.com/mauriciopoppe)) [SIG Storage]

### Other (Cleanup or Flake)

- Kube-apiserver: sets an upper-bound on the lifetime of idle keep-alive connections and time to read the headers of incoming requests ([#103958](https://github.com/kubernetes/kubernetes/pull/103958), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Node]

## Dependencies

### Added
_Nothing has changed._

### Changed
_Nothing has changed._

### Removed
_Nothing has changed._



# v1.22.0

[Documentation](https://docs.k8s.io)

## Downloads for v1.22.0

### Source Code

filename | sha512 hash
-------- | -----------
[kubernetes.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes.tar.gz) | `d1145ec29a8581a4c94a83cefa3658a73bfc7d8e2624d31e735d53551718c9212e477673f74cfa4e430a8367a47bba65e2573162711613e60db54563dc912f00`
[kubernetes-src.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-src.tar.gz) | `94d4430765ae8463c2509492050ea8925e7a5f1f3e58fb76e2b87602aa89d9a321110fb0c9f9003a8640c53adec12c82200cc5c126eb6e7a6a5716ecae67305b`

### Client Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-client-darwin-amd64.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-client-darwin-amd64.tar.gz) | `83022159507b761b806551062293c4a88fc513041b758d28fef26f38911b49f6a9581e600f23329563eaf5c62965177a298f7f7919a5bee7170dd34b16348aa6`
[kubernetes-client-darwin-arm64.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-client-darwin-arm64.tar.gz) | `c3e253a20e2b91a3f83c7f742e7064aa8ec9c89f3d8a37ef593d4ccad88877844074c4bfbfa5d19408e065b7a83eecbac170745561d8a5f6b10637c64d9b1c41`
[kubernetes-client-linux-386.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-client-linux-386.tar.gz) | `8c9ac2f45cb475a6c5191a67d27d3bd6e287f71391eb7afcf8fb195322dcaba052bcabd36999961cf07ab38aea68e1bdda49df1bc0c4c4e4e98055bbecc82b58`
[kubernetes-client-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-client-linux-amd64.tar.gz) | `15707fc968fdb8e3d5cd80bb23fbb4e579e8642d9724ad3b179c6d0f5b7dfe425f1c878a6120101137d4cba2ad2bbd19d677dd84b5bc6f6023c82f6a06e4153e`
[kubernetes-client-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-client-linux-arm.tar.gz) | `a710cff509469d3a35ecd254346b093a7c358f7118b92e52df600bfbdd1230e7340c9a09de6c9fac30996b8c46e3d7b1e2391a2b39a38d43668bd25ddb1782d8`
[kubernetes-client-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-client-linux-arm64.tar.gz) | `d8088e12154654cd5da7e0225b54f0d052132774d37e14d42e31e87a8a4bc34ec1fe18e7574ed5e4fc0b08591979bf788827ac9a0a59de26d0b7ac629bfe1cd8`
[kubernetes-client-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-client-linux-ppc64le.tar.gz) | `3739185084afcb725ad7612c05fed7c3655fe57bfa06c825736b43568d3672cd1075e7e463edcfa4c3cb429f3c1d349a6127c7deaf1c0542f03ca2b8b6180411`
[kubernetes-client-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-client-linux-s390x.tar.gz) | `be8cf34b3361b479622b5173ec99d5885ced493cbba9e42a1b0587a062e54076f6ea6543f08e9fb55b5a9f41ba64967a3237b1621e75f24684c924d748a5e42e`
[kubernetes-client-windows-386.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-client-windows-386.tar.gz) | `e8d199b2b124f6fdc9849127791279325814c5fdaeaf54443e878a3616b08ae3a3bf3181432ef64d946d96448e2a6b48f0f0dd5be3902cfbe9d14f34b255da40`
[kubernetes-client-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-client-windows-amd64.tar.gz) | `03b988b4d184f30e9956736e6fc60841f3b46f5d2017b8ed3b0f790b6f85380c85009050e65b331f76112404550f54db77e42659212623948fcfa969fe25026b`

### Server Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-server-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-server-linux-amd64.tar.gz) | `d54435de50214faabc49e3659625a689623508128ca9a4f97b4f2c54b40bc9e14dd17e1971c06c90aa74fc335d0038a7ac4b7b90882edb0944af99354d6c9762`
[kubernetes-server-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-server-linux-arm.tar.gz) | `b7349715480fc0245e83a364f1d413831e8677d9c27569302addd4746a83f0c94430a30dddc3394dec31cd542130e4e6a09a5fe90f1dde42a0ed9f11cbf831ac`
[kubernetes-server-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-server-linux-arm64.tar.gz) | `7517349b33b1d49514276be23c3c52ba946bc3f33b98b6c9aefc8dba1cc034364ccca609f4dbb1f8880696c15e7204a9a584de7abab1184a5ad55ff662bf4f00`
[kubernetes-server-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-server-linux-ppc64le.tar.gz) | `a007a714128a08e7cfa42152d63b3cb99da9e7198a0908d8baeccc56e52a4a6ce50d7a442c020ce0651067193b01cf82230b7bbfc8dd99b7ee9958eaea387645`
[kubernetes-server-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-server-linux-s390x.tar.gz) | `f642555c23121ab5cbeb74f98de054138cce2a929475364794b1a60a1a64197b0d1b28ad5e78279765d389d84ca0d57759f6cdb790c63d6afc80f6cdf2751b8c`

### Node Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-node-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-node-linux-amd64.tar.gz) | `aa990405a1c6bd6737a8ff89fd536ba28ad62dec7de2e44ae223f4fcb42d6a9ffdfb324144def946b777ac7ba6fac085a49a7977cb79289a3256cced783bf215`
[kubernetes-node-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-node-linux-arm.tar.gz) | `d99a535163c67a4e4fe5c2eca828255bc071f9f9aae0b0c71132980df772844fa493fa4c2ac2f422e76192f2318057301eb451a681eae14ba037632dd47352bf`
[kubernetes-node-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-node-linux-arm64.tar.gz) | `6cc6de072827944314b4162918ac2eead9900347669e4507f1bb4ddb119ca87c4bc2c15a7dc6305a8b6cb29dea80ea764a145fc62f5ab87fec4a7c3ef7daed66`
[kubernetes-node-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-node-linux-ppc64le.tar.gz) | `dbe9e45152fce851bdfdac9443f3eca4f8f98e4d20dce3dab044ee70c87365ab44158d5fde4f82963816bcb4d4d20fa8d56bf6f6b3e378dc5da5faeec6e8fd55`
[kubernetes-node-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-node-linux-s390x.tar.gz) | `9d8674ab7590a4e2c3151d709e36c0756af38754d5835611b4f964cc3eaa46064fdaa46acd04f10a3a82cb5599574947de9812d9a7bd05e99d3b7c0dcd1acc5e`
[kubernetes-node-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0/kubernetes-node-windows-amd64.tar.gz) | `9cc73fb1d3f9ec926fd09bc3904d62ec79da4a3c4fb9a5c4c784bc1f08c650711c21fb30874b05db4bd354e4d04b0153296180d89a53c04d9241dd6a1384510d`

## Changelog since v1.21.0

## What's New (Major Themes)

### Removal of several beta Kubernetes APIs

A number of APIs are no longer serving specific Beta versions in favour of the GA version of those APIs. All existing objects can be interacted with via general availability APIs. This removal includes beta versions of `ValidatingWebhookConfiguration`, `MutatingWebhookConfiguration`, `CustomResourceDefinition`, `APIService`, `TokenReview`, `SubjectAccessReview`, `CertificateSigningRequest`, `Lease`, `Ingress`, and `IngressClass` APIs. For the full list check out [Deprecated API Migration Guide](https://kubernetes.io/docs/reference/using-api/deprecation-guide/#v1-22) and the blog post [Kubernetes API and Feature Removals In 1.22: Here’s What You Need To Know](https://blog.k8s.io/2021/07/14/upcoming-changes-in-kubernetes-1-22/).

### Kubernetes release cadence change

We all have to adapt to change in our lives, and especially so in the past year. The Kubernetes release team was also affected from the COVID-19 pandemic and has listened to its user base regarding the number of releases in a calendar year. From April 23, 2021 it was made official that Kubernetes release cadence has reduced from 4 releases per year to 3 releases per year.

You can read more in the official blog post [Kubernetes Release Cadence Change: Here’s What You Need To Know](https://blog.k8s.io/2021/07/20/new-kubernetes-release-cadence/).

### External credential providers

Kubernetes client [credential plugins](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins) have been in beta since 1.11, a few eons ago. With the release of Kubernetes 1.22, this feature set graduates to stable. The GA feature set includes improved support for plugins that provide interactive login flows.
This release also contains a number of bug fixes to the feature set. Aspiring plugin authors can look at [sample-exec-plugin](https://github.com/ankeesler/sample-exec-plugin) as a way to get started.

Related to this topic, the in-tree Azure and GCP authentication plugins have been [deprecated](https://github.com/kubernetes/kubernetes/pull/102181) in favor of out-of-tree implementations.

### Server-side Apply graduates to GA

Server-side Apply is a new object merge algorithm, as well as tracking of field ownership, running on the Kubernetes API server. Server-side Apply helps users and controllers manage their resources via declarative configurations. It allows them to create and/or modify their objects declaratively, simply by sending their fully specified intent. After being in beta for a couple releases, [Server-side Apply](https://kubernetes.io/docs/reference/using-api/server-side-apply/) is now generally available.

### Cluster Storage Interface graduations

CSI support for Windows nodes moves to GA in the 1.22 release. In Kubernetes v1.22, Windows privileged containers are only an alpha feature. To allow using CSI storage on Windows nodes, [CSIProxy](https://github.com/kubernetes-csi/csi-proxy) enables CSI node plugins to be deployed as unprivileged pods, using the proxy to perform privileged storage operations on the node.

Another feature moving to GA in v1.22 is CSI Service Account Token support. This feature allows CSI drivers to use pods' [bound service account tokens](https://kubernetes.io/docs/reference/access-authn-authz/service-accounts-admin/#bound-service-account-token-volume) instead of a more privileged identity. It also provides control over to re-publishing these volumes, so that short-lived tokens can be refreshed.

### SIG Windows development tools

To grow the developer community, SIG Windows  released multiple [tools](https://github.com/kubernetes-sigs/sig-windows-dev-tools/). The new tools support multiple CNI providers (Antrea, Calico), can run on multiple platforms (any vagrant compatible provider, such as Hyper-V, VirtualBox, or vSphere).
There is also a new way to run bleeding edge Windows features from scratch by compiling the windows kubelet and kube-proxy, then using them along with daily builds of other Kubernetes components.

### Deploy a more secure control plane with kubeadm

A new alpha feature allows running the kubeadm control plane components as non-root users. This is a long requested security measure in kubeadm. To try it you must enable the kubeadm-specific `RootlessControlPlane` feature gate. When you deploy a cluster using this alpha feature, your control plane runs with lower privileges.

A new [v1beta3 configuration API](https://github.com/kubernetes/kubeadm/issues/1796). It [iterates over v1beta2](https://pkg.go.dev/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta3) by adding some long requested features and deprecating some existing ones. The `V1beta3` is now the preferred API version; the `v1beta2` API also remains available and is not yet deprecated.

### etcd moves to version 3.5.0

Kubernetes' default backend storage, etcd, has a new release 3.5.0 and the community embraced it. The new release comes with improvements to the Security, performance, monitoring and developer experience. There are numerous bug fixes to lease objects causing memory leaks, and compact operation causing deadlocks and more. A couple of new features are also introduced like the migration to structured logging and build in log rotation. The release comes with a detailed future roadmap to implement a solution to traffic overload. A full and detailed list of changes can be read in the [3.5.0 release announcement](https://etcd.io/blog/2021/announcing-etcd-3.5/).

### Kubernetes Node system swap support

Every system administrator or Kubernetes user has been in the same boat regarding setting up and using Kubernetes: disable swap space. With the release of Kubernetes 1.22, *alpha* support is available to run nodes with swap memory. This change lets administrators opt in to configuring swap on Linux nodes, treating a portion of block storage as additional virtual memory.

### Cluster-wide seccomp defaults

A new alpha feature gate `SeccompDefault` has been added to the kubelet, together with a corresponding command line flag `--seccomp-default` and kubelet configuration. If both are enabled, then the kubelet's behavior changes for pods that don't explicitly set a seccomp profile.
With cluster-wide seccomp defaults, the kubelet uses the `RuntimeDefault` seccomp profile by default, rather than than `Unconfined`. This allows enhancing the default cluster wide workload security of the Kubernetes deployment. Security administrators will now sleep better knowing there is some security by default for the workloads.

To learn more about the feature, please refer to the official [seccomp tutorial](https://kubernetes.io/docs/tutorials/clusters/seccomp/#enable-the-use-of-runtimedefault-as-the-default-seccomp-profile-for-all-workloads).

### Quality of Service for memory resources

Originally, Kubernetes used the v1 cgroups API. With that design, the QoS class for a pod only applied to CPU resources (such as `cpu_shares`). The Kubernetes cgroup manager uses `memory.limit_in_bytes` in v1 cgroups to limit the memory capacity for a container, and uses `oom_scores` to recommend an order for killing container processes if an out-of-memory event occurs. This implementation has shortcomings: for `Guaranteed` pods, memory can not be fully reserved, and the page cache is at risk of being recycled. For `Burstable` pods, overcommitting memory (setting `request` less than `limit` ) could increase the risk of a container being killed when the Linux kernel detects an out of memory condition.

As an alpha feature, Kubernetes v1.22 can use the cgroups v2 API to control memory allocation and isolation. This feature is designed to improve workload and node availability when there is contention for memory resources.

### API changes and improvements for ephemeral containers

The API used to create [Ephemeral Containers](https://kubernetes.io/docs/concepts/workloads/pods/ephemeral-containers/) changed in 1.22. The Ephemeral Containers feature is alpha and disabled by default, and the new API does not work with clients that attempt to use the old API.

For stable features, the `kubectl` tool follows the Kubernetes [version skew policy](https://kubernetes.io/releases/version-skew-policy/);
however, kubectl v1.21 and older do not support the new API for ephemeral containers.
Users who create ephemeral containers using `kubectl debug` should note that kubectl version 1.22 will attempt to fall back to the old API; older versions of kubectl will not work with cluster versions of 1.22 or later. Please update kubectl to 1.22 if you wish to use `kubectl debug` with a mix of cluster versions.

## Known Issues

### CPU and Memory manager are not working correctly for Guaranteed Pods with multiple containers

A regression bug was found where guaranteed Pods with multiple containers do not work properly with set allocations for CPU, Memory, and Device manager. [The fix will be availability in coming releases](https://github.com/kubernetes/kubernetes/pull/103979).

### `CSIMigrationvSphere` feature gate has not migrated to new CRD APIs

If CSIMigrationvSphere feature gate is enabled, user should not upgrade to Kubernetes v1.22. vSphere CSI Driver does not support Kubernetes v1.22 yet because it uses v1beta1 CRD APIs. Support for v1.22 will be added at a later release. Check the following document for supported Kubernetes releases for a given [vSphere CSI Driver version](https://vsphere-csi-driver.sigs.k8s.io/compatiblity_matrix.html#compatibility-matrix-for-vsphere-csi-driver).

## Urgent Upgrade Notes 

### (No, really, you MUST read this before you upgrade)

- Audit log files are now created with a mode of 0600. Existing file permissions will not be changed. If you need the audit file to be readable by a non-root user, you can pre-create the file with the desired permissions. ([#95387](https://github.com/kubernetes/kubernetes/pull/95387), [@JAORMX](https://github.com/JAORMX)) [SIG API Machinery and Auth]
 - CSI migration of AWS EBS volumes requires AWS EBS CSI driver ver. 1.0 that supports `allowAutoIOPSPerGBIncrease` parameter in StorageClass. ([#101082](https://github.com/kubernetes/kubernetes/pull/101082), [@jsafrane](https://github.com/jsafrane))
 - Conformance image is now built with Distroless. Users running Conformance testing should rely on container entrypoint instead of manual invocation to `/run_e2e.sh` or `/gorunner`, as they are now deprecated and will be removed in 1.25 release. Invoking `ginkgo` and `e2e.test` are still supported through overriding entrypoint (docker) or defining container `spec.command` (kubernetes). ([#99178](https://github.com/kubernetes/kubernetes/pull/99178), [@wilsonehusin](https://github.com/wilsonehusin))
 - Default `StreamingProxyRedirects` to disabled. If there is a >= 2 version skew between master and nodes, and the old nodes were enabling `--redirect-container-streaming`, this will break them. In this case, the `StreamingProxyRedirects` can still be manually enabled. ([#101647](https://github.com/kubernetes/kubernetes/pull/101647), [@pacoxu](https://github.com/pacoxu))
 - Intree volume plugin scaleIO support has been completely removed from Kubernetes. ([#101685](https://github.com/kubernetes/kubernetes/pull/101685), [@Jiawei0227](https://github.com/Jiawei0227))
 - Kubeadm: remove the automatic detection and matching of cgroup drivers for Docker. For new clusters if you have not configured the cgroup driver explicitly you might get a failure in the `kubelet` on driver mismatch (kubeadm clusters should be using the `systemd` driver). Also remove the `IsDockerSystemdCheck` preflight check (warning) that checks if the Docker cgroup driver is set to `systemd`. Ideally such detection / coordination should be on the side of CRI implementers and the kubelet (tracked [here](https://github.com/kubernetes/kubernetes/issues/99808)). Please see the [page](https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/configure-cgroup-driver/) on how to configure cgroup drivers with kubeadm manually ([#99647](https://github.com/kubernetes/kubernetes/pull/99647), [@neolit123](https://github.com/neolit123))
 - Kubeadm: the flag `--cri-socket` is no longer allowed in a mixture with the flag `--config`. Please use the kubeadm configuration for setting the CRI socket for a node using `{Init|Join}Configuration.nodeRegistration.criSocket`. ([#101600](https://github.com/kubernetes/kubernetes/pull/101600), [@KofClubs](https://github.com/KofClubs))
 - Newly provisioned PVs by Azure disk will no longer have the beta `FailureDomain` label. Azure disk volume plugin will start to have GA topology label instead. ([#101534](https://github.com/kubernetes/kubernetes/pull/101534), [@kassarl](https://github.com/kassarl))
 - Scheduler's CycleState now embeds internal read/write locking inside its `Read()` and `Write()` functions. Meanwhile, `Lock()` and `Unlock()` function are removed. Scheduler plugin developers are now required to remove `CycleState#Lock()` and `CycleState#Unlock()`. Just simply use `Read()` and `Write()` as they're natively thread-safe now. ([#101542](https://github.com/kubernetes/kubernetes/pull/101542), [@Huang-Wei](https://github.com/Huang-Wei))
 - The `CSIMigrationVSphereComplete` feature flag is removed. `InTreePluginvSphereUnregister` will be the way moving forward. ([#101272](https://github.com/kubernetes/kubernetes/pull/101272), [@Jiawei0227](https://github.com/Jiawei0227))
 - The flag `--experimental-patches` is now deprecated and will be removed in a future release. You can migrate to using the new flag `--patches`. Add a new field `{Init|Join}Configuration.patches.directory` that can be used for the same purpose. For `init` and `join` it is now recommended that you migrate to configure patches via `{Init|Join}Configuration.patches.directory`. For the time being, these flags can be mixed with `--config`, but that might change in the future. On a command line, the last *patches flag takes precedence over previous flags and the value in config. `kubeadm upgrade --patches` will continue to be the only available option, since `upgrade` does not support a configuration file yet. ([#103063](https://github.com/kubernetes/kubernetes/pull/103063), [@neolit123](https://github.com/neolit123))
 
## Changes by Kind

### Deprecation

- Controller-manager: the following flags have no effect and would be removed in v1.24:
  - `--port`
  - `--address`
  The insecure port flags `--port` may only be set to 0 now.
  
  In addtion, please be careful that:
  - controller-manager MUST start with `--authorization-kubeconfig` and `--authentication-kubeconfig` correctly set to get authentication/authorization working.
  - liveness/readiness probes to controller-manager MUST use HTTPS now, and the default port has been changed to 10257.
  - Applications that fetch metrics from controller-manager should use a dedicated service account which is allowed to access nonResourceURLs `/metrics`. ([#96216](https://github.com/kubernetes/kubernetes/pull/96216), [@knight42](https://github.com/knight42)) [SIG API Machinery, Cloud Provider, Instrumentation and Testing]
- Deprecate `--record` flag in `kubectl`. The `--record` flag is being replaced with the [mechanism](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/859-kubectl-headers) which annotates HTTP requests with kubectl command details. ([#102873](https://github.com/kubernetes/kubernetes/pull/102873), [@soltysh](https://github.com/soltysh))
- E2e.test: removed the  `--viper-config` flag. If you were previously using this to pass flags to `e2e.test` via a file, you will need to pass them directly on the command line, e.g. `e2e.test --e2e-output-dir`. ([#102598](https://github.com/kubernetes/kubernetes/pull/102598), [@dims](https://github.com/dims))
- For `kubeadm`: remove the ClusterStatus API from v1beta3 and its management in the kube-system/kubeadm-config ConfigMap. This method of keeping track of what API endpoints exists in the cluster was replaced (in a prior release) by a method to annotate the etcd Pods that `kubeadm` creates in "stacked etcd" clusters. The following CLI sub-phases are deprecated and are now a NO-OP: for `kubeadm join`: "control-plane-join/update-status", for `kubeadm reset`: "update-cluster-status". Unless you are using these phases explicitly, you should not be affected. ([#101915](https://github.com/kubernetes/kubernetes/pull/101915), [@neolit123](https://github.com/neolit123))
- Kubeadm: remove the deprecated `--csr-only` and `--csr-dir` flags from `kubeadm init phase certs`. Deprecate the same flags under `kubeadm certs renew`. In both the cases the command `kubeadm certs generate-csr` should be used instead. ([#102108](https://github.com/kubernetes/kubernetes/pull/102108), [@neolit123](https://github.com/neolit123))
- Kubeadm: Remove the deprecated command `kubeadm alpha kubeconfig`. Please use `kubeadm kubeconfig` instead. ([#101938](https://github.com/kubernetes/kubernetes/pull/101938), [@knight42](https://github.com/knight42))
- Kubeadm: Remove the deprecated hyperkube image support in `v1beta3`. This implies removal of `ClusterConfiguration.UseHyperKubeImage.` ([#101537](https://github.com/kubernetes/kubernetes/pull/101537), [@neolit123](https://github.com/neolit123))
- Kubeadm: Remove the field `ClusterConfiguration.DNS.Type` in v1beta3 since CoreDNS is the only supported DNS type. ([#101547](https://github.com/kubernetes/kubernetes/pull/101547), [@neolit123](https://github.com/neolit123))
- Kubeadm: remove the deprecated command `kubeadm config view`. A replacement for this command is `kubectl get cm -n kube-system kubeadm-config -o=jsonpath="{.data.ClusterConfiguration}"` ([#102071](https://github.com/kubernetes/kubernetes/pull/102071), [@neolit123](https://github.com/neolit123))
- Kubeadm: remove the deprecated flag '--image-pull-timeout' for 'kubeadm upgrade apply' command ([#102093](https://github.com/kubernetes/kubernetes/pull/102093), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Kubeadm: remove the deprecated flag `--insecure-port` from the kube-apiserver manifest that kubeadm manages. The flag had no effect since 1.20, since the insecure serving of the component was disabled in the same version. ([#102121](https://github.com/kubernetes/kubernetes/pull/102121), [@pacoxu](https://github.com/pacoxu))
- Kubeadm: remove the deprecated kubeadm API `v1beta1`. Introduce a new kubeadm API `v1beta3`. See [kubeadm/v1beta3](https://pkg.go.dev/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta3) for a list of changes since `v1beta2`. Note that `v1beta2` is not yet deprecated, but will be in a future release. ([#101129](https://github.com/kubernetes/kubernetes/pull/101129), [@neolit123](https://github.com/neolit123))
- Newly provisioned PVs by vSphere in-tree plugin will no longer have the beta `FailureDomain` label. vSphere volume plugin will start to have GA topology label ([#102414](https://github.com/kubernetes/kubernetes/pull/102414), [@divyenpatel](https://github.com/divyenpatel))
- Removal of the CSI `NodePublish` path by the kubelet is deprecated. This must be done by the CSI plugin according to the CSI spec. ([#101441](https://github.com/kubernetes/kubernetes/pull/101441), [@dobsonj](https://github.com/dobsonj))
- Remove support for the Service `topologyKeys` field (alpha) and the `kube-proxy` implementation of it. This field was deprecated several cycles ago. This functionality is replaced by the combination of automatic topology hints per-endpoint (alpha) and the Service `internalTrafficPolicy` field (alpha). ([#102412](https://github.com/kubernetes/kubernetes/pull/102412), [@andrewsykim](https://github.com/andrewsykim))
- The `PodUnknown` phase is now deprecated. ([#95286](https://github.com/kubernetes/kubernetes/pull/95286), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev))
- The `storageos`, `quobyte` and `flocker` storage volume plugins are deprecated and will be removed in a later release. ([#101773](https://github.com/kubernetes/kubernetes/pull/101773), [@Jiawei0227](https://github.com/Jiawei0227))
- The deprecated flag `--hard-pod-affinity-symmetric-weight` and `--scheduler-name` have been removed from `kube-scheduler`. Use `ComponentConfig` instead to configure those parameters. ([#102805](https://github.com/kubernetes/kubernetes/pull/102805), [@ahg-g](https://github.com/ahg-g))
- The feature Dynamic Kubelet Configuration is deprecated and kubelet will report warning when the flag `--dynamic-config-dir` is used. Feature gate `DynamicKubeletConfig` is disabled out of the box and needs to be explicitly enabled. ([#102966](https://github.com/kubernetes/kubernetes/pull/102966), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev)) [SIG Cloud Provider, Instrumentation and Node]
- The in-tree azure and gcp auth plugins have been deprecated.  The https://github.com/Azure/kubelogin and gcloud commands serve as out-of-tree replacements via the kubectl/client-go credential plugin mechanism. ([#102181](https://github.com/kubernetes/kubernetes/pull/102181), [@enj](https://github.com/enj)) [SIG API Machinery and Auth]
- The ingress `v1beta1` has been deprecated. ([#102030](https://github.com/kubernetes/kubernetes/pull/102030), [@aojea](https://github.com/aojea))

### API Change

- A new score extension for NodeResourcesFit plugin that merges the functionality of `NodeResourcesLeastAllocated`, `NodeResourcesMostAllocated`, `RequestedToCapacityRatio` plugins, which are marked as deprecated as of v1beta2. In v1beta1, the three plugins can still be used in v1beta1 but not at the same time with the score extension of `NodeResourcesFit`. ([#101822](https://github.com/kubernetes/kubernetes/pull/101822), [@yuzhiquan](https://github.com/yuzhiquan))
- A value of `Auto` is now a valid for the `service.kubernetes.io/topology-aware-hints` annotation. ([#100728](https://github.com/kubernetes/kubernetes/pull/100728), [@robscott](https://github.com/robscott))
- Add `DataSourceRef` alpha field to PVC spec, which allows contents other than `PVCs` and `VolumeSnapshots` to be data sources. ([#103276](https://github.com/kubernetes/kubernetes/pull/103276), [@bswartz](https://github.com/bswartz))
- Add `PersistentVolumeClaimDeletePoilcy` to StatefulSet API. ([#99378](https://github.com/kubernetes/kubernetes/pull/99378), [@mattcary](https://github.com/mattcary))
- Add a new Priority and Fairness rule that exempts all probes (`/readyz`, `/healthz`, `/livez`) to prevent restarting of healthy `kube-apiserver` instance by kubelet. ([#100678](https://github.com/kubernetes/kubernetes/pull/100678), [@tkashem](https://github.com/tkashem))
- Add alpha support for HostProcess containers on Windows ([#99576](https://github.com/kubernetes/kubernetes/pull/99576), [@marosset](https://github.com/marosset)) [SIG API Machinery, Apps, Node, Testing and Windows]
- Add distributed tracing to the `kube-apiserver`. It is can be enabled with the feature gate `APIServerTracing` ([#94942](https://github.com/kubernetes/kubernetes/pull/94942), [@dashpole](https://github.com/dashpole))
- Add three metrics to the job controller to monitor if a job works in healthy condition.
  `IndexedJob` has been promoted to Beta. ([#101292](https://github.com/kubernetes/kubernetes/pull/101292), [@AliceZhang2016](https://github.com/AliceZhang2016))
- Added field `.status.uncountedTerminatedPods` to the Job resource. This field is used by the job controller to keep track of finished pods before adding them to the Job status counters. Pods created by the job controller get the finalizer `batch.kubernetes.io/job-tracking`
  Jobs that are tracked using this mechanism get the annotation `batch.kubernetes.io/job-tracking`. This is a temporary measure. Two releases after this feature graduates to beta, the annotation won't be added to Jobs anymore. ([#98817](https://github.com/kubernetes/kubernetes/pull/98817), [@alculquicondor](https://github.com/alculquicondor))
- Added new kubelet alpha feature `SeccompDefault`. This feature enables falling back to
  the `RuntimeDefault` (former `runtime/default`) seccomp profile if nothing else is specified
  in the pod/container `SecurityContext` or the pod annotation level. To use the feature, enable
  the feature gate as well as set the kubelet configuration option `SeccompDefault`
  (`--seccomp-default`) to `true`. ([#101943](https://github.com/kubernetes/kubernetes/pull/101943), [@saschagrunert](https://github.com/saschagrunert)) [SIG Node]
- Adds the `ReadWriteOncePod` access mode for `PersistentVolumes` and `PersistentVolumeClaims`. Restricts volume access to a single pod on a single node. ([#102028](https://github.com/kubernetes/kubernetes/pull/102028), [@chrishenzie](https://github.com/chrishenzie))
- Alpha swap support can now be enabled on Kubernetes nodes with the `NodeSwapEnabled` feature flag. See [KEP-2400](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/2400-node-swap/README.md#design-details) for details. ([#102823](https://github.com/kubernetes/kubernetes/pull/102823), [@ehashman](https://github.com/ehashman))
- Because of the implementation logic of `time.Format` in golang, the displayed time zone is not consistent. ([#102366](https://github.com/kubernetes/kubernetes/pull/102366), [@cndoit18](https://github.com/cndoit18))
- Corrected the documentation for escaping dollar signs in a container's env, command and args property. ([#101916](https://github.com/kubernetes/kubernetes/pull/101916), [@MartinKanters](https://github.com/MartinKanters)) [SIG Apps]
- Enable `MaxSurge` for `DaemonSet` by default. ([#101742](https://github.com/kubernetes/kubernetes/pull/101742), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla))
- Enforce the `ReadWriteOncePod` PVC access mode during scheduling ([#103082](https://github.com/kubernetes/kubernetes/pull/103082), [@chrishenzie](https://github.com/chrishenzie))
- Ephemeral containers are now allowed to configure a `securityContext` that differs from that of the Pod. Cluster administrators should ensure that security policy controllers support `EphemeralContainers` before enabling this feature in clusters. ([#99023](https://github.com/kubernetes/kubernetes/pull/99023), [@verb](https://github.com/verb))
- Exec plugin authors can override default handling of standard input via new `interactiveMode` kubeconfig field. ([#99310](https://github.com/kubernetes/kubernetes/pull/99310), [@ankeesler](https://github.com/ankeesler))
- If someone had the `ProbeTerminationGracePeriod` alpha feature enabled in 1.21, they should update/delete any workloads/pods with probe `terminationGracePeriods` < 1 before upgrading ([#103245](https://github.com/kubernetes/kubernetes/pull/103245), [@wzshiming](https://github.com/wzshiming))
- Improved parsing of label selectors ([#102188](https://github.com/kubernetes/kubernetes/pull/102188), [@alculquicondor](https://github.com/alculquicondor)) [SIG API Machinery]
- Introduce `minReadySeconds` api to the `StatefulSets`. ([#100842](https://github.com/kubernetes/kubernetes/pull/100842), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla))
- Introducing Memory quality of service support with `cgroups v2 (Alpha)`. The `MemoryQoS` feature is now in Alpha. This allows `kubelet` running with `cgroups v2` to set memory QoS at container, pod and QoS level to protect and guarantee better memory quality. This feature can be enabled through feature gate Memory QoS. ([#102970](https://github.com/kubernetes/kubernetes/pull/102970), [@borgerli](https://github.com/borgerli))
- Kube API server accepts `Impersonate-Uid` header to impersonate a user with a specific UID, in the same way that you can currently use `Impersonate-User`, `Impersonate-Group` and `Impersonate-Extra`. ([#99961](https://github.com/kubernetes/kubernetes/pull/99961), [@margocrawf](https://github.com/margocrawf))
- Kube-apiserver: `--service-account-issuer` can be specified multiple times now, to enable non-disruptive change of issuer. ([#101155](https://github.com/kubernetes/kubernetes/pull/101155), [@zshihang](https://github.com/zshihang)) [SIG API Machinery, Auth, Node and Testing]
- Kube-controller-manager: the `--horizontal-pod-autoscaler-use-rest-clients` flag and Heapster support in the horizontal pod autoscaler, deprecated since 1.12, is removed. ([#90368](https://github.com/kubernetes/kubernetes/pull/90368), [@serathius](https://github.com/serathius))
- Kube-scheduler: a plugin enabled in a v1beta2 configuration file takes precedence over the default configuration for that plugin. This simplifies enabling default plugins with custom configuration without needing to explicitly disable those default plugins. ([#99582](https://github.com/kubernetes/kubernetes/pull/99582), [@chendave](https://github.com/chendave))
- New `node-high` priority-level has been added to Suggested API Priority and ([#101151](https://github.com/kubernetes/kubernetes/pull/101151), [@mborsz](https://github.com/mborsz))
- NodeSwapEnabled feature flag was renamed to NodeSwap
  
  The flag was only available in the 1.22.0-beta.1 release, and the new flag should be used going forward. ([#103553](https://github.com/kubernetes/kubernetes/pull/103553), [@ehashman](https://github.com/ehashman)) [SIG Node]
- Omit comparison with boolean constant ([#101523](https://github.com/kubernetes/kubernetes/pull/101523), [@chuntaochen](https://github.com/chuntaochen)) [SIG CLI and Cloud Provider]
- Removed the feature flag for probe-level termination grace period from Kubelet. If a user wants to disable this feature on already created pods, they will have to delete and recreate the pods. ([#103168](https://github.com/kubernetes/kubernetes/pull/103168), [@raisaat](https://github.com/raisaat)) [SIG Apps and Node]
- Revert addition of Add `PersistentVolumeClaimDeletePoilcy` to `StatefulSet`API. ([#103747](https://github.com/kubernetes/kubernetes/pull/103747), [@mattcary](https://github.com/mattcary))
- Scheduler could be configured  to consider new resources beside CPU and memory,  GPU for example, for the score plugin of `NodeResourcesBalancedAllocation`. ([#101946](https://github.com/kubernetes/kubernetes/pull/101946), [@chendave](https://github.com/chendave)) [SIG Scheduling]
- Server Side Apply now treats all <Some>Selector fields as atomic (meaning the entire selector is managed by a single writer and updated together), since they contain interrelated and inseparable fields that do not merge in intuitive ways. ([#97989](https://github.com/kubernetes/kubernetes/pull/97989), [@Danil-Grigorev](https://github.com/Danil-Grigorev)) [SIG API Machinery]
- Suspend Job feature graduated to beta. Added the `action` label to Job controller sync metrics `job_sync_total` and `job_sync_duration_seconds`. ([#102022](https://github.com/kubernetes/kubernetes/pull/102022), [@adtac](https://github.com/adtac))
- The API documentation for the DaemonSet's `spec.updateStrategy.rollingUpdate.maxUnavailable` field was corrected to state that the value is rounded up. ([#101296](https://github.com/kubernetes/kubernetes/pull/101296), [@Miciah](https://github.com/Miciah))
- The `CSIServiceAccountToken` graduates to Ga and is unconditionally enabled. ([#103001](https://github.com/kubernetes/kubernetes/pull/103001), [@zshihang](https://github.com/zshihang))
- The `CertificateSigningRequest.certificates.k8s.io` API supports an optional expirationSeconds field to allow the client to request a particular duration for the issued certificate.  The default signer implementations provided by the Kubernetes controller manager will honor this field as long as it does not exceed the --cluster-signing-duration flag. ([#99494](https://github.com/kubernetes/kubernetes/pull/99494), [@enj](https://github.com/enj))
- The `EndpointSlicen Mirroring controller` no longer mirrors the `last-applied-configuration` annotation created by `kubectl` to update `EndpointSlices`. ([#102731](https://github.com/kubernetes/kubernetes/pull/102731), [@sharmarajdaksh](https://github.com/sharmarajdaksh))
- The `NetworkPolicyEndPort` is graduated to beta and is enabled by default. ([#102834](https://github.com/kubernetes/kubernetes/pull/102834), [@rikatz](https://github.com/rikatz))
- The `PodDeletionCost` feature has been promoted to beta, and enabled by default. ([#101080](https://github.com/kubernetes/kubernetes/pull/101080), [@ahg-g](https://github.com/ahg-g))
- The `Server Side Apply` treats certain structs as atomic. Meaning the entire selector field is managed by a single writer and updated together. ([#100684](https://github.com/kubernetes/kubernetes/pull/100684), [@Jefftree](https://github.com/Jefftree))
- The `ServiceAppProtocol` feature gate has been removed. It reached GA in Kubernetes ([#103190](https://github.com/kubernetes/kubernetes/pull/103190), [@robscott](https://github.com/robscott))
- The `TerminationGracePeriodSeconds` on pod specs and container probes should not be negative. Negative values of `TerminationGracePeriodSeconds` will be treated as the value `1s` on the delete path. Immutable field validation will be relaxed in order to update negative values. In a future release, negative values will not be permitted. ([#98866](https://github.com/kubernetes/kubernetes/pull/98866), [@wzshiming](https://github.com/wzshiming))
- The `kube-scheduler` component config `v1beta2` API available
  Three scheduler plugins deprecated (`NodeLabel`, `ServiceAffinity`, `NodePreferAvoidPods`). ([#99597](https://github.com/kubernetes/kubernetes/pull/99597), [@adtac](https://github.com/adtac))
- The `pod/eviction` subresource now accepts `policy/v1` eviction requests in addition to `policy/v1beta1` eviction requests ([#100724](https://github.com/kubernetes/kubernetes/pull/100724), [@liggitt](https://github.com/liggitt))
- The `podAffinity`, `NamespaceSelector` and the associated `CrossNamespaceAffinity` quota scope features graduate to Beta and they are now enabled by default. ([#101496](https://github.com/kubernetes/kubernetes/pull/101496), [@ahg-g](https://github.com/ahg-g))
- The `pods/ephemeralcontainers` API now returns and expects a `Pod` object instead of `EphemeralContainers`. This is incompatible with the previous alpha-level API. ([#101034](https://github.com/kubernetes/kubernetes/pull/101034), [@verb](https://github.com/verb)) [SIG Apps, Auth, CLI and Testing]
- The `v1.Node` and `.status.images[].names` are  now optional. ([#102159](https://github.com/kubernetes/kubernetes/pull/102159), [@roycaihw](https://github.com/roycaihw))
- The deprecated flag `--algorithm-provider` has been removed from `kube-scheduler`. Use instead `ComponentConfig` to configure the set of enabled plugins. ([#102239](https://github.com/kubernetes/kubernetes/pull/102239), [@Haleygo](https://github.com/Haleygo))
- The options `--ssh-user` and `--ssh-key` are removed. They only functioned on GCE, and only in-tree. Use the apiserver network proxy instead. ([#102297](https://github.com/kubernetes/kubernetes/pull/102297), [@deads2k](https://github.com/deads2k))
- Track Job completion through status and Pod finalizers, removing dependency on Pod tombstones. ([#98238](https://github.com/kubernetes/kubernetes/pull/98238), [@alculquicondor](https://github.com/alculquicondor)) [SIG API Machinery, Apps, Auth and Testing]
- Track ownership of scale subresource for all scalable resources i.e. Deployment, ReplicaSet, StatefulSet, ReplicationController, and Custom Resources. ([#98377](https://github.com/kubernetes/kubernetes/pull/98377), [@nodo](https://github.com/nodo)) [SIG API Machinery and Testing]

### Feature

- A `system-cluster-critical` pod should not get a low OOM Score. 
  
  As of now both `system-node-critical` and `system-cluster-critical` pods have -997 OOM score, making them one of the last processes to be OOMKilled. By definition `system-cluster-critical` pods can be scheduled elsewhere if there is a resource crunch on the node where as `system-node-critical` pods cannot be rescheduled. This was the reason for `system-node-critical` to have higher priority value than `system-cluster-critical`.  This change allows only `system-node-critical` priority class to have low OOMScore.
  
  action required
  If the user wants to have the pod to be OOMKilled last and the pod has `system-cluster-critical` priority class, it has to be changed to `system-node-critical` priority class to preserve the existing behavior ([#99729](https://github.com/kubernetes/kubernetes/pull/99729), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla))
- API Server tracing can now trace re-entrant api requests. ([#103218](https://github.com/kubernetes/kubernetes/pull/103218), [@dashpole](https://github.com/dashpole)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle and Instrumentation]
- APIServerTracing now collects spans from etcd client calls, and propagates context to etcd. ([#103216](https://github.com/kubernetes/kubernetes/pull/103216), [@dashpole](https://github.com/dashpole)) [SIG API Machinery, Cloud Provider and Instrumentation]
- APIServerTracing now collects spans from outgoing requests to admission webhooks. ([#103601](https://github.com/kubernetes/kubernetes/pull/103601), [@dashpole](https://github.com/dashpole)) [SIG API Machinery]
- Add a namespace label for all `apiserver_admission_* metrics`.
  Expand the histogram range to 0-10s for all `apiserver_admission_*_duration_seconds` metrics. ([#101208](https://github.com/kubernetes/kubernetes/pull/101208), [@voutcn](https://github.com/voutcn))
- Add unified map on CRI to support `cgroup v2`. Refer to https://github.com/opencontainers/runtime-spec/blob/master/config-linux.md#unified. ([#102578](https://github.com/kubernetes/kubernetes/pull/102578), [@payall4u](https://github.com/payall4u))
- Added BinaryData description to `kubectl describe` command. ([#100568](https://github.com/kubernetes/kubernetes/pull/100568), [@lauchokyip](https://github.com/lauchokyip))
- Added a new metric `apiserver_flowcontrol_request_concurrency_in_use` that shows the number of
  seats (concurrency) occupied by the currently executing requests in the API Priority and Fairness system. ([#102795](https://github.com/kubernetes/kubernetes/pull/102795), [@tkashem](https://github.com/tkashem))
- Added field-selector option for `kubectl top pod` ([#102155](https://github.com/kubernetes/kubernetes/pull/102155), [@lauchokyip](https://github.com/lauchokyip)) [SIG CLI]
- Added new metrics about API Priority and Fairness.  Each one has a label `priority_level`.  The last two also have a label `bound` taking values `min` and `max.
  - apiserver_flowcontrol_current_r: R(the time of the last change in state of the queues)
  - apiserver_flowcontrol_dispatch_r: R(the time of the latest request dispatch)
  - apiserver_flowcontrol_latest_s: S(the request last dispatched) = R(when that request starts executing in the virtual world)
  - apiserver_flowcontrol_next_s_bounds: min and max next S among non-empty queues
  - apiserver_flowcontrol_next_discounted_s_bounds: min and max next S - (sum [over requests executing] width * estimatedDuration) among non-empty queues ([#102859](https://github.com/kubernetes/kubernetes/pull/102859), [@MikeSpreitzer](https://github.com/MikeSpreitzer)) [SIG API Machinery and Instrumentation]
- Adding `--restart-kubelet` flag on E2E Node test suite ([#97028](https://github.com/kubernetes/kubernetes/pull/97028), [@knabben](https://github.com/knabben)) [SIG Node and Testing]
- Adds feature gate `KubeletInUserNamespace` which enables support for running kubelet in a user namespace.
  
  The user namespace has to be created before running kubelet.
  All the node components such as CRI need to be running in the same user namespace.
  
  When the feature gate is enabled, kubelet ignores errors that happens during setting the following sysctl values: `vm.overcommit_memory`, `vm.panic_on_oom`, `kernel.panic`, `kernel.panic_on_oops`, `kernel.keys.root_maxkeys`, `kernel.keys.root_maxbytes`. (These sysctl values for the host, not for the containers)
  
  kubelet also ignores an error during opening `/dev/kmsg`.
  This feature gate also allows kube-proxy to ignore an error during setting `RLIMIT_NOFILE`.
  
  This feature gate is especially useful for running Kubernetes inside Rootless Docker/Podman with `kind` or `minikube`. ([#92863](https://github.com/kubernetes/kubernetes/pull/92863), [@AkihiroSuda](https://github.com/AkihiroSuda)) [SIG Network, Node and Testing]
- Adds metrics for the delegated authenticator used by extension APIs that delegate authentication logic to the Kube API server. ([#99364](https://github.com/kubernetes/kubernetes/pull/99364), [@p0lyn0mial](https://github.com/p0lyn0mial))
- Adds metrics for the delegated authorizer used by extension APIs that delegate authorization logic to the Kube API server. ([#100339](https://github.com/kubernetes/kubernetes/pull/100339), [@p0lyn0mial](https://github.com/p0lyn0mial))
- Adds two kubemark flags, `--max-pods` and `--extended-resources`. ([#100267](https://github.com/kubernetes/kubernetes/pull/100267), [@Jeffwan](https://github.com/Jeffwan))
- An audit log entry will be generated when a `ValidatingAdmissionWebhook` is failing to open. ([#92739](https://github.com/kubernetes/kubernetes/pull/92739), [@cnphil](https://github.com/cnphil))
- Base images: Updated to
  - debian-base:buster-v1.6.0
  - debian-iptables:buster-v1.6.0 ([#100976](https://github.com/kubernetes/kubernetes/pull/100976), [@jindijamie](https://github.com/jindijamie))
- Base-images: Update to `debian-base:buster-v1.7.1` ([#102594](https://github.com/kubernetes/kubernetes/pull/102594), [@mengjiao-liu](https://github.com/mengjiao-liu))
- Deprecated warning message for `igonre-errors` flag. ([#102677](https://github.com/kubernetes/kubernetes/pull/102677), [@yuzhiquan](https://github.com/yuzhiquan))
- Endpoints that have more than 1000 endpoints will be truncated and the `endpoints.kubernetes.io/over-capacity` annotation on the Endpoints resource will be set to `truncated`. ([#103520](https://github.com/kubernetes/kubernetes/pull/103520), [@swetharepakula](https://github.com/swetharepakula)) [SIG Apps and Network]
- Expose `/debug/flags/v` to allow dynamically setting log level for kube-proxy. ([#98306](https://github.com/kubernetes/kubernetes/pull/98306), [@borgerli](https://github.com/borgerli)) [SIG Network]
- Expose container start time as `container_start_time_seconds` in the kubelet `/metrics/resource` endpoint. ([#102444](https://github.com/kubernetes/kubernetes/pull/102444), [@sanwishe](https://github.com/sanwishe))
- Extended resources defined in `LeastAllocated`, `MostAllocated` and `RequestedToCapacityRatio` plugin argument are bypassed by the scheduler if the incoming Pod doesn't request them in the pod spec. ([#103169](https://github.com/kubernetes/kubernetes/pull/103169), [@Huang-Wei](https://github.com/Huang-Wei))
- Feat: change parittion style to GPT on Windows ([#101412](https://github.com/kubernetes/kubernetes/pull/101412), [@andyzhangx](https://github.com/andyzhangx)) [SIG Storage and Windows]
- Features gates `EndpointSliceProxying` & `WindowsEndpointSliceProxying` graduates to GA and are unconditionally enabled. Kube-proxy will use EndpointSlices for endpoint information. ([#103451](https://github.com/kubernetes/kubernetes/pull/103451), [@swetharepakula](https://github.com/swetharepakula))
- Fluentd: isolate logging resources in separate namespace `logging` ([#68004](https://github.com/kubernetes/kubernetes/pull/68004), [@saravanan30erd](https://github.com/saravanan30erd))
- For `kubeadm`: add `--validity-period` flag for `kubeadm kubeconfig user` command. ([#100907](https://github.com/kubernetes/kubernetes/pull/100907), [@SataQiu](https://github.com/SataQiu))
- Implement `minReadySeconds` for the `StatefulSets`. ([#101316](https://github.com/kubernetes/kubernetes/pull/101316), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla))
- Improve logging of `APIService` availability changes in kube-apiserver. ([#101420](https://github.com/kubernetes/kubernetes/pull/101420), [@sttts](https://github.com/sttts))
- Introduce a feature gate `DisableCloudProviders` allowing to disable cloud-provider initialization in KAPI, KCM and kubelet.
  `DisableCloudProviders` FeatureGate is currently in Alpha, which means is currently disabled by default. Once the FeatureGate moves to beta, in-tree cloud providers would be disabled by default, and a user won't be able to specify `--cloud-provider=<aws|openstack|azure|gcp|vsphere>` anymore to any of KCM, KAPI or kubelet. Only a '--cloud-provider=external' would be allowed. CCM would have to run out-of-tree with CSI. ([#100136](https://github.com/kubernetes/kubernetes/pull/100136), [@Danil-Grigorev](https://github.com/Danil-Grigorev))
- JSON logging format is no longer available by default in non-core Kubernetes Components and require owners to opt in. ([#102869](https://github.com/kubernetes/kubernetes/pull/102869), [@mengjiao-liu](https://github.com/mengjiao-liu)) [SIG API Machinery, Cluster Lifecycle and Instrumentation]
- Kube-apiserver: the alpha PodSecurity feature can be enabled by passing `--feature-gates=PodSecurity=true`, and enables controlling allowed pods using namespace labels. See https://git.k8s.io/enhancements/keps/sig-auth/2579-psp-replacement for more details. ([#103099](https://github.com/kubernetes/kubernetes/pull/103099), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Auth, Instrumentation, Release, Security and Testing]
- Kube-proxy uses V1 `EndpointSlices`. ([#103306](https://github.com/kubernetes/kubernetes/pull/103306), [@swetharepakula](https://github.com/swetharepakula))
- Kubeadm: Add the `RootlessControlPlane` kubeadm specific feature gate (Alpha in 1.22, disabled by default). It can be used to enable an experimental feature that makes the control plane component static Pod containers for `kube-apiserver`, `kube-controller-manager`, `kube-scheduler` and `etcd` to run as a non-root users. ([#102158](https://github.com/kubernetes/kubernetes/pull/102158), [@vinayakankugoyal](https://github.com/vinayakankugoyal))
- Kubeadm: Set the `seccompProfile` to `runtime/default` in the `PodSecurityContext` of the control-plane components that run as static Pods. ([#100234](https://github.com/kubernetes/kubernetes/pull/100234), [@vinayakankugoyal](https://github.com/vinayakankugoyal))
- Kubeadm: add a new field `skipPhases` to `v1beta3` `InitConfiguration` and `JoinConfiguration` that can contain a list of phases to skip during "kubeadm init" and "kubeadm join". The flag "--skip-phases" takes precedence over this field. ([#101923](https://github.com/kubernetes/kubernetes/pull/101923), [@neolit123](https://github.com/neolit123))
- Kubeadm: add the `--dry-run` flag to the `control-plane` phase of "kubeadm init". ([#102722](https://github.com/kubernetes/kubernetes/pull/102722), [@vinayakankugoyal](https://github.com/vinayakankugoyal))
- Kubeadm: add the `imagePullPolicy` field in the `nodeRegistration` section of `InitConfiguration` and `JoinConfiguration` in `v1beta3`. This allows the user to specify the image pull policy during "kubeadm init" and "kubeadm join". The value of this field must be one of `Always`, `IfNotPresent` or `Never`. The default behavior continues to be `IfNotPresent`. ([#102901](https://github.com/kubernetes/kubernetes/pull/102901), [@wangyysde](https://github.com/wangyysde))
- Kubeadm: during "kubeadm init/join/upgrade", always default the `cgroupDriver` value in the `KubeletConfiguration` to `systemd`, unless the user was explicit about the value. See [configure-cgroup-driver](https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/configure-cgroup-driver/) for more details. ([#102133](https://github.com/kubernetes/kubernetes/pull/102133), [@pacoxu](https://github.com/pacoxu))
- Kubeadm: update `CoreDNS` to 1.8.4. Grant `CoreDNS` permissions to "list" and "watch" `EndpointSlice` objects to accommodate dual-stack support. ([#102466](https://github.com/kubernetes/kubernetes/pull/102466), [@pacoxu](https://github.com/pacoxu))
- Kubectl: add `LAST RESTART` column to `kubectl get pods` output. ([#100142](https://github.com/kubernetes/kubernetes/pull/100142), [@Ethyling](https://github.com/Ethyling))
- Kubemark's hollow-node will now print flags before starting. ([#101181](https://github.com/kubernetes/kubernetes/pull/101181), [@mm4tt](https://github.com/mm4tt))
- Kubernetes is now built with Golang 1.16.3 ([#101206](https://github.com/kubernetes/kubernetes/pull/101206), [@justaugustus](https://github.com/justaugustus)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Kubernetes is now built with Golang 1.16.4 ([#101809](https://github.com/kubernetes/kubernetes/pull/101809), [@justaugustus](https://github.com/justaugustus)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Kubernetes is now built with Golang 1.16.5. ([#102689](https://github.com/kubernetes/kubernetes/pull/102689), [@cpanato](https://github.com/cpanato))
- Kubernetes is now built with Golang 1.16.6 ([#103669](https://github.com/kubernetes/kubernetes/pull/103669), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Leader Migration for controller managers graduated to beta. ([#103533](https://github.com/kubernetes/kubernetes/pull/103533), [@jiahuif](https://github.com/jiahuif)) [SIG API Machinery and Cloud Provider]
- Make kubectl command headers default for beta. ([#103238](https://github.com/kubernetes/kubernetes/pull/103238), [@seans3](https://github.com/seans3)) [SIG CLI]
- Mark `net.ipv4.ip_unprivileged_port_start` as safe `sysctl`. ([#103326](https://github.com/kubernetes/kubernetes/pull/103326), [@pacoxu](https://github.com/pacoxu))
- Metrics server nanny has now poll period set to 30s (previously 5 minutes) to allow faster scaling of metrics server. ([#101869](https://github.com/kubernetes/kubernetes/pull/101869), [@olagacek](https://github.com/olagacek)) [SIG Cloud Provider and Instrumentation]
- NetworkPolicy validation framework support for windows. ([#98077](https://github.com/kubernetes/kubernetes/pull/98077), [@jayunit100](https://github.com/jayunit100))
- New feature gate `ExpandedDNSConfig` is now available. This feature allows Kubernetes to have expanded DNS configuration. ([#100651](https://github.com/kubernetes/kubernetes/pull/100651), [@gjkim42](https://github.com/gjkim42))
- New metrics: `apiserver_kube_aggregator_x509_missing_san_total` and `apiserver_webhooks_x509_missing_san_total`. This metric measures a number of connections to webhooks/aggregated API servers that use certificates without Subject Alternative Names. It being non-zero is a warning sign that these connections will stop functioning in the future since Golang is going to deprecate x509 certificate subject Common Names for server hostname verification. ([#95396](https://github.com/kubernetes/kubernetes/pull/95396), [@stlaz](https://github.com/stlaz)) [SIG API Machinery, Auth and Instrumentation]
- Node Problem Detector is now available for GCE Windows nodes. ([#101539](https://github.com/kubernetes/kubernetes/pull/101539), [@jeremyje](https://github.com/jeremyje)) [SIG Cloud Provider, Node and Windows]
- Promote Cronjobs storage version to `batch/v1`. ([#102363](https://github.com/kubernetes/kubernetes/pull/102363), [@mengjiao-liu](https://github.com/mengjiao-liu))
- Promote `CronJobControllerV2` flag to GA, with removal in 1.23. ([#102529](https://github.com/kubernetes/kubernetes/pull/102529), [@soltysh](https://github.com/soltysh))
- Promote `EndpointSliceTerminatingCondition` to Beta. This enables the `terminating` and `serving` conditions for EndpointSlice by default. ([#103596](https://github.com/kubernetes/kubernetes/pull/103596), [@andrewsykim](https://github.com/andrewsykim))
- Run etcd as non-root on GCE provider ([#100635](https://github.com/kubernetes/kubernetes/pull/100635), [@cindy52](https://github.com/cindy52))
- Scheduler nows provides an option for plugin developers to move `Pods` to activeQ. ([#103383](https://github.com/kubernetes/kubernetes/pull/103383), [@Huang-Wei](https://github.com/Huang-Wei))
- Secret values are now masked by default in `kubectl diff` output. ([#96084](https://github.com/kubernetes/kubernetes/pull/96084), [@loozhengyuan](https://github.com/loozhengyuan))
- Services with `externalTrafficPolicy: Local` now support graceful termination when using
  the iptables or ipvs mode of kube-proxy with `EndpointSlices` enabled. Specifically, if a
  connection for such a service arrives on a node when there are no "Ready" endpoints
  for the service, but there is at least one Terminating pod for that service on the node,
  then kube-proxy will send the traffic to the Terminating pod rather than dropping it. This
  patches up a race condition between when a pod is killed and when the external load
  balancer notices that it has been killed. ([#97238](https://github.com/kubernetes/kubernetes/pull/97238), [@andrewsykim](https://github.com/andrewsykim))
- Shell completion has been migrated to Cobra's go solution.  `kubectl` is now smarter about disabling file completion when it does not apply.  Furthermore, completion for the `cp` command does not show all files unless the user has started typing something. ([#96087](https://github.com/kubernetes/kubernetes/pull/96087), [@marckhouzam](https://github.com/marckhouzam)) [SIG CLI]
- Some of the in-tree storage drivers indicate support for the `MetricsProvider` interface, but fail to configure this for BlockMode volumes. With a recent change, `Kubelet` will call `GetMetrics()` for BlockMode volumes, and the in-tree drivers that miss the support cause a Go panic. Now the in-tree storage drivers that support BlockMode volumes, will return the Capacity of the volume in the `GetMetrics()` call. ([#101587](https://github.com/kubernetes/kubernetes/pull/101587), [@nixpanic](https://github.com/nixpanic))
- Support `FakeClientset` match subresource. ([#100939](https://github.com/kubernetes/kubernetes/pull/100939), [@wzshiming](https://github.com/wzshiming))
- The "Leader Migration" now support a wildcard component name and the default value. ([#102711](https://github.com/kubernetes/kubernetes/pull/102711), [@jiahuif](https://github.com/jiahuif))
- The CSI driver supports the NodeServiceCapability `VOLUME_MOUNT_GROUP` and the `DelegateFSGroupToCSIDriver` feature gate is enabled, kubelet will delegate applying FSGroup to the driver by passing it to `NodeStageVolume` and `NodePublishVolume`, regardless of what other `FSGroup` policies are set, this is an alpha feature. ([#103244](https://github.com/kubernetes/kubernetes/pull/103244), [@verult](https://github.com/verult))
- The Memory Manager feature graduates to Beta and it is enabled by default. ([#101947](https://github.com/kubernetes/kubernetes/pull/101947), [@cynepco3hahue](https://github.com/cynepco3hahue))
- The `BoundServiceAccountTokenVolume` graduates to GA and thus will be unconditionally enabled. The feature gate is going to be removed in 1.23. ([#101992](https://github.com/kubernetes/kubernetes/pull/101992), [@zshihang](https://github.com/zshihang))
- The `EmptyDir` memory backed volumes are sized as the the minimum of pod allocatable memory on a host and an optional explicit user provided value. ([#101048](https://github.com/kubernetes/kubernetes/pull/101048), [@dims](https://github.com/dims))
- The `HugePageStorageMediumSize` feature graduates to GA and unconditionally enabled. Allowing unconditional usage of multiple sizes huge page resources on a container level. ([#99144](https://github.com/kubernetes/kubernetes/pull/99144), [@bart0sh](https://github.com/bart0sh))
- The `IngressClassNamespacedParams` feature gate has graduated to beta and is enabled by default. This means IngressClass resource will now have two new fields - `spec.paramters.namespace` and `spec.parameters.scope`. ([#101711](https://github.com/kubernetes/kubernetes/pull/101711), [@hbagdi](https://github.com/hbagdi))
- The `LogarithmicScaleDown` feature graduates to Beta and enabled by default. ([#101767](https://github.com/kubernetes/kubernetes/pull/101767), [@damemi](https://github.com/damemi))
- The `NamespaceDefaultLabelName` is promoted to GA in this release. All Namespace API objects have a `kubernetes.io/metadata.name` label matching their `metadata.name` field to allow selecting any `namespace` by its name using a label selector. ([#101342](https://github.com/kubernetes/kubernetes/pull/101342), [@rosenhouse](https://github.com/rosenhouse))
- The `ServiceInternalTrafficPolicy` feature graduates to Beta and enable by default, which enables the `internalTrafficPolicy` field of Service by default. ([#103462](https://github.com/kubernetes/kubernetes/pull/103462), [@andrewsykim](https://github.com/andrewsykim))
- The `ServiceLBNodePortControl` graduates to Beta and is enabled by default. ([#100412](https://github.com/kubernetes/kubernetes/pull/100412), [@hanlins](https://github.com/hanlins))
- The `SetHostnameAsFQDN` graduates to GA and thus will be unconditionally disabled. ([#101294](https://github.com/kubernetes/kubernetes/pull/101294), [@javidiaz](https://github.com/javidiaz))
- The `WarningHeader` feature is now GA and is unconditionally enabled. The `apiserver_requested_deprecated_apis` metric has graduated to stable status. The `WarningHeader` feature-gate is no longer operative and will be removed in v1.24. ([#100754](https://github.com/kubernetes/kubernetes/pull/100754), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Instrumentation and Testing]
- The `kubectl debug` is able to create ephemeral containers in pre-1.22 clusters with the `EphemeralContainers` feature enabled. Note that versions of kubectl prior to 1.22 are unable to create ephemeral containers in clusters version 1.22 and greater due to an API change. ([#103292](https://github.com/kubernetes/kubernetes/pull/103292), [@verb](https://github.com/verb))
- The client-go credential plugins are now GA and are enabled by default. ([#102890](https://github.com/kubernetes/kubernetes/pull/102890), [@ankeesler](https://github.com/ankeesler))
- The feature gate `SSA` graduated to GA in v1.22 and therefore is unconditionally enabled. ([#100139](https://github.com/kubernetes/kubernetes/pull/100139), [@Jefftree](https://github.com/Jefftree))
- The job controller removes running pods when the number of completions is achieved. ([#99963](https://github.com/kubernetes/kubernetes/pull/99963), [@alculquicondor](https://github.com/alculquicondor))
- The kubeconfig is now exposed in `the kube-scheduler` framework handle. Out-of-tree plugins can leverage that to build CRD informers easily. ([#100644](https://github.com/kubernetes/kubernetes/pull/100644), [@Huang-Wei](https://github.com/Huang-Wei))
- The new flag `--chunk-size=SIZE` for `kubectl drain` has been promoted to beta, and enabled by default. This flag  may be used to alter the number of items or disable this feature when `0` is passed. ([#100148](https://github.com/kubernetes/kubernetes/pull/100148), [@KnVerey](https://github.com/KnVerey))
- The new flag `--chunk-size=SIZE` has been added to `kubectl describe`. This flag may be used to alter the number of items or disable this feature when `0` is passed. ([#101171](https://github.com/kubernetes/kubernetes/pull/101171), [@KnVerey](https://github.com/KnVerey))
- The pod resource API will provide memory manager metrics in the case when the memory manager feature gate is enabled, and the memory manager policy is static. ([#101030](https://github.com/kubernetes/kubernetes/pull/101030), [@cynepco3hahue](https://github.com/cynepco3hahue))
- The prefer nominated node graduates to Beta and enabld by default. ([#102201](https://github.com/kubernetes/kubernetes/pull/102201), [@chendave](https://github.com/chendave))
- Update etcd version to 3.5.0-beta.3. ([#102062](https://github.com/kubernetes/kubernetes/pull/102062), [@serathius](https://github.com/serathius))
- Update the Debian images to pick up CVE fixes in the base images:
  - Update the `debian-base` image to v1.7.0
  - Update the `debian-iptables` image to v1.6.1 ([#102302](https://github.com/kubernetes/kubernetes/pull/102302), [@xmudrii](https://github.com/xmudrii))
- Update the setcap image to buster-v2.0.1. ([#102377](https://github.com/kubernetes/kubernetes/pull/102377), [@xmudrii](https://github.com/xmudrii))
- Update the system-validators library to v1.5.0. Includes validation for `seccomp` and fixes a stdout/stderr problem in the Docker validator. ([#103390](https://github.com/kubernetes/kubernetes/pull/103390), [@ironyman](https://github.com/ironyman))
- Updates the following images to pick up CVE fixes:
  - `debian` to v1.8.0
  - `debian-iptables` to v1.6.5
  - `setcap` to v2.0.3 ([#103235](https://github.com/kubernetes/kubernetes/pull/103235), [@thejoycekung](https://github.com/thejoycekung)) [SIG API Machinery, Release and Testing]
- Warnings for the use of deprecated and known-bad values in pod specs are now sent. ([#101688](https://github.com/kubernetes/kubernetes/pull/101688), [@liggitt](https://github.com/liggitt))
- Watch requests are now handled throttled by priority and fairness filter in `kube-apiserver`. ([#102171](https://github.com/kubernetes/kubernetes/pull/102171), [@wojtek-t](https://github.com/wojtek-t))
- You can use this Builder function to create events Field Selector ([#101817](https://github.com/kubernetes/kubernetes/pull/101817), [@cndoit18](https://github.com/cndoit18)) [SIG API Machinery and Scalability]
- `Scheduler` now registers event handlers dynamically. ([#101394](https://github.com/kubernetes/kubernetes/pull/101394), [@Huang-Wei](https://github.com/Huang-Wei))
- `kubectl`: Enable using protocol buffers to request Metrics API. ([#102039](https://github.com/kubernetes/kubernetes/pull/102039), [@serathius](https://github.com/serathius))

### Documentation

- The command`kubectl debug` will now print a warning message when using the `--target` option since many container runtimes do not support this yet. ([#101074](https://github.com/kubernetes/kubernetes/pull/101074), [@verb](https://github.com/verb))

### Failing Test

- Fixed generic ephemeal volumes with `OwnerReferencesPermissionEnforcement` admission plugin enabled. ([#101186](https://github.com/kubernetes/kubernetes/pull/101186), [@jsafrane](https://github.com/jsafrane))
- Fixes `kubectl drain --dry-run=server`. ([#100206](https://github.com/kubernetes/kubernetes/pull/100206), [@KnVerey](https://github.com/KnVerey))
- Fixes an overly restrictive conformance test to accept service account tokens signed by an ECDSA key ([#100680](https://github.com/kubernetes/kubernetes/pull/100680), [@smira](https://github.com/smira)) [SIG Architecture, Auth and Testing]
- Fixes the `should receive events on concurrent watches in same order` conformance test to work properly on clusters that auto-create additional configmaps in namespaces. ([#101950](https://github.com/kubernetes/kubernetes/pull/101950), [@liggitt](https://github.com/liggitt))
- Resolves an issue with the "ServiceAccountIssuerDiscovery should support OIDC discovery" conformance test failing on clusters which are configured with issuers outside the cluster ([#101589](https://github.com/kubernetes/kubernetes/pull/101589), [@mtaufen](https://github.com/mtaufen)) [SIG Auth and Testing]

### Bug or Regression

- Added jitter factor to lease controller that better smears load on kube-apiserver over time. ([#101652](https://github.com/kubernetes/kubernetes/pull/101652), [@marseel](https://github.com/marseel)) [SIG API Machinery and Scalability]
- Added privileges for `EndpointSlice` to the default view & edit RBAC roles. ([#101203](https://github.com/kubernetes/kubernetes/pull/101203), [@mtougeron](https://github.com/mtougeron))
- After DBus restarts, make `GracefulNodeShutdown` work again ([#100369](https://github.com/kubernetes/kubernetes/pull/100369), [@wzshiming](https://github.com/wzshiming))
- Aggregate errors when putting vmss. ([#98350](https://github.com/kubernetes/kubernetes/pull/98350), [@nilo19](https://github.com/nilo19))
- Aggregate write permissions on events to users with edit and admin role. ([#102858](https://github.com/kubernetes/kubernetes/pull/102858), [@tumido](https://github.com/tumido))
- Aggregated roles no longer include write access to `EndpointSlices`. This rolls back part of a change that was introduced earlier in the Kubernetes 1.22 cycle. ([#103703](https://github.com/kubernetes/kubernetes/pull/103703), [@robscott](https://github.com/robscott))
- Applying fix for not deleting existing public IP when a service is deleted in Azure. ([#100694](https://github.com/kubernetes/kubernetes/pull/100694), [@nilo19](https://github.com/nilo19))
- Applying fix for not tagging static public IP. ([#101752](https://github.com/kubernetes/kubernetes/pull/101752), [@nilo19](https://github.com/nilo19))
- Applying fix so that deleting non-existing disk returns success. ([#102083](https://github.com/kubernetes/kubernetes/pull/102083), [@andyzhangx](https://github.com/andyzhangx))
- Applying fix: cleanup outdated routes. ([#102935](https://github.com/kubernetes/kubernetes/pull/102935), [@nilo19](https://github.com/nilo19))
- Avoid caching the Azure VMSS instances whose network profile is nil ([#100948](https://github.com/kubernetes/kubernetes/pull/100948), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Azure: Avoid setting cached Sku when updating VMSS and VMSS instances. ([#102005](https://github.com/kubernetes/kubernetes/pull/102005), [@feiskyer](https://github.com/feiskyer))
- Azurefile: Normalize share name to not include the capital letters ([#100731](https://github.com/kubernetes/kubernetes/pull/100731), [@kassarl](https://github.com/kassarl))
- Chain the field manager creation calls in `newDefaultFieldManager` to be explicit about the order of operations. ([#101076](https://github.com/kubernetes/kubernetes/pull/101076), [@kevindelgado](https://github.com/kevindelgado))
- Disruption controller shouldn't error while syncing for unmanaged pods. ([#103414](https://github.com/kubernetes/kubernetes/pull/103414), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Apps and Testing]
- Ensure service is deleted when the Azure resource group has been deleted. ([#100944](https://github.com/kubernetes/kubernetes/pull/100944), [@feiskyer](https://github.com/feiskyer))
- Ensures `ExecProbeTimeout=false` kubelet feature gate with dockershim is taken into account, when the exec probe takes longer than `timeoutSeconds` configuration. ([#100200](https://github.com/kubernetes/kubernetes/pull/100200), [@jackfrancis](https://github.com/jackfrancis))
- Expose `rest_client_rate_limiter_duration_seconds` metric to component-base to track client side rate limiter latency in seconds. Broken down by verb and URL. ([#100311](https://github.com/kubernetes/kubernetes/pull/100311), [@IonutBajescu](https://github.com/IonutBajescu)) [SIG API Machinery, Cluster Lifecycle and Instrumentation]
- Fire an event when failing to open `NodePort`. ([#100599](https://github.com/kubernetes/kubernetes/pull/100599), [@masap](https://github.com/masap))
- Fix Azure node public IP fetching issues from instance metadata service when the node is part of standard load balancer backend pool. ([#100690](https://github.com/kubernetes/kubernetes/pull/100690), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix `EndpointSlice` describe panic when an Endpoint doesn't have zone. ([#101025](https://github.com/kubernetes/kubernetes/pull/101025), [@tnqn](https://github.com/tnqn))
- Fix `kubectl set env` or `resources` not working for initcontainers. ([#101669](https://github.com/kubernetes/kubernetes/pull/101669), [@carlory](https://github.com/carlory))
- Fix `kubectl` alpha debug node does not work on tainted(NoExecute) nodes and tolerate everything. ([#98431](https://github.com/kubernetes/kubernetes/pull/98431), [@wawa0210](https://github.com/wawa0210))
- Fix a bug on the `endpointslicemirroring` controller where endpoint `NotReadyAddresses` were mirrored as Ready to the corresponding `EndpointSlice`. ([#102683](https://github.com/kubernetes/kubernetes/pull/102683), [@aojea](https://github.com/aojea))
- Fix a bug that a preemptor pod may exist as a phantom in the scheduler. ([#102498](https://github.com/kubernetes/kubernetes/pull/102498), [@Huang-Wei](https://github.com/Huang-Wei))
- Fix a number of race conditions in the kubelet when pods are starting up or shutting down that might cause pods to take a long time to shut down. ([#102344](https://github.com/kubernetes/kubernetes/pull/102344), [@smarterclayton](https://github.com/smarterclayton)) [SIG Apps, Node, Storage and Testing]
- Fix an issue with `kubectl` on certain older version of Windows or when legacy console mode is enabled on Windows 8 which causes `kubectl exec` to crash. ([#102825](https://github.com/kubernetes/kubernetes/pull/102825), [@n4j](https://github.com/n4j))
- Fix availability set cache in vmss cache ([#100110](https://github.com/kubernetes/kubernetes/pull/100110), [@CecileRobertMichon](https://github.com/CecileRobertMichon)) [SIG Cloud Provider]
- Fix how nulls are handled in array and objects in [json patches](https://github.com/evanphx/json-patch). ([#102467](https://github.com/kubernetes/kubernetes/pull/102467), [@pacoxu](https://github.com/pacoxu))
- Fix panic when `kubectl create ingress` has annotation flag and an empty value set. ([#101377](https://github.com/kubernetes/kubernetes/pull/101377), [@rikatz](https://github.com/rikatz))
- Fix performance regression for update and apply operations on large CRDs. ([#103318](https://github.com/kubernetes/kubernetes/pull/103318), [@jpbetz](https://github.com/jpbetz)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Storage]
- Fix raw block mode CSI `NodePublishVolume` stage miss pod info. ([#99069](https://github.com/kubernetes/kubernetes/pull/99069), [@phantooom](https://github.com/phantooom))
- Fix resource enforcement when using systemd cgroup driver ([#102147](https://github.com/kubernetes/kubernetes/pull/102147), [@kolyshkin](https://github.com/kolyshkin))
- Fix rounding of volume storage requests. ([#100100](https://github.com/kubernetes/kubernetes/pull/100100), [@maxlaverse](https://github.com/maxlaverse))
- Fix runtime container status for `PostStart` hook error. ([#100608](https://github.com/kubernetes/kubernetes/pull/100608), [@pacoxu](https://github.com/pacoxu))
- Fix scoring for `NodeResourcesMostAllocated` and `NodeResourcesBalancedAllocation` plugins when nodes have containers with no requests. This was leaving to under-utilization of small nodes. ([#102925](https://github.com/kubernetes/kubernetes/pull/102925), [@alculquicondor](https://github.com/alculquicondor))
- Fix the code is leaking the defaulting between unrelated pod instances. ([#103284](https://github.com/kubernetes/kubernetes/pull/103284), [@kebe7jun](https://github.com/kebe7jun)) [SIG CLI]
- Fix winkernel kube-proxy to only use dual stack when host and networking supports it ([#101047](https://github.com/kubernetes/kubernetes/pull/101047), [@jsturtevant](https://github.com/jsturtevant)) [SIG Network and Windows]
- Fix: Azure file inline volume namespace issue in CSI migration translation ([#101235](https://github.com/kubernetes/kubernetes/pull/101235), [@andyzhangx](https://github.com/andyzhangx))
- Fix: Bug in `kube-proxy` latency metrics to calculate only the latency value for the `Endpoints` that are created after it starts running. This is needed because all the `Endpoints` objects are processed on restarts, independently when they were. ([#100861](https://github.com/kubernetes/kubernetes/pull/100861), [@aojea](https://github.com/aojea))
- Fix: avoid nil-pointer panic when checking the frontend IP configuration ([#101739](https://github.com/kubernetes/kubernetes/pull/101739), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: display of Job completion mode in `kubectl describe`. ([#101160](https://github.com/kubernetes/kubernetes/pull/101160), [@alculquicondor](https://github.com/alculquicondor))
- Fix: return empty VMAS name if using standalone VM ([#103470](https://github.com/kubernetes/kubernetes/pull/103470), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: set "host is down" as corrupted mount. When SMB server is down, there is no way to terminate pod which is using SMB mount, would get an error. ([#101398](https://github.com/kubernetes/kubernetes/pull/101398), [@andyzhangx](https://github.com/andyzhangx))
- Fix: using NVMe AWS EBS volumes partitions. ([#100500](https://github.com/kubernetes/kubernetes/pull/100500), [@jsafrane](https://github.com/jsafrane))
- Fixed 'kubelet' runtime panic for timed-out portforward streams. ([#102489](https://github.com/kubernetes/kubernetes/pull/102489), [@saschagrunert](https://github.com/saschagrunert))
- Fixed SELinux relabeling of CSI volumes after CSI driver failure. ([#103154](https://github.com/kubernetes/kubernetes/pull/103154), [@jsafrane](https://github.com/jsafrane)) [SIG Node and Storage]
- Fixed `garbage collection` of dangling `VolumeAttachments` for `PersistentVolumes` migrated to CSI on startup of `kube-controller-manager`. ([#102176](https://github.com/kubernetes/kubernetes/pull/102176), [@timebertt](https://github.com/timebertt))
- Fixed `port-forward` memory leak for long-running and heavily used connections. ([#99839](https://github.com/kubernetes/kubernetes/pull/99839), [@saschagrunert](https://github.com/saschagrunert))
- Fixed a bug due to which the controller was not populating the `lastSuccessfulTime` field added to `cronjob.status` in `batch/v1`. ([#102642](https://github.com/kubernetes/kubernetes/pull/102642), [@alaypatel07](https://github.com/alaypatel07))
- Fixed a bug that `kubectl create configmap` always returns zero exit code when failed. ([#101780](https://github.com/kubernetes/kubernetes/pull/101780), [@nak3](https://github.com/nak3)) [SIG CLI]
- Fixed a bug that scheduler extenders are not called on preemptions. ([#103019](https://github.com/kubernetes/kubernetes/pull/103019), [@ordovicia](https://github.com/ordovicia))
- Fixed a bug where `startupProbe` stopped working after a container's first restart. ([#101093](https://github.com/kubernetes/kubernetes/pull/101093), [@wzshiming](https://github.com/wzshiming))
- Fixed an issue blocking azure auth to prompt to device code authentication flow when refresh token expires. ([#102063](https://github.com/kubernetes/kubernetes/pull/102063), [@tdihp](https://github.com/tdihp))
- Fixed false-positive uncertain volume attachments, which led to unexpected detachment of CSI migrated volumes ([#101737](https://github.com/kubernetes/kubernetes/pull/101737), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG Apps and Storage]
- Fixed mounting of NFS volumes when IPv6 address is used as a server. ([#101067](https://github.com/kubernetes/kubernetes/pull/101067), [@Elbehery](https://github.com/Elbehery)) [SIG Storage]
- Fixed starting new pods after previous pod timed out unmounting its volumes. ([#100183](https://github.com/kubernetes/kubernetes/pull/100183), [@jsafrane](https://github.com/jsafrane))
- Fixed very rare volume corruption when a pod is deleted while kubelet is offline. ([#102059](https://github.com/kubernetes/kubernetes/pull/102059), [@jsafrane](https://github.com/jsafrane))
- Fixes a data race issue in the priority and fairness API server filter. ([#100638](https://github.com/kubernetes/kubernetes/pull/100638), [@tkashem](https://github.com/tkashem))
- Fixes issue with websocket-based watches of Service objects not closing correctly on timeout. ([#102539](https://github.com/kubernetes/kubernetes/pull/102539), [@liggitt](https://github.com/liggitt))
- For `kubeadm`: support for custom imagetags for etcd images which contain build metadata, when imagetags are in the form of version_metadata. For instance, if the etcd version is v3.4.13+patch.0, the supported imagetag would be v3.4.13_patch.0 ([#100350](https://github.com/kubernetes/kubernetes/pull/100350), [@jr0d](https://github.com/jr0d))
- For vSphere: fix regression during attach disk if datastore is within a storage folder or datastore cluster. ([#102892](https://github.com/kubernetes/kubernetes/pull/102892), [@gnufied](https://github.com/gnufied))
- GCE Windows clusters have their TCP/IP parameters are set to GCE's recommended values. ([#103057](https://github.com/kubernetes/kubernetes/pull/103057), [@jeremyje](https://github.com/jeremyje)) [SIG Cloud Provider and Windows]
- GCE Windows will no longer install Docker on containerd nodes. ([#101747](https://github.com/kubernetes/kubernetes/pull/101747), [@jeremyje](https://github.com/jeremyje)) [SIG Cloud Provider and Windows]
- Generated OpenAPI now correctly specifies 201 as a possible response code for PATCH operations. ([#100141](https://github.com/kubernetes/kubernetes/pull/100141), [@brendandburns](https://github.com/brendandburns))
- Graceful termination will now be honored when deleting a collection of pods. ([#100101](https://github.com/kubernetes/kubernetes/pull/100101), [@deads2k](https://github.com/deads2k))
- If `kube-proxy` mode  is userspace do not enable `EndpointSlices`. ([#100913](https://github.com/kubernetes/kubernetes/pull/100913), [@JornShen](https://github.com/JornShen))
- Kubeadm: allow passing the flag `--log-file` if `--config` is passed. If you wish to log to a file you must also pass `--logtostderr=false` or `--alsologtostderr=true`. Alternatively you can pipe to a file using "kubeadm ... | tee ...". ([#101449](https://github.com/kubernetes/kubernetes/pull/101449), [@CaoDonghui123](https://github.com/CaoDonghui123))
- Kubeadm: enable `--experimental-patches` flag for `kubeadm join phase control-plane-join all` command. ([#101110](https://github.com/kubernetes/kubernetes/pull/101110), [@SataQiu](https://github.com/SataQiu))
- Kubeadm: fix a bug where `kubeadm join` for control plane nodes would download certificates and keys from the cluster, but would not write publicly readable certificates and public keys with mode `0644` and instead use mode `0600`. ([#103313](https://github.com/kubernetes/kubernetes/pull/103313), [@neolit123](https://github.com/neolit123))
- Kubeadm: fix the bug that `kubeadm` only uses the first hash in `caCertHashes` to verify the root CA. ([#101977](https://github.com/kubernetes/kubernetes/pull/101977), [@SataQiu](https://github.com/SataQiu))
- Kubeadm: remove the "ephemeral_storage" request from the etcd static pod that kubeadm deploys on stacked etcd control plane nodes. This request has caused sporadic failures on some setups due to a problem in the kubelet with cadvisor and the LocalStorageCapacityIsolation feature gate. See this issue for more details: https://github.com/kubernetes/kubernetes/issues/99305 ([#102673](https://github.com/kubernetes/kubernetes/pull/102673), [@jackfrancis](https://github.com/jackfrancis)) [SIG Cluster Lifecycle]
- Kubeadm: when using a custom image repository for CoreDNS kubeadm now will append the `coredns` image name instead of `coredns/coredns`, thus restoring the behaviour existing before the v1.21 release. Users who rely on nested folder for the coredns image should set the `clusterConfiguration.dns.imageRepository` value including the nested path name (e.g using `registry.company.xyz/coredns` will force kubeadm to use `registry.company.xyz/coredns/coredns` image). No action is needed if using the default registry (k8s.gcr.io). ([#102502](https://github.com/kubernetes/kubernetes/pull/102502), [@ykakarap](https://github.com/ykakarap))
- Kubelet: improve the performance when waiting for a synchronization of the node list with the `kube-apiserver`. ([#99336](https://github.com/kubernetes/kubernetes/pull/99336), [@neolit123](https://github.com/neolit123))
- Kubelet: the returned value for PodIPs is the same in the Downward API and in the `pod.status.PodIPs` field ([#103307](https://github.com/kubernetes/kubernetes/pull/103307), [@aojea](https://github.com/aojea))
- Limit vSphere volume name to 63 characters long. ([#100404](https://github.com/kubernetes/kubernetes/pull/100404), [@gnufied](https://github.com/gnufied))
- Logging for GCE Windows clusters will be more accurate and complete when using Fluent bit. ([#101271](https://github.com/kubernetes/kubernetes/pull/101271), [@jeremyje](https://github.com/jeremyje))
- Metrics Server will use Addon Manager 1.8.3 ([#103541](https://github.com/kubernetes/kubernetes/pull/103541), [@jbartosik](https://github.com/jbartosik)) [SIG Cloud Provider and Instrumentation]
- Output for `kubectl describe podsecuritypolicy` is now kind specific and cleaner ([#101436](https://github.com/kubernetes/kubernetes/pull/101436), [@KnVerey](https://github.com/KnVerey))
- Parsing of cpuset information now properly detects more invalid input such as `1--3` or `10-6`. ([#100565](https://github.com/kubernetes/kubernetes/pull/100565), [@lack](https://github.com/lack))
- Pods that are known to the kubelet to have previously been Running should not revert to Pending state, the kubelet will now infer a termination. ([#102821](https://github.com/kubernetes/kubernetes/pull/102821), [@ehashman](https://github.com/ehashman))
- Prevent Kubelet stuck in `DiskPressure` when `imagefs.minReclaim` is set ([#99095](https://github.com/kubernetes/kubernetes/pull/99095), [@maxlaverse](https://github.com/maxlaverse))
- Reduces delay initializing on non-AWS platforms docker runtime. ([#93260](https://github.com/kubernetes/kubernetes/pull/93260), [@nckturner](https://github.com/nckturner)) [SIG Cloud Provider]
- Register/Deregister Targets in chunks for AWS TargetGroup ([#101592](https://github.com/kubernetes/kubernetes/pull/101592), [@M00nF1sh](https://github.com/M00nF1sh)) [SIG Cloud Provider]
- Removed `/sbin/apparmor_parser` requirement for the AppArmor host validation.
  This allows using AppArmor on distributions which ship the binary in a different path. ([#97968](https://github.com/kubernetes/kubernetes/pull/97968), [@saschagrunert](https://github.com/saschagrunert)) [SIG Node and Testing]
- Renames the timeout field for the `DelegatingAuthenticationOptions` to `TokenRequestTimeout` and set the timeout only for the token review client. Previously the timeout was also applied to watches making them reconnecting every 10 seconds. ([#100959](https://github.com/kubernetes/kubernetes/pull/100959), [@p0lyn0mial](https://github.com/p0lyn0mial))
- Reorganized iptables rules to reduce rules in `KUBE-SERVICES` and `KUBE-NODEPORTS`. ([#96959](https://github.com/kubernetes/kubernetes/pull/96959), [@tssurya](https://github.com/tssurya))
- Respect annotation size limit for server-side apply updates to the client-side apply annotation. Also, fix opt-out of this behavior by setting the client-side apply annotation to the empty string. ([#102105](https://github.com/kubernetes/kubernetes/pull/102105), [@julianvmodesto](https://github.com/julianvmodesto)) [SIG API Machinery]
- Retry `FibreChannel` devices cleanup after error to ensure `FibreChannel` device is detached before it can be used on another node. ([#101862](https://github.com/kubernetes/kubernetes/pull/101862), [@jsafrane](https://github.com/jsafrane))
- Support correct sorting for cpu, memory, storage, ephemeral-storage, hugepages, and attachable-volumes. ([#100435](https://github.com/kubernetes/kubernetes/pull/100435), [@lauchokyip](https://github.com/lauchokyip))
- Switch scheduler to generate the merge patch on pod status instead of the full pod ([#103133](https://github.com/kubernetes/kubernetes/pull/103133), [@marwanad](https://github.com/marwanad)) [SIG Scheduling]
- The `EndpointSlice` IP validation now matches `Endpoints` IP validation. ([#101084](https://github.com/kubernetes/kubernetes/pull/101084), [@robscott](https://github.com/robscott))
- The `kube-apiserver` now reports the synthetic verb when logging requests, better explaining the user intent and matching what is reported in the metrics. ([#102934](https://github.com/kubernetes/kubernetes/pull/102934), [@lavalamp](https://github.com/lavalamp))
- The `kube-controller-manager'` sets the upper-bound timeout limit for outgoing requests to 70s. Previously ([#99358](https://github.com/kubernetes/kubernetes/pull/99358), [@p0lyn0mial](https://github.com/p0lyn0mial))
- The `kube-proxy` log now shows the "Skipping topology aware endpoint filtering since no hints were provided for zone" warning under the right conditions. ([#101857](https://github.com/kubernetes/kubernetes/pull/101857), [@dervoeti](https://github.com/dervoeti))
- The `kubectl create service` now respects the `namespace` flag. ([#101005](https://github.com/kubernetes/kubernetes/pull/101005), [@zxh326](https://github.com/zxh326))
- The `kubectl get` now truncates multi-line strings to avoid breaking printing ([#103514](https://github.com/kubernetes/kubernetes/pull/103514), [@soltysh](https://github.com/soltysh))
- The `kubectl wait --for=delete` command now ignores the not found error correctly. ([#96702](https://github.com/kubernetes/kubernetes/pull/96702), [@lingsamuel](https://github.com/lingsamuel))
- The `kubelet` now reports distinguishes log messages about certificate rotation for its client cert and server cert separately to make debugging problems with one or the other easier. ([#101252](https://github.com/kubernetes/kubernetes/pull/101252), [@smarterclayton](https://github.com/smarterclayton))
- The `serviceOwnsFrontendIP` shouldn't report error when the public IP doesn't match. ([#102516](https://github.com/kubernetes/kubernetes/pull/102516), [@nilo19](https://github.com/nilo19))
- The `system:aggregate-to-edit` role no longer includes write access to the Endpoints API. For new Kubernetes 1.22 clusters, the `edit` and `admin` roles will no longer include that access in newly created Kubernetes 1.22 clusters. This will have no affect on existing clusters upgrading to Kubernetes 1.22. To retain write access to Endpoints in the aggregated `edit` and `admin` roles for newly created 1.22 clusters, refer to https://github.com/kubernetes/website/pull/29025. ([#103704](https://github.com/kubernetes/kubernetes/pull/103704), [@robscott](https://github.com/robscott)) [SIG Auth and Network]
- The conformance tests:
  - Services should serve multiport endpoints from pods
  - Services should serve a basic endpoint from pods
  were only validating the API objects, not performing any validation on the actual Services implementation.
  Those tests now validate that the Services under test are able to forward traffic to the endpoints. ([#101709](https://github.com/kubernetes/kubernetes/pull/101709), [@aojea](https://github.com/aojea)) [SIG Network and Testing]
- The current behavior for Services that `IPFamilyPolicy` set as `PreferDualstack`. The current behavior when the cluster is upgraded to dual-stack is:
  - Services that have been set to IPFamilyPolicy = PreferDualstack will be upgraded when the service object is updated. e.g., when a user change a label.
  
  This behavior will change to:
  - Services that have been set  IPFamilyPolicy = PreferDualstack will not be upgraded when the service object is updated. User can still change policy, type etc and existing behaviors remain the same. ([#102898](https://github.com/kubernetes/kubernetes/pull/102898), [@khenidak](https://github.com/khenidak)) [SIG Network and Testing]
- The reason and message fields for pod status are no longer reset unless the phase also changes. ([#103785](https://github.com/kubernetes/kubernetes/pull/103785), [@smarterclayton](https://github.com/smarterclayton)) [SIG Node]
- Treat VSphere "File (vmdk path here) was not found" errors as success during volume deletion ([#92372](https://github.com/kubernetes/kubernetes/pull/92372), [@breunigs](https://github.com/breunigs)) [SIG Cloud Provider and Storage]
- Update `kube-proxy` base image `debian-iptables` to v1.6.2 to pickup [documentation](https://github.com/kubernetes/release/pull/2106)
  \n"- `debian-iptables`: select nft mode if ntf lines > legacy lines, matching [iptables-wrappers](https://github.com/kubernetes-sigs/iptables-wrappers/)" ([#102590](https://github.com/kubernetes/kubernetes/pull/102590), [@BenTheElder](https://github.com/BenTheElder))
- Update klog v2.9.0. ([#102332](https://github.com/kubernetes/kubernetes/pull/102332), [@pacoxu](https://github.com/pacoxu))
- Updated the Graceful Node Shutdown Pod termination reason and message.
  Updated the Graceful Node Shutdown Pod rejection reason and message. ([#102840](https://github.com/kubernetes/kubernetes/pull/102840), [@Kissy](https://github.com/Kissy))
- Updates dependency `sigs.k8s.io/structured-merge-diff` to v4.1.1. ([#100784](https://github.com/kubernetes/kubernetes/pull/100784), [@kevindelgado](https://github.com/kevindelgado))
- Updates hostprocess tests to specify user. ([#102965](https://github.com/kubernetes/kubernetes/pull/102965), [@jsturtevant](https://github.com/jsturtevant))
- Upgrades functionality of `kubectl kustomize` as described at
  https://github.com/kubernetes-sigs/kustomize/releases/tag/kustomize%2Fv4.2.0 ([#103419](https://github.com/kubernetes/kubernetes/pull/103419), [@natasha41575](https://github.com/natasha41575)) [SIG CLI]
- Upgrades functionality of `kubectl kustomize` as described at [kustomize/v4.1.2](https://github.com/kubernetes-sigs/kustomize/releases/tag/kustomize%2Fv4.1.2) ([#101120](https://github.com/kubernetes/kubernetes/pull/101120), [@monopole](https://github.com/monopole))
- Upgrading etcd: `kubeadm` upgrade etcd to 3.4.13-3 ([#100612](https://github.com/kubernetes/kubernetes/pull/100612), [@pacoxu](https://github.com/pacoxu))
- Use default timeout of 10s for Azure ACR credential provider. ([#100686](https://github.com/kubernetes/kubernetes/pull/100686), [@hasheddan](https://github.com/hasheddan)) [SIG Cloud Provider]
- We no longer allow the cluster operator to delete any suggested priority & fairness bootstrap configuration object. If a cluster operator removes a suggested configuration, it will be restored by the apiserver. ([#102067](https://github.com/kubernetes/kubernetes/pull/102067), [@tkashem](https://github.com/tkashem))
- When `DisableAcceleratorUsageMetrics` is set, do not collect accelerator metrics using cAdvisor. ([#101712](https://github.com/kubernetes/kubernetes/pull/101712), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev)) [SIG Instrumentation and Node]
- YAML documents separators ("---") can now be followed by whitespace and comments ("# ....") on the same line. This fixes a bug where documents starting with a comment after the separator were ignored. Other types of content on the same line will result in an error. ([#103457](https://github.com/kubernetes/kubernetes/pull/103457), [@codearky](https://github.com/codearky)) [SIG API Machinery]
- `oc describe quota` used has the same unit format as hard ([#102177](https://github.com/kubernetes/kubernetes/pull/102177), [@atiratree](https://github.com/atiratree)) [SIG CLI]

### Other (Cleanup or Flake)

- After the deprecation period,now the Kubelet's  `--chaos-chance` flag  are removed. ([#101057](https://github.com/kubernetes/kubernetes/pull/101057), [@wangyysde](https://github.com/wangyysde)) [SIG Node]
- Allow CSI drivers to just run offline expansion tests. ([#102665](https://github.com/kubernetes/kubernetes/pull/102665), [@gnufied](https://github.com/gnufied))
- Changed buildmode of non static Kubernetes binaries to produce position independent executables (PIE). ([#102323](https://github.com/kubernetes/kubernetes/pull/102323), [@saschagrunert](https://github.com/saschagrunert))
- Clarified the description of a test in the e2e suite that mentions "SCTP" but is
  actually intended to be testing the behavior of network plugins that don't
  implement SCTP. ([#102509](https://github.com/kubernetes/kubernetes/pull/102509), [@danwinship](https://github.com/danwinship))
- Client-go: reduce verbosity of `Starting/Stopping reflector` messages to 3 again. ([#102788](https://github.com/kubernetes/kubernetes/pull/102788), [@pohly](https://github.com/pohly))
- Disable log sampling when using json logging format. ([#102620](https://github.com/kubernetes/kubernetes/pull/102620), [@serathius](https://github.com/serathius))
- Exposes `WithCustomRoundTripper` method for specifying a middleware function for custom HTTP behaviour for the delegated auth clients. ([#99775](https://github.com/kubernetes/kubernetes/pull/99775), [@p0lyn0mial](https://github.com/p0lyn0mial))
- Fake clients now implement a `FakeClient` interface ([#100940](https://github.com/kubernetes/kubernetes/pull/100940), [@markusthoemmes](https://github.com/markusthoemmes)) [SIG API Machinery and Instrumentation]
- Featuregate `ServiceLoadBalancerClass` graduates to Beta and is enables by default. ([#103129](https://github.com/kubernetes/kubernetes/pull/103129), [@XudongLiuHarold](https://github.com/XudongLiuHarold))
- Improve func `ToSelectableFields`' performance for event. ([#102461](https://github.com/kubernetes/kubernetes/pull/102461), [@goodluckbot](https://github.com/goodluckbot))
- Increased `CSINodeIDMaxLength` from 128 bytes to 192 bytes. Prepare to increase the length limit to 256 bytes in 1.23 release. ([#101256](https://github.com/kubernetes/kubernetes/pull/101256), [@Jiawei0227](https://github.com/Jiawei0227))
- JSON logging now supports having information about source code location in the logging format, source code information is available under the key "caller". ([#102437](https://github.com/kubernetes/kubernetes/pull/102437), [@MadhavJivrajani](https://github.com/MadhavJivrajani))
- Kubeadm: move the BootstrapToken* API and related utilities from v1beta3 to a separate API group/version - bootstraptoken/v1. ([#102964](https://github.com/kubernetes/kubernetes/pull/102964), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: the `CriticalAddonsOnly` toleration has been removed from `kube-proxy` DaemonSet ([#101966](https://github.com/kubernetes/kubernetes/pull/101966), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Metrics Server updated to use `0.4.4` image that doesn't depend on deprecated `authorization.k8s.io/v1beta1` subjectaccessreviews API version. ([#101477](https://github.com/kubernetes/kubernetes/pull/101477), [@x13n](https://github.com/x13n))
- Migrate `proxy/ipvs/proxier.go` logs to structured logging. ([#97796](https://github.com/kubernetes/kubernetes/pull/97796), [@JornShen](https://github.com/JornShen))
- Migrate `staging/src/k8s.io/apiserver/pkg/registry` logs to structured logging. ([#98287](https://github.com/kubernetes/kubernetes/pull/98287), [@lala123912](https://github.com/lala123912))
- Migrate some log messages to structured logging in `pkg/volume/plugins.go`. ([#101510](https://github.com/kubernetes/kubernetes/pull/101510), [@huchengze](https://github.com/huchengze))
- Migrate some log messages to structured logging in `pkg/volume/volume_linux.go`. ([#99566](https://github.com/kubernetes/kubernetes/pull/99566), [@huchengze](https://github.com/huchengze))
- Official binaries now include the golang generated build ID `buildid` instead of an empty string. ([#101411](https://github.com/kubernetes/kubernetes/pull/101411), [@saschagrunert](https://github.com/saschagrunert))
- Remove balanced attached node volumes feature. ([#102443](https://github.com/kubernetes/kubernetes/pull/102443), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla))
- Remove deprecated `--generator` flag from `kubectl autoscale`. ([#99900](https://github.com/kubernetes/kubernetes/pull/99900), [@MadhavJivrajani](https://github.com/MadhavJivrajani))
- Remove the deprecated flag `--generator` from `kubectl create deployment` command. ([#99915](https://github.com/kubernetes/kubernetes/pull/99915), [@BLasan](https://github.com/BLasan))
- Remove the duplicate packet import. ([#101187](https://github.com/kubernetes/kubernetes/pull/101187), [@chuntaochen](https://github.com/chuntaochen))
- Replace `go-bindata` with `//go:embed`. ([#99829](https://github.com/kubernetes/kubernetes/pull/99829), [@palnabarun](https://github.com/palnabarun))
- The `DynamicFakeClient` now exposes its tracker via a `Tracker()` function. ([#100085](https://github.com/kubernetes/kubernetes/pull/100085), [@markusthoemmes](https://github.com/markusthoemmes))
- The `VolumeSnapshotDataSource` feature gate that is GA since v1.20 is unconditionally enabled, and can no longer be specified via the `--feature-gates` argument. ([#101531](https://github.com/kubernetes/kubernetes/pull/101531), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG Storage]
- The deprecated `CRIContainerLogRotation` feature-gate has been removed, since the `CRIContainerLogRotation` feature graduated to GA in 1.21 and was unconditionally enabled. ([#101578](https://github.com/kubernetes/kubernetes/pull/101578), [@carlory](https://github.com/carlory))
- The deprecated `RootCAConfigMap` feature-gate has been removed, since the `RootCAConfigMap` feature graduated to GA in 1.21 and is unconditionally enabled. ([#101579](https://github.com/kubernetes/kubernetes/pull/101579), [@carlory](https://github.com/carlory))
- The deprecated `runAsGroup` feature-gate has been removed, since the `runAsGroup` feature graduated to GA in 1.21. ([#101581](https://github.com/kubernetes/kubernetes/pull/101581), [@carlory](https://github.com/carlory))
- The etcd client has been updated to 3.5.0; `github.com/golang/protobuf`, `google.golang.org/protobuf`, and `google.golang.org/grpc` have been updated to current versions. ([#100488](https://github.com/kubernetes/kubernetes/pull/100488), [@liggitt](https://github.com/liggitt))
- Update Azure Go SDK to v55.0.0. ([#102441](https://github.com/kubernetes/kubernetes/pull/102441), [@feiskyer](https://github.com/feiskyer))
- Update Azure Go SDK version to v53.1.0 ([#101357](https://github.com/kubernetes/kubernetes/pull/101357), [@feiskyer](https://github.com/feiskyer)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle and Instrumentation]
- Update CNI plugins to v0.9.1. ([#102328](https://github.com/kubernetes/kubernetes/pull/102328), [@lentzi90](https://github.com/lentzi90))
- Update Calico to v3.19.1. ([#102386](https://github.com/kubernetes/kubernetes/pull/102386), [@JornShen](https://github.com/JornShen))
- Update cri-tools dependency to `v1.21.0`. ([#100956](https://github.com/kubernetes/kubernetes/pull/100956), [@saschagrunert](https://github.com/saschagrunert))
- Update dep `google/gnostic` and `google/go-cmp` to v0.5.5 and updating transitive dependencies `protobuf`. ([#102783](https://github.com/kubernetes/kubernetes/pull/102783), [@mcbenjemaa](https://github.com/mcbenjemaa))
- Update golang.org/x/net to v0.0.0-20210520170846-37e1c6afe023 ([#103176](https://github.com/kubernetes/kubernetes/pull/103176), [@CaoDonghui123](https://github.com/CaoDonghui123)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Node and Storage]
- Updated command descriptions and examples for grammar and punctuation consistency. ([#103524](https://github.com/kubernetes/kubernetes/pull/103524), [@bergerhoffer](https://github.com/bergerhoffer)) [SIG Auth and CLI]
- Updated pause image to version 3.5, which now runs per default as pseudo user and group `65535:65535`. This does not have any effect on remote container runtimes like CRI-O and containerd, which setup the pod sandbox user and group on their own. ([#100292](https://github.com/kubernetes/kubernetes/pull/100292), [@saschagrunert](https://github.com/saschagrunert))
- Upgrade functionality of `kubectl kustomize` as described at [kustomize/v4.1.3](https://github.com/kubernetes-sigs/kustomize/releases/tag/kustomize%2Fv4.1.3). ([#102193](https://github.com/kubernetes/kubernetes/pull/102193), [@gautierdelorme](https://github.com/gautierdelorme))

## Dependencies

### Added
- github.com/antihax/optional: [v1.0.0](https://github.com/antihax/optional/tree/v1.0.0)
- github.com/benbjohnson/clock: [v1.0.3](https://github.com/benbjohnson/clock/tree/v1.0.3)
- github.com/bits-and-blooms/bitset: [v1.2.0](https://github.com/bits-and-blooms/bitset/tree/v1.2.0)
- github.com/certifi/gocertifi: [2c3bb06](https://github.com/certifi/gocertifi/tree/2c3bb06)
- github.com/checkpoint-restore/go-criu/v5: [v5.0.0](https://github.com/checkpoint-restore/go-criu/v5/tree/v5.0.0)
- github.com/cncf/udpa/go: [5459f2c](https://github.com/cncf/udpa/go/tree/5459f2c)
- github.com/cockroachdb/errors: [v1.2.4](https://github.com/cockroachdb/errors/tree/v1.2.4)
- github.com/cockroachdb/logtags: [eb05cc2](https://github.com/cockroachdb/logtags/tree/eb05cc2)
- github.com/coredns/caddy: [v1.1.0](https://github.com/coredns/caddy/tree/v1.1.0)
- github.com/felixge/httpsnoop: [v1.0.1](https://github.com/felixge/httpsnoop/tree/v1.0.1)
- github.com/frankban/quicktest: [v1.11.3](https://github.com/frankban/quicktest/tree/v1.11.3)
- github.com/getsentry/raven-go: [v0.2.0](https://github.com/getsentry/raven-go/tree/v0.2.0)
- github.com/go-kit/log: [v0.1.0](https://github.com/go-kit/log/tree/v0.1.0)
- github.com/gofrs/uuid: [v4.0.0+incompatible](https://github.com/gofrs/uuid/tree/v4.0.0)
- github.com/josharian/intern: [v1.0.0](https://github.com/josharian/intern/tree/v1.0.0)
- github.com/jpillora/backoff: [v1.0.0](https://github.com/jpillora/backoff/tree/v1.0.0)
- github.com/nxadm/tail: [v1.4.4](https://github.com/nxadm/tail/tree/v1.4.4)
- github.com/opentracing/opentracing-go: [v1.1.0](https://github.com/opentracing/opentracing-go/tree/v1.1.0)
- github.com/robfig/cron/v3: [v3.0.1](https://github.com/robfig/cron/v3/tree/v3.0.1)
- github.com/stoewer/go-strcase: [v1.2.0](https://github.com/stoewer/go-strcase/tree/v1.2.0)
- go.etcd.io/etcd/api/v3: v3.5.0
- go.etcd.io/etcd/client/pkg/v3: v3.5.0
- go.etcd.io/etcd/client/v2: v2.305.0
- go.etcd.io/etcd/client/v3: v3.5.0
- go.etcd.io/etcd/pkg/v3: v3.5.0
- go.etcd.io/etcd/raft/v3: v3.5.0
- go.etcd.io/etcd/server/v3: v3.5.0
- go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc: v0.20.0
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.20.0
- go.opentelemetry.io/contrib: v0.20.0
- go.opentelemetry.io/otel/exporters/otlp: v0.20.0
- go.opentelemetry.io/otel/metric: v0.20.0
- go.opentelemetry.io/otel/oteltest: v0.20.0
- go.opentelemetry.io/otel/sdk/export/metric: v0.20.0
- go.opentelemetry.io/otel/sdk/metric: v0.20.0
- go.opentelemetry.io/otel/sdk: v0.20.0
- go.opentelemetry.io/otel/trace: v0.20.0
- go.opentelemetry.io/otel: v0.20.0
- go.opentelemetry.io/proto/otlp: v0.7.0
- go.uber.org/goleak: v1.1.10

### Changed
- github.com/Azure/azure-sdk-for-go: [v43.0.0+incompatible → v55.0.0+incompatible](https://github.com/Azure/azure-sdk-for-go/compare/v43.0.0...v55.0.0)
- github.com/Azure/go-ansiterm: [d6e3b33 → d185dfc](https://github.com/Azure/go-ansiterm/compare/d6e3b33...d185dfc)
- github.com/Azure/go-autorest/autorest/adal: [v0.9.5 → v0.9.13](https://github.com/Azure/go-autorest/autorest/adal/compare/v0.9.5...v0.9.13)
- github.com/Azure/go-autorest/autorest/to: [v0.2.0 → v0.4.0](https://github.com/Azure/go-autorest/autorest/to/compare/v0.2.0...v0.4.0)
- github.com/Azure/go-autorest/autorest: [v0.11.12 → v0.11.18](https://github.com/Azure/go-autorest/autorest/compare/v0.11.12...v0.11.18)
- github.com/Azure/go-autorest/logger: [v0.2.0 → v0.2.1](https://github.com/Azure/go-autorest/logger/compare/v0.2.0...v0.2.1)
- github.com/alecthomas/units: [c3de453 → f65c72e](https://github.com/alecthomas/units/compare/c3de453...f65c72e)
- github.com/auth0/go-jwt-middleware: [5493cab → v1.0.1](https://github.com/auth0/go-jwt-middleware/compare/5493cab...v1.0.1)
- github.com/aws/aws-sdk-go: [v1.35.24 → v1.38.49](https://github.com/aws/aws-sdk-go/compare/v1.35.24...v1.38.49)
- github.com/cilium/ebpf: [v0.2.0 → v0.6.2](https://github.com/cilium/ebpf/compare/v0.2.0...v0.6.2)
- github.com/cockroachdb/datadriven: [80d97fb → bf6692d](https://github.com/cockroachdb/datadriven/compare/80d97fb...bf6692d)
- github.com/container-storage-interface/spec: [v1.3.0 → v1.5.0](https://github.com/container-storage-interface/spec/compare/v1.3.0...v1.5.0)
- github.com/containerd/console: [v1.0.1 → v1.0.2](https://github.com/containerd/console/compare/v1.0.1...v1.0.2)
- github.com/containernetworking/cni: [v0.8.0 → v0.8.1](https://github.com/containernetworking/cni/compare/v0.8.0...v0.8.1)
- github.com/coredns/corefile-migration: [v1.0.11 → v1.0.12](https://github.com/coredns/corefile-migration/compare/v1.0.11...v1.0.12)
- github.com/coreos/go-systemd/v22: [v22.1.0 → v22.3.2](https://github.com/coreos/go-systemd/v22/compare/v22.1.0...v22.3.2)
- github.com/envoyproxy/go-control-plane: [5f8ba28 → 668b12f](https://github.com/envoyproxy/go-control-plane/compare/5f8ba28...668b12f)
- github.com/evanphx/json-patch: [v4.9.0+incompatible → v4.11.0+incompatible](https://github.com/evanphx/json-patch/compare/v4.9.0...v4.11.0)
- github.com/form3tech-oss/jwt-go: [v3.2.2+incompatible → v3.2.3+incompatible](https://github.com/form3tech-oss/jwt-go/compare/v3.2.2...v3.2.3)
- github.com/go-logfmt/logfmt: [v0.4.0 → v0.5.0](https://github.com/go-logfmt/logfmt/compare/v0.4.0...v0.5.0)
- github.com/go-openapi/jsonpointer: [v0.19.3 → v0.19.5](https://github.com/go-openapi/jsonpointer/compare/v0.19.3...v0.19.5)
- github.com/go-openapi/jsonreference: [v0.19.3 → v0.19.5](https://github.com/go-openapi/jsonreference/compare/v0.19.3...v0.19.5)
- github.com/go-openapi/swag: [v0.19.5 → v0.19.14](https://github.com/go-openapi/swag/compare/v0.19.5...v0.19.14)
- github.com/godbus/dbus/v5: [v5.0.3 → v5.0.4](https://github.com/godbus/dbus/v5/compare/v5.0.3...v5.0.4)
- github.com/golang/groupcache: [8c9f03a → 41bb18b](https://github.com/golang/groupcache/compare/8c9f03a...41bb18b)
- github.com/golang/protobuf: [v1.4.3 → v1.5.2](https://github.com/golang/protobuf/compare/v1.4.3...v1.5.2)
- github.com/google/btree: [v1.0.0 → v1.0.1](https://github.com/google/btree/compare/v1.0.0...v1.0.1)
- github.com/google/cadvisor: [v0.39.0 → v0.39.2](https://github.com/google/cadvisor/compare/v0.39.0...v0.39.2)
- github.com/google/go-cmp: [v0.5.2 → v0.5.5](https://github.com/google/go-cmp/compare/v0.5.2...v0.5.5)
- github.com/googleapis/gnostic: [v0.4.1 → v0.5.5](https://github.com/googleapis/gnostic/compare/v0.4.1...v0.5.5)
- github.com/gopherjs/gopherjs: [0766667 → fce0ec3](https://github.com/gopherjs/gopherjs/compare/0766667...fce0ec3)
- github.com/grpc-ecosystem/go-grpc-middleware: [f849b54 → v1.3.0](https://github.com/grpc-ecosystem/go-grpc-middleware/compare/f849b54...v1.3.0)
- github.com/grpc-ecosystem/grpc-gateway: [v1.9.5 → v1.16.0](https://github.com/grpc-ecosystem/grpc-gateway/compare/v1.9.5...v1.16.0)
- github.com/heketi/heketi: [v10.2.0+incompatible → v10.3.0+incompatible](https://github.com/heketi/heketi/compare/v10.2.0...v10.3.0)
- github.com/jonboulle/clockwork: [v0.1.0 → v0.2.2](https://github.com/jonboulle/clockwork/compare/v0.1.0...v0.2.2)
- github.com/json-iterator/go: [v1.1.10 → v1.1.11](https://github.com/json-iterator/go/compare/v1.1.10...v1.1.11)
- github.com/julienschmidt/httprouter: [v1.2.0 → v1.3.0](https://github.com/julienschmidt/httprouter/compare/v1.2.0...v1.3.0)
- github.com/kr/pretty: [v0.2.0 → v0.2.1](https://github.com/kr/pretty/compare/v0.2.0...v0.2.1)
- github.com/mailru/easyjson: [v0.7.0 → v0.7.6](https://github.com/mailru/easyjson/compare/v0.7.0...v0.7.6)
- github.com/mattn/go-isatty: [v0.0.4 → v0.0.3](https://github.com/mattn/go-isatty/compare/v0.0.4...v0.0.3)
- github.com/miekg/dns: [v1.1.35 → v1.0.14](https://github.com/miekg/dns/compare/v1.1.35...v1.0.14)
- github.com/moby/sys/mountinfo: [v0.4.0 → v0.4.1](https://github.com/moby/sys/mountinfo/compare/v0.4.0...v0.4.1)
- github.com/moby/term: [df9cb8a → 9d4ed18](https://github.com/moby/term/compare/df9cb8a...9d4ed18)
- github.com/mwitkow/go-conntrack: [cc309e4 → 2f06839](https://github.com/mwitkow/go-conntrack/compare/cc309e4...2f06839)
- github.com/onsi/ginkgo: [v1.11.0 → v1.14.0](https://github.com/onsi/ginkgo/compare/v1.11.0...v1.14.0)
- github.com/onsi/gomega: [v1.7.0 → v1.10.1](https://github.com/onsi/gomega/compare/v1.7.0...v1.10.1)
- github.com/opencontainers/runc: [v1.0.0-rc93 → v1.0.1](https://github.com/opencontainers/runc/compare/v1.0.0-rc93...v1.0.1)
- github.com/opencontainers/runtime-spec: [e6143ca → 1c3f411](https://github.com/opencontainers/runtime-spec/compare/e6143ca...1c3f411)
- github.com/opencontainers/selinux: [v1.8.0 → v1.8.2](https://github.com/opencontainers/selinux/compare/v1.8.0...v1.8.2)
- github.com/prometheus/client_golang: [v1.7.1 → v1.11.0](https://github.com/prometheus/client_golang/compare/v1.7.1...v1.11.0)
- github.com/prometheus/common: [v0.10.0 → v0.26.0](https://github.com/prometheus/common/compare/v0.10.0...v0.26.0)
- github.com/prometheus/procfs: [v0.2.0 → v0.6.0](https://github.com/prometheus/procfs/compare/v0.2.0...v0.6.0)
- github.com/rogpeppe/fastuuid: [6724a57 → v1.2.0](https://github.com/rogpeppe/fastuuid/compare/6724a57...v1.2.0)
- github.com/sirupsen/logrus: [v1.7.0 → v1.8.1](https://github.com/sirupsen/logrus/compare/v1.7.0...v1.8.1)
- github.com/smartystreets/assertions: [b2de0cb → v1.1.0](https://github.com/smartystreets/assertions/compare/b2de0cb...v1.1.0)
- github.com/soheilhy/cmux: [v0.1.4 → v0.1.5](https://github.com/soheilhy/cmux/compare/v0.1.4...v0.1.5)
- github.com/spf13/cobra: [v1.1.1 → v1.1.3](https://github.com/spf13/cobra/compare/v1.1.1...v1.1.3)
- github.com/spf13/jwalterweatherman: [v1.1.0 → v1.0.0](https://github.com/spf13/jwalterweatherman/compare/v1.1.0...v1.0.0)
- github.com/stretchr/testify: [v1.6.1 → v1.7.0](https://github.com/stretchr/testify/compare/v1.6.1...v1.7.0)
- github.com/tmc/grpc-websocket-proxy: [0ad062e → e5319fd](https://github.com/tmc/grpc-websocket-proxy/compare/0ad062e...e5319fd)
- github.com/yuin/goldmark: [v1.2.1 → v1.3.5](https://github.com/yuin/goldmark/compare/v1.2.1...v1.3.5)
- go.etcd.io/bbolt: v1.3.5 → v1.3.6
- go.uber.org/atomic: v1.4.0 → v1.7.0
- go.uber.org/multierr: v1.1.0 → v1.6.0
- go.uber.org/zap: v1.10.0 → v1.17.0
- golang.org/x/lint: 738671d → 6edffad
- golang.org/x/mod: ce943fd → v0.4.2
- golang.org/x/net: 3d97a24 → 37e1c6a
- golang.org/x/sync: 67f06af → 036812b
- golang.org/x/sys: a50acf3 → 59db8d7
- golang.org/x/text: v0.3.4 → v0.3.6
- golang.org/x/time: f8bda1e → 1f47c86
- golang.org/x/tools: v0.1.0 → v0.1.2
- google.golang.org/genproto: 8816d57 → f16073e
- google.golang.org/grpc: v1.27.1 → v1.38.0
- google.golang.org/protobuf: v1.25.0 → v1.26.0
- gopkg.in/yaml.v3: 9f266ea → 496545a
- k8s.io/klog/v2: v2.8.0 → v2.9.0
- k8s.io/kube-openapi: 591a79e → 9528897
- k8s.io/system-validators: v1.4.0 → v1.5.0
- k8s.io/utils: 67b214c → 4b05e18
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.15 → v0.0.22
- sigs.k8s.io/kustomize/api: v0.8.5 → v0.8.11
- sigs.k8s.io/kustomize/cmd/config: v0.9.7 → v0.9.13
- sigs.k8s.io/kustomize/kustomize/v4: v4.0.5 → v4.2.0
- sigs.k8s.io/kustomize/kyaml: v0.10.15 → v0.11.0
- sigs.k8s.io/structured-merge-diff/v4: v4.1.0 → v4.1.2

### Removed
- github.com/agnivade/levenshtein: [v1.0.1](https://github.com/agnivade/levenshtein/tree/v1.0.1)
- github.com/alecthomas/template: [fb15b89](https://github.com/alecthomas/template/tree/fb15b89)
- github.com/andreyvit/diff: [c7f18ee](https://github.com/andreyvit/diff/tree/c7f18ee)
- github.com/bifurcation/mint: [93c51c6](https://github.com/bifurcation/mint/tree/93c51c6)
- github.com/caddyserver/caddy: [v1.0.3](https://github.com/caddyserver/caddy/tree/v1.0.3)
- github.com/cenkalti/backoff: [v2.1.1+incompatible](https://github.com/cenkalti/backoff/tree/v2.1.1)
- github.com/checkpoint-restore/go-criu/v4: [v4.1.0](https://github.com/checkpoint-restore/go-criu/v4/tree/v4.1.0)
- github.com/cheekybits/genny: [9127e81](https://github.com/cheekybits/genny/tree/9127e81)
- github.com/go-acme/lego: [v2.5.0+incompatible](https://github.com/go-acme/lego/tree/v2.5.0)
- github.com/go-bindata/go-bindata: [v3.1.1+incompatible](https://github.com/go-bindata/go-bindata/tree/v3.1.1)
- github.com/go-openapi/analysis: [v0.19.5](https://github.com/go-openapi/analysis/tree/v0.19.5)
- github.com/go-openapi/errors: [v0.19.2](https://github.com/go-openapi/errors/tree/v0.19.2)
- github.com/go-openapi/loads: [v0.19.4](https://github.com/go-openapi/loads/tree/v0.19.4)
- github.com/go-openapi/runtime: [v0.19.4](https://github.com/go-openapi/runtime/tree/v0.19.4)
- github.com/go-openapi/spec: [v0.19.5](https://github.com/go-openapi/spec/tree/v0.19.5)
- github.com/go-openapi/strfmt: [v0.19.5](https://github.com/go-openapi/strfmt/tree/v0.19.5)
- github.com/go-openapi/validate: [v0.19.8](https://github.com/go-openapi/validate/tree/v0.19.8)
- github.com/gobuffalo/here: [v0.6.0](https://github.com/gobuffalo/here/tree/v0.6.0)
- github.com/hpcloud/tail: [v1.0.0](https://github.com/hpcloud/tail/tree/v1.0.0)
- github.com/jimstudt/http-authentication: [3eca13d](https://github.com/jimstudt/http-authentication/tree/3eca13d)
- github.com/klauspost/cpuid: [v1.2.0](https://github.com/klauspost/cpuid/tree/v1.2.0)
- github.com/kr/logfmt: [b84e30a](https://github.com/kr/logfmt/tree/b84e30a)
- github.com/kylelemons/godebug: [d65d576](https://github.com/kylelemons/godebug/tree/d65d576)
- github.com/lucas-clemente/aes12: [cd47fb3](https://github.com/lucas-clemente/aes12/tree/cd47fb3)
- github.com/lucas-clemente/quic-clients: [v0.1.0](https://github.com/lucas-clemente/quic-clients/tree/v0.1.0)
- github.com/lucas-clemente/quic-go-certificates: [d2f8652](https://github.com/lucas-clemente/quic-go-certificates/tree/d2f8652)
- github.com/lucas-clemente/quic-go: [v0.10.2](https://github.com/lucas-clemente/quic-go/tree/v0.10.2)
- github.com/markbates/pkger: [v0.17.1](https://github.com/markbates/pkger/tree/v0.17.1)
- github.com/marten-seemann/qtls: [v0.2.3](https://github.com/marten-seemann/qtls/tree/v0.2.3)
- github.com/mholt/certmagic: [6a42ef9](https://github.com/mholt/certmagic/tree/6a42ef9)
- github.com/naoina/go-stringutil: [v0.1.0](https://github.com/naoina/go-stringutil/tree/v0.1.0)
- github.com/naoina/toml: [v0.1.1](https://github.com/naoina/toml/tree/v0.1.1)
- github.com/robfig/cron: [v1.1.0](https://github.com/robfig/cron/tree/v1.1.0)
- github.com/satori/go.uuid: [v1.2.0](https://github.com/satori/go.uuid/tree/v1.2.0)
- github.com/thecodeteam/goscaleio: [v0.1.0](https://github.com/thecodeteam/goscaleio/tree/v0.1.0)
- github.com/tidwall/pretty: [v1.0.0](https://github.com/tidwall/pretty/tree/v1.0.0)
- github.com/vektah/gqlparser: [v1.1.2](https://github.com/vektah/gqlparser/tree/v1.1.2)
- github.com/willf/bitset: [v1.1.11](https://github.com/willf/bitset/tree/v1.1.11)
- go.etcd.io/etcd: dd1b699
- go.mongodb.org/mongo-driver: v1.1.2
- gopkg.in/cheggaaa/pb.v1: v1.0.25
- gopkg.in/fsnotify.v1: v1.4.7
- gopkg.in/mcuadros/go-syslog.v2: v2.2.1
- gopkg.in/resty.v1: v1.12.0
- k8s.io/heapster: v1.2.0-beta.1



# v1.22.0-rc.0


## Downloads for v1.22.0-rc.0

### Source Code

filename | sha512 hash
-------- | -----------
[kubernetes.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes.tar.gz) | 94c906dddb353fd3de30cb22c587efd0a683d2865cfad64fdf04a63132cd79be700cdb63ca8752a4af29d192ce151dde83c07192a4d3af73d28f89c08c246236
[kubernetes-src.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-src.tar.gz) | ad8a80e925779cfdc0b31021aaac8931ecdb419a7b319d264ae6cdc7fe5f7f63e3b999e83e26074f94ea429b4c7c9345e221653b7b7c18984e29d488f66f7da0

### Client Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-client-darwin-amd64.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-client-darwin-amd64.tar.gz) | 8e0133fabe29e00c107d7d8973bbbbac9fc30abc9bfee8e06595910c4e5f58c5e37bc069adcfef306efa68efc121ed2818cce24e585a791f753ca0c86ee62f7e
[kubernetes-client-darwin-arm64.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-client-darwin-arm64.tar.gz) | 55baf22e323ef47ea8241b0f98bd23818ad54c87498458c976fdd619362297c5739ed25751111f6fe83fa7ba06480e0011e67f785d84e996e9480def250620a7
[kubernetes-client-linux-386.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-client-linux-386.tar.gz) | e53416e5542278103457bc69f95998fa1cd5ecb5b0d5f2f144096d717d0db50bc4b0ccc43f220c10e21867e6a3dcc1dd21c1a329291cfdff90c7391c106c273f
[kubernetes-client-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-client-linux-amd64.tar.gz) | 7089d702765c474cbf365fd92190d69e3e461b8d880ebceb20d3e1bddd112f387f74b1c85f8c90639984e7fb033539d89ef1f8bb6b753dbe87fd593cae7e489b
[kubernetes-client-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-client-linux-arm.tar.gz) | 852349deda1eaf159a5f9d1dab0db115c564110130ff4e39f0252583f0eaa5a3ffc3c24214e93d36508f33173a1b5952e377e87c6b39be456184dcc856b83ceb
[kubernetes-client-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-client-linux-arm64.tar.gz) | 42583f79eab40f2e840e81f5675a3eaed670cf2ea02ab35eda1c7fcc559cc7d730fd15b8e40ca5a138d9a1a29b52e8644d40c5cc7dfc48a366bc7f61049aad4d
[kubernetes-client-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-client-linux-ppc64le.tar.gz) | 4f3580b128c994027fdcac5c57cc0fe17993ef3ff156a782b5d9ca081988bf9fe07e1b0ee8f8beedf3e272c0c579cad6ca198a1787aa502df94d7053a9699815
[kubernetes-client-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-client-linux-s390x.tar.gz) | f63fa853ba3b605a9b2081eac1336f9400572bf0435bd2f8dd9f1ad72045ef522dbcf5bec516944d6288cb91e8497dc8a627c272fc2504404a5553d7058df92b
[kubernetes-client-windows-386.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-client-windows-386.tar.gz) | 752ad0a3028c207044a4d882fe603fd325107a508e567fe3bf939cf06d67d2d1c523031b3197cf4a3edfa7b2c36b37e9943e5b3567149e36669bae6bc99b1cce
[kubernetes-client-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-client-windows-amd64.tar.gz) | 88d1dd951eaeeda51de9c728036dfb1ee428ae5944b657ecdc702a90366031bdd75ce30092f179abd861c211873aa52c30386bc1975e8a92a24f31440b8229ad

### Server Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-server-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-server-linux-amd64.tar.gz) | 9372ea845dafc41949bbfa4c45c197f282b61950eefa19bb7486c58991ce76c7f7eeb1e97a38726f4b627d367614d00cbac24a4f46d13f4530fc0fa84cc05423
[kubernetes-server-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-server-linux-arm.tar.gz) | 9c8be6ca9f3b27ddbf1aff9cb17925656ac81927f76a68784fb898886887aebce9011a5d8726ecbd2d82a6b18819d0df0ff9eb1fb62c18035e12430458d0c777
[kubernetes-server-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-server-linux-arm64.tar.gz) | 9b72565025216f3dd8859ed3b56cbb29d1ed2b34d470fc959b9b4f5ccc34cdcee1e590ebf11979c0166e179ada27612ee8a3943bf9462146caa752e3e670fcc3
[kubernetes-server-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-server-linux-ppc64le.tar.gz) | 3a8674858cb4c5dfe5e76fa156e38c5f7684711bcfad1e1b8f3c48b4350eb446bd3032b21d2a6b1da5be23b2d6d8565096ad091c2af8453676388fafcda10e57
[kubernetes-server-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-server-linux-s390x.tar.gz) | fac5de2ea843490bd20fb4e91ccccd692fa42593baee7d396322a17b79d781cecf34a57fe17617173fb27cfe5778236957176b62deee7afe721793e84eb7d0a3

### Node Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-node-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-node-linux-amd64.tar.gz) | 4ae063266207869ba4a6f9290b75785a4e555f3c8ebc7856fcb57b14a3b29993afc106e68bc91b918391830ff1585b8772040e3a8a55ffbef7b3b2968f286971
[kubernetes-node-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-node-linux-arm.tar.gz) | ec2a1769e69882ba7bbdbd51d35c62b839e39eb2d49db217ceb2c9e8bb2673e45b2946bdd020b859b1272e3318518d4706dbd1e757adf51c23c1959dcfc80c0e
[kubernetes-node-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-node-linux-arm64.tar.gz) | de98240a500ac4d1fb739eca31326c283323c884419eacfac92632a86273a19ac69773560c310218cd8f4f3546db4ece0de56961e621e8004188e04b584d34ba
[kubernetes-node-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-node-linux-ppc64le.tar.gz) | c12674ddcd0a5f74d00873fc2682a60379353ec49cb54bb14e13fec361ceeaf0bb46b625fb6994b05bb9ae37a507d1c2bff0c6584619ff0658529742d2112459
[kubernetes-node-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-node-linux-s390x.tar.gz) | b1e5780a211ff1e2651bca78bfc611eb1cfab2fddf74eb0a6ea48746db4db07b519d88a1b4bc3e7af43132062c760af7244800b3f42ecc66a6ded8dbdf31ac72
[kubernetes-node-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-rc.0/kubernetes-node-windows-amd64.tar.gz) | 4c0a3b0731c2110d0fce1e0d67e30752618555840a23f133cab8fa928e863d4b024cdd03e3f0e24c84192d99a939ed4b6e152b8ac58c3ce0885933e04ccce682

## Changelog since v1.22.0-beta.2

## Changes by Kind

### API Change

- Revert addition of Add PersistentVolumeClaimDeletePoilcy to StatefulSet API. ([#103747](https://github.com/kubernetes/kubernetes/pull/103747), [@mattcary](https://github.com/mattcary)) [SIG API Machinery and Apps]

### Bug or Regression

- Aggregated roles no longer include write access to EndpointSlices. This rolls back part of a change that was introduced earlier in the Kubernetes 1.22 cycle. ([#103703](https://github.com/kubernetes/kubernetes/pull/103703), [@robscott](https://github.com/robscott)) [SIG Auth and Network]
- Fix winkernel kube-proxy to only use dual stack when host and networking supports it ([#101047](https://github.com/kubernetes/kubernetes/pull/101047), [@jsturtevant](https://github.com/jsturtevant)) [SIG Network and Windows]
- The `system:aggregate-to-edit` role no longer includes write access to the Endpoints API. For new Kubernetes 1.22 clusters, the `edit` and `admin` roles will no longer include that access in newly created Kubernetes 1.22 clusters. This will have no affect on existing clusters upgrading to Kubernetes 1.22. To retain write access to Endpoints in the aggregated `edit` and `admin` roles for newly created 1.22 clusters, refer to https://github.com/kubernetes/website/pull/29025. ([#103704](https://github.com/kubernetes/kubernetes/pull/103704), [@robscott](https://github.com/robscott)) [SIG Auth and Network]
- The reason and message fields for pod status are no longer reset unless the phase also changes. ([#103785](https://github.com/kubernetes/kubernetes/pull/103785), [@smarterclayton](https://github.com/smarterclayton)) [SIG Node]

## Dependencies

### Added
_Nothing has changed._

### Changed
_Nothing has changed._

### Removed
_Nothing has changed._



# v1.22.0-beta.2


## Downloads for v1.22.0-beta.2

### Source Code

filename | sha512 hash
-------- | -----------
[kubernetes.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes.tar.gz) | 564ee1ca1e345ff46cf4085a9b3c2084b6b53c0094837671261e798b18d78e9eef391f1f1d30081f787897f8dcdde9317f1d283ceb75c93b889aab7aed505ce1
[kubernetes-src.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-src.tar.gz) | 13353ed167e7671c16e1abda3f88082e974c863f865dd6bb09ef1fdfbb097f6e71f7dc5a8f605bbe2c1e68de15cd339b76caf5d69234211f1720c2f3da1d8eba

### Client Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-client-darwin-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-client-darwin-amd64.tar.gz) | 7cb2d30bbd50d44df1fb0275f8ea8330654311cc566b89dcf2bf8d3ee9d045728a49c7b61a72eeb7ce13e22d604cff367a28fc89d1b8940fc9cc1e014866b354
[kubernetes-client-darwin-arm64.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-client-darwin-arm64.tar.gz) | 7810a1076e416d1a1721c0c419403f454b0ae06194e5ef9edabfd0d03f515dfb7194a300b7ace200b273509c7521645a3fe3867aa54c9827d1d54480e6df3b4e
[kubernetes-client-linux-386.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-client-linux-386.tar.gz) | 67daedc35190364b9f7af31672ceceb17b0770979d19cfa291aef4cb781eb2b8edd082ac613443850d41e5122d82fa320e9d72d6dc5ee32b6667f1d2e4bba51f
[kubernetes-client-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-client-linux-amd64.tar.gz) | a11cab14ad46c57596e73fdd6e25763df8a36be324cde275344db5e05b05bc0cb9b92bd17788df0ce949ea6ea3f73406d873f5f2f171f3fb25b05bd5dc456e72
[kubernetes-client-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-client-linux-arm.tar.gz) | 34e89e0bf6ebd0051f5162342f574aebf3f716c4902f747e4f9de38169be70f1298f4ee7cca2efde60c73d6e017d4435e20735bb255ca813c8dfc9c3b9afdb68
[kubernetes-client-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-client-linux-arm64.tar.gz) | 6f581c5eae7808ea2847c56d77dd9ddd403ec7d65386c0e20665c00bd63e82c2b6982bb18a7358d52a56eeee29fd5409835e7abd05c198eb80e5ad1a45612860
[kubernetes-client-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-client-linux-ppc64le.tar.gz) | bdfd91a51f2341a0afbe6aa5cb0a08b31adad179117eef4d720bb616bf33a58ab99cc042661c2087bc8ecf6ca79ad9647ea5000ce7f72ae26157b347652b7b14
[kubernetes-client-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-client-linux-s390x.tar.gz) | 4acee93c50011e1c7c99da7cb5cab07f18301fca71c4855c281bca6547a81984ab6e591429a948deae94c0dc379f5cb4e877a7dd35d9f1a0221d2c7e945de476
[kubernetes-client-windows-386.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-client-windows-386.tar.gz) | 1286fc074ac03e6f0e4080f408e696881741627387133cea8c75d5b2824d88b53610b476e7b0d71a1a4b61afcf1c9f1e173e2b3f4ea497a9a85e7d7f326430f9
[kubernetes-client-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-client-windows-amd64.tar.gz) | d2037202680a56f47eb2944493f54677f8da6d581b39930f02dfea526234dc2ea8c7a56ead12298b19be5ff613c54b1b35b65d8e716fcd155ddc038d57ac06be

### Server Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-server-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-server-linux-amd64.tar.gz) | ea6c13e9c2f76ad6dee59be352b694537fd2e647ba9e38c0c0655c838eed99b37c575eb0c5836eb7f35931314fd7f950623f7c73a8fba328bebf7fa345d45516
[kubernetes-server-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-server-linux-arm.tar.gz) | 1a421d745853bd03519f14dbeca631036c4e95fb658872ed69bd0824126479b91a7229384f4c178cfac57dc78546c989d1d49a5dc7d3f5c773dc5550cc0e9ad3
[kubernetes-server-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-server-linux-arm64.tar.gz) | e2cc43e00134963587e947b5400a03ee56a76786551eeef3f2b9f0e8345ed337b0691acf98f1c0279af761fae5e830073f0ffb4c817c288b28cc9b581cad47e2
[kubernetes-server-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-server-linux-ppc64le.tar.gz) | b2a58a5b0621334227f95607b87d3d3ebf5ea75ff91bd0617ddc11507ecf9a98376a9bff3cfa03d852ada65a7ddbbd03802ae380b7d50cac1fa6f752c5cb852e
[kubernetes-server-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-server-linux-s390x.tar.gz) | 5fe52138d593f835ff2444b5826a4f3c5ebad0f07ca30044cfa22d9b40e5de06454e803fa82e7a070bffc4070fb8de6e2fa4db7f6be1f3b07ec9bca28dec80f7

### Node Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-node-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-node-linux-amd64.tar.gz) | 0c68a611493f13fbaef2848d13558472e0a9f23cfe439f3767ab9e77b6fcbfb26c988a2343688fc0d21b9decabd6d2d414ae45a6d55d11f73beadc4d659f358a
[kubernetes-node-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-node-linux-arm.tar.gz) | a6222cbaf9e54f9ccf2b95e97178dc86de9b45d5d2bd4ecc64a7b47761ae03f643612dd03d5958fdfe37116ac74293d52980c15a320cc7dec3179d938b5b819f
[kubernetes-node-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-node-linux-arm64.tar.gz) | a4117aeba3171a1ea776efa7ed7ba1d1aca57e1f3860d816c4daa7e0efee3aba529ba355b6726f747b65efbd519c20f3bc0a5bd13b60d4453daa811a049fa8c9
[kubernetes-node-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-node-linux-ppc64le.tar.gz) | b1dbaa70fc88845297dbaba65b0df212972be4871c94f961922f95c2e411c4ae2aa065ce8912604480950bdb0657586427d7835aedf5eee7cb4b30059c47f6bf
[kubernetes-node-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-node-linux-s390x.tar.gz) | ba82a3bfe6b3d6f351a711597f5e3b42fddee69c681ff1e67009ef8a37120c0f4b43b5231b39eb0b5c6c026d9c7ef632b6e0431beb4f9ade7f9ff30cd2264b81
[kubernetes-node-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.2/kubernetes-node-windows-amd64.tar.gz) | bd2b657f91ea4f66eee0fcc72bbb524548fbe86c93aa3b89ead332012e19a5e109c7b3a34186b72d603756724fabf41af2fa14b2d2513343b2c47ea5d07343bb

## Changelog since v1.22.0-beta.1

## Changes by Kind

### API Change

- Added field .status.uncountedTerminatedPods to the Job resource. This field is used by the job controller to keep track of finished pods before adding them to the Job status counters.
  
  Pods created by the job controller get the finalizer batch.kubernetes.io/job-tracking
  
  Jobs that are tracked using this mechanism get the annotation batch.kubernetes.io/job-tracking. This is a temporary measure. Two releases after this feature graduates to beta, the annotation won't be added to Jobs anymore. ([#98817](https://github.com/kubernetes/kubernetes/pull/98817), [@alculquicondor](https://github.com/alculquicondor)) [SIG API Machinery, Apps, Auth and CLI]
- Ephemeral containers are now allowed to configure a securityContext that differs from that of the Pod.
  
  Cluster administrators should ensure that security policy controllers support EphemeralContainers before enabling this feature in clusters. ([#99023](https://github.com/kubernetes/kubernetes/pull/99023), [@verb](https://github.com/verb)) [SIG API Machinery, Apps, Auth and Node]
- If someone had the ProbeTerminationGracePeriod alpha feature enabled in 1.21, they should update/delete any workloads/pods with probe terminationGracePeriods < 1 before upgrading ([#103245](https://github.com/kubernetes/kubernetes/pull/103245), [@wzshiming](https://github.com/wzshiming)) [SIG Apps and Node]
- Introducing Memory QoS support with cgroups v2 (Alpha)
  The MemoryQoS feature is now in Alpha. This allows kubelet running with cgroups v2 to set memory QoS at container, pod and QoS level to protect and guarantee better memory quality. This feature can be enabled through feature gate MemoryQoS. ([#102970](https://github.com/kubernetes/kubernetes/pull/102970), [@borgerli](https://github.com/borgerli)) [SIG Node and Storage]
- NodeSwapEnabled feature flag was renamed to NodeSwap
  
  The flag was only available in the 1.22.0-beta.1 release, and the new flag should be used going forward. ([#103553](https://github.com/kubernetes/kubernetes/pull/103553), [@ehashman](https://github.com/ehashman)) [SIG Node]
- Removed the feature flag for probe-level termination grace period from Kubelet. If a user wants to disable this feature on already created pods, they will have to delete and recreate the pods. ([#103168](https://github.com/kubernetes/kubernetes/pull/103168), [@raisaat](https://github.com/raisaat)) [SIG Apps and Node]
- Track Job completion through status and Pod finalizers, removing dependency on Pod tombstones. ([#98238](https://github.com/kubernetes/kubernetes/pull/98238), [@alculquicondor](https://github.com/alculquicondor)) [SIG API Machinery, Apps, Auth and Testing]
- When using `kubectl replace` (or the equivalent API call) on a Service, the caller no longer needs to do a read-modify-write cycle to fetch the allocated values for `.spec.clusterIP` and `.spec.ports[].nodePort`.  Instead the API server will automatically carry these forward from the original object when the new object does not specify them. ([#103532](https://github.com/kubernetes/kubernetes/pull/103532), [@thockin](https://github.com/thockin)) [SIG Apps and Network]

### Feature

- APIServerTracing now collects spans from etcd client calls, and propagates context to etcd. ([#103216](https://github.com/kubernetes/kubernetes/pull/103216), [@dashpole](https://github.com/dashpole)) [SIG API Machinery, Cloud Provider and Instrumentation]
- APIServerTracing now collects spans from outgoing requests to admission webhooks. ([#103601](https://github.com/kubernetes/kubernetes/pull/103601), [@dashpole](https://github.com/dashpole)) [SIG API Machinery]
- Kubernetes is now built with Golang 1.16.6 ([#103669](https://github.com/kubernetes/kubernetes/pull/103669), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Leader Migration for controller managers graduated to beta. ([#103533](https://github.com/kubernetes/kubernetes/pull/103533), [@jiahuif](https://github.com/jiahuif)) [SIG API Machinery and Cloud Provider]
- Promote EndpointSliceTerminatingCondition to Beta. This enables the 'terminating' and 'serving' conditions for EndpointSlice by default. ([#103596](https://github.com/kubernetes/kubernetes/pull/103596), [@andrewsykim](https://github.com/andrewsykim)) [SIG Network and Testing]

### Bug or Regression

- Fix "verb" and "scope" tag in apiserver_longrunning_gauge and apiserver_request_terminations_total  such that a GET namespaced object by name call will have verb "GET" instead of "LIST", and scope "resource" instead of "namespace" ([#103565](https://github.com/kubernetes/kubernetes/pull/103565), [@zhan849](https://github.com/zhan849)) [SIG API Machinery and Instrumentation]
- Fix a number of race conditions in the kubelet when pods are starting up or shutting down that might cause pods to take a long time to shut down. ([#102344](https://github.com/kubernetes/kubernetes/pull/102344), [@smarterclayton](https://github.com/smarterclayton)) [SIG Apps, Node, Storage and Testing]
- Fix the code is leaking the defaulting between unrelated pod instances. ([#103284](https://github.com/kubernetes/kubernetes/pull/103284), [@kebe7jun](https://github.com/kebe7jun)) [SIG CLI]
- Fix: return empty VMAS name if using standalone VM ([#103470](https://github.com/kubernetes/kubernetes/pull/103470), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- GCE Windows clusters have their TCP/IP parameters are set to GCE's recommended values. ([#103057](https://github.com/kubernetes/kubernetes/pull/103057), [@jeremyje](https://github.com/jeremyje)) [SIG Cloud Provider and Windows]
- Graceful termination will now be honored when deleting a collection of pods. ([#100101](https://github.com/kubernetes/kubernetes/pull/100101), [@deads2k](https://github.com/deads2k)) [SIG API Machinery, Node and Testing]
- YAML documents separators ("---") can now be followed by whitespace and comments ("# ....") on the same line. This fixes a bug where documents starting with a comment after the separator were ignored. Other types of content on the same line will result in an error. ([#103457](https://github.com/kubernetes/kubernetes/pull/103457), [@codearky](https://github.com/codearky)) [SIG API Machinery]
- `oc describe quota` used has the same unit format as hard ([#102177](https://github.com/kubernetes/kubernetes/pull/102177), [@atiratree](https://github.com/atiratree)) [SIG CLI]

## Dependencies

### Added
_Nothing has changed._

### Changed
_Nothing has changed._

### Removed
_Nothing has changed._



# v1.22.0-beta.1


## Downloads for v1.22.0-beta.1

### Source Code

filename | sha512 hash
-------- | -----------
[kubernetes.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes.tar.gz) | 87201338ac3506b65fc0473302b59cd9e37af600d6a99b2542498832200b41a5106654199edd69816a61b2b4008ec5302fbc2dcb56d84ecf0917dfbbc64ab40e
[kubernetes-src.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-src.tar.gz) | 41e2fa601b7c4b2695ff081898a19261ce6c6381cc0d0547cd87378b210c7cdc60d66e75e7126a7b3c77ecc7574ff4ef9108814a6066e2dbd3b52d8693252c38

### Client Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-client-darwin-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-client-darwin-amd64.tar.gz) | 3782d576437ad78702caa2d7243d00a352245f01a72c3e8ebad690eb56ed4f9a215c40e99b0ef7b4e4afcf1c05d6f747798f697fe2b9e171168efa15c68c8b86
[kubernetes-client-darwin-arm64.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-client-darwin-arm64.tar.gz) | 42dbe3f4ccf968754665a73444cee465c00caf1d13bfead15e037d309633da5caddee100a313bf6c7da50075b7909c42a1f9c9b4ed7bf6bc73456238c9d93178
[kubernetes-client-linux-386.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-client-linux-386.tar.gz) | 11a417180848ffed539733e8483eb152a53f71f2be1d2b8d98f584334feca5cd9dc2ba2bd6b519e463d0da6c460409f56dae0728923e370a9a0489943e332a79
[kubernetes-client-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-client-linux-amd64.tar.gz) | 91b9c1e284a5de454518edd6cfd1eec6b6fc4b2b523f2f68a09f349cc951828cd7af85e36a52f46c6cc7f3f418da10bbacc4bb0a8c7ca349231b76cf6d0a241e
[kubernetes-client-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-client-linux-arm.tar.gz) | 101baed202c7e6b9e9fba5c6fb6e2b4b967717e6ea0a2440d7801490eff73e34cbee7f58afec2d270c07faf6bc499cba4472d15f41ca489940e3703a91cc681d
[kubernetes-client-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-client-linux-arm64.tar.gz) | ea1adb89bb898cc57bd3d6641a7da3320609b56c7b16dc2631ae35e23d969c4d5c58ed14779fd2ef61238e08b0dfc4851662d69786284af02872a5b4bde19ca2
[kubernetes-client-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-client-linux-ppc64le.tar.gz) | 7c2dfb68dcb5ca46c503d88f10ae0d68820fe2dcc47ba62c1a1ee676fd47ec1e6fdec773603e1155a8d04324a38e177440eda07c13f48d588e4f01d34070aa61
[kubernetes-client-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-client-linux-s390x.tar.gz) | 8a6fa28269f734fd752c8f3560a476d93768df25c1aa39e235a443cd26fa4f37a177e5508f75ed143335cadbb589afe68eaedb1ea4d0edf3d122ac231eca9196
[kubernetes-client-windows-386.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-client-windows-386.tar.gz) | ae0029235bdf049349d66a38e095834fc14ad9c4d380f0befd798a6267d2a1dc1e6e36242c466053a05dc2e4c8684e9d2b6aa923c82d9d283c9f4e06a05cbd03
[kubernetes-client-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-client-windows-amd64.tar.gz) | e197227390a93093426705d05091c85bbab5c705dc52073c5d087c0026bdf4ab882e314904b451987152c242b4accdc2c8092ca333c5b228161902039beef9a3

### Server Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-server-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-server-linux-amd64.tar.gz) | ccd594e9b0ca70ef90c01f5f0c6c34f7c0531bbdaaac54e8f61f12d3cfe8060f312d6865d60aad12a52aaad4674923d22ec8bb3c3dd692a88b42b0ae49801b7e
[kubernetes-server-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-server-linux-arm.tar.gz) | 4378866e6190a0315c4e19645fe09e9de5f5c180f7201a1a5d733e106e7bf24dad2748a26a3c629b988b8928c7b92486f404f7981f53deb4c39150e95635a818
[kubernetes-server-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-server-linux-arm64.tar.gz) | 977ce8965f3f55161652da41bfc51674dbd7f63d53ee3eaad69d92a448746d0f2eaf973af7189267275cdfb537018df2c54cd89ff4b46cc44cb964a5c7e7e4e3
[kubernetes-server-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-server-linux-ppc64le.tar.gz) | 196276a28a4dc2e04b4a0a8fc3b369b69cc58bfecafd29b4d026f01b95b186ec1bb8ea5ab26e3b51f478d9689a42be76de42318d7cbc9cf199e7b4cda76f89be
[kubernetes-server-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-server-linux-s390x.tar.gz) | ea2f0ebd438dfe3373bca05a1382945bbaa05ae3ec8474af1dcf641c82da8f3bfff5f309039243629ca485c699a46bc85b3f0f6fbd682bd340bc4869d3dc4c6d

### Node Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-node-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-node-linux-amd64.tar.gz) | 2f8238cab106e21f9e37c1423b2a04ff4e6e3947a7aeb786dc576af8d1fbff67a4719f79e79bb113abe47a42610c30ee3dde13322c6240b46d8ed641bf881de3
[kubernetes-node-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-node-linux-arm.tar.gz) | fa9038d9fbdfca5e984c86697f78464121393086f5f4a478898656a81a5f37314a5e31a54026a04afa7af3eb62cb8fffe21a5f8f233fd8731aa1fa10c40a6ea9
[kubernetes-node-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-node-linux-arm64.tar.gz) | 9d772b9e5cc59000c46fc3984c702ea569c27cfd4f1f7795eb4c1a79dac8a84cd834402faa282dbae9041c61443a2c3c677b9b6a2c8fe3d0a9f8df68cb34e7f0
[kubernetes-node-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-node-linux-ppc64le.tar.gz) | 8213b88a902aaf9c6a90e80a3c7af492e5356ddf78edb2f9666090362e12a428b782b70ad6a7e0742873d6ac05e746b0a2a819be9b94ae443524963468054df7
[kubernetes-node-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-node-linux-s390x.tar.gz) | f26af3bf8dc9499592cdaba971f087996ba8756237da7f1f5bed1e04ad87d90ca980ac325819222da7dadbf36bc6fa0c86bd227160c9b009be551ed6045b4eb3
[kubernetes-node-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.1/kubernetes-node-windows-amd64.tar.gz) | 87fd57ba83810e17dd8b64da9e17161abf5c1ba92fa6dbe063de073506a0fb031cb175484a0fe4ea8be85796f7c10b2e288b6d736520eb5c15277046ec57617d

## Changelog since v1.22.0-beta.0

## Urgent Upgrade Notes

### (No, really, you MUST read this before you upgrade)

 - The flag --experimental-patches is now deprecated and will be removed in a future release. You can migrate to using the new flag --patches. Add a new field {Init|Join}Configuration.patches.directory that can be used for the same purpose. For "init" and "join" it is now recommended that you migrate to configure patches via {Init|Join}Configuration.patches.directory. For the time being, these flags can be mixed with --config, but that might change in the future. On a command line, the last *patches flag takes precedence over previous flags and the value in config. "kubeadm upgrade" --patches will continue to be the only available option, since "upgrade" does not support a configuration file yet. ([#103063](https://github.com/kubernetes/kubernetes/pull/103063), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
 
## Changes by Kind

### Deprecation

- The feature Dynamic Kubelet Configuration is deprecated and kubelet will report warning when the flag `--dynamic-config-dir` is used. Feature gate `DynamicKubeletConfig` is disabled out of the box and needs to be explicitly enabled. ([#102966](https://github.com/kubernetes/kubernetes/pull/102966), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev)) [SIG Cloud Provider, Instrumentation and Node]
- The in-tree azure and gcp auth plugins have been deprecated.  The https://github.com/Azure/kubelogin and gcloud commands serve as out-of-tree replacements via the kubectl/client-go credential plugin mechanism. ([#102181](https://github.com/kubernetes/kubernetes/pull/102181), [@enj](https://github.com/enj)) [SIG API Machinery and Auth]

### API Change

- A new score extension for NodeResourcesFit plugin that merges the functionality of NodeResourcesLeastAllocated,NodeResourcesMostAllocated,RequestedToCapacityRatio plugins, which are marked as deprecated as of v1beta2. In v1beta1, the three plugins can still be used in v1beta1 but not at the same time with the score extension of NodeResourcesFit
  ```
  
  #### Additional documentation e.g., KEPs (Kubernetes Enhancement Proposals), usage docs, etc.:
  
  <!--
  This section can be blank if this pull request does not require a release note.
  
  When adding links which point to resources within git repositories, like
  KEPs or supporting documentation, please reference a specific commit and avoid
  linking directly to the master branch. This ensures that links reference a
  specific point in time, rather than a document that may change over time.
  
  See here for guidance on getting permanent links to files: https://help.github.com/en/articles/getting-permanent-links-to-files
  
  Please use the following format for linking documentation:
  - [KEP]: <link>
  - [Usage]: <link>
  - [Other doc]: <link>
  --> ([#101822](https://github.com/kubernetes/kubernetes/pull/101822), [@yuzhiquan](https://github.com/yuzhiquan)) [SIG API Machinery, Apps, Auth, Cloud Provider, Instrumentation, Node, Scheduling, Storage and Testing]
- Add DataSourceRef alpha field to PVC spec, which allows contents other than PVCs and VolumeSnapshots to be data sources. ([#103276](https://github.com/kubernetes/kubernetes/pull/103276), [@bswartz](https://github.com/bswartz)) [SIG API Machinery, Apps and Storage]
- Add PersistentVolumeClaimDeletePoilcy to StatefulSet API. ([#99378](https://github.com/kubernetes/kubernetes/pull/99378), [@mattcary](https://github.com/mattcary)) [SIG API Machinery and Apps]
- Add distributed tracing to the kube-apiserver.  It is can be enabled with the feature gate: APIServerTracing=true ([#94942](https://github.com/kubernetes/kubernetes/pull/94942), [@dashpole](https://github.com/dashpole)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Storage and Testing]
- Added new kubelet alpha feature `SeccompDefault`. This feature enables falling back to
  the `RuntimeDefault` (former `runtime/default`) seccomp profile if nothing else is specified
  in the pod/container `SecurityContext` or the pod annotation level. To use the feature, enable
  the feature gate as well as set the kubelet configuration option `SeccompDefault`
  (`--seccomp-default`) to `true`. ([#101943](https://github.com/kubernetes/kubernetes/pull/101943), [@saschagrunert](https://github.com/saschagrunert)) [SIG Node]
- Adds the ReadWriteOncePod access mode for PersistentVolumes and PersistentVolumeClaims. Restricts volume access to a single pod on a single node. ([#102028](https://github.com/kubernetes/kubernetes/pull/102028), [@chrishenzie](https://github.com/chrishenzie)) [SIG Apps, CLI, Node, Scheduling and Storage]
- Alpha swap support can now be enabled on Kubernetes nodes with the NodeSwapEnabled feature flag. See <website link> for details. ([#102823](https://github.com/kubernetes/kubernetes/pull/102823), [@ehashman](https://github.com/ehashman)) [SIG Node]
- CSIServiceAccountToken is GA. ([#103001](https://github.com/kubernetes/kubernetes/pull/103001), [@zshihang](https://github.com/zshihang)) [SIG Auth and Storage]
- Enforce the ReadWriteOncePod PVC access mode during scheduling ([#103082](https://github.com/kubernetes/kubernetes/pull/103082), [@chrishenzie](https://github.com/chrishenzie)) [SIG Apps, CLI, Node, Scheduling and Storage]
- Improved parsing of label selectors ([#102188](https://github.com/kubernetes/kubernetes/pull/102188), [@alculquicondor](https://github.com/alculquicondor)) [SIG API Machinery]
- Kube API server accepts Impersonate-Uid header to impersonate a user with a specific UID, in the same way that you can currently use Impersonate-User, Impersonate-Group and Impersonate-Extra ([#99961](https://github.com/kubernetes/kubernetes/pull/99961), [@margocrawf](https://github.com/margocrawf)) [SIG API Machinery, Auth and Testing]
- Kube-scheduler: a plugin enabled in a v1beta2 configuration file takes precedence over the default configuration for that plugin; this simplifies enabling default plugins with custom configuration without needing to explicitly disable those default plugins. ([#99582](https://github.com/kubernetes/kubernetes/pull/99582), [@chendave](https://github.com/chendave)) [SIG Scheduling]
- Scheduler could be configured  to consider new resources beside CPU and memory,  GPU for example, for the score plugin of `NodeResourcesBalancedAllocation`. ([#101946](https://github.com/kubernetes/kubernetes/pull/101946), [@chendave](https://github.com/chendave)) [SIG Scheduling]
- Suspend Job feature graduated to beta
  Added the "action" label to Job controller sync metrics job_sync_total and job_sync_duration_seconds ([#102022](https://github.com/kubernetes/kubernetes/pull/102022), [@adtac](https://github.com/adtac)) [SIG Apps, Instrumentation and Testing]
- TerminationGracePeriodSeconds on pod specs and container probes should not be negative.
  Negative values of TerminationGracePeriodSeconds will be treated as the value `1s` on the delete path.
  Immutable field validation will be relaxed in order to update negative values. 
  In a future release, negative values will not be permitted. ([#98866](https://github.com/kubernetes/kubernetes/pull/98866), [@wzshiming](https://github.com/wzshiming)) [SIG API Machinery, Apps and Node]
- The API documentation for the DaemonSet's spec.updateStrategy.rollingUpdate.maxUnavailable field was corrected to state that the value is rounded up. ([#101296](https://github.com/kubernetes/kubernetes/pull/101296), [@Miciah](https://github.com/Miciah)) [SIG Apps and CLI]
- The CertificateSigningRequest.certificates.k8s.io API supports an optional expirationSeconds field to allow the client to request a particular duration for the issued certificate.  The default signer implementations provided by the Kubernetes controller manager will honor this field as long as it does not exceed the --cluster-signing-duration flag. ([#99494](https://github.com/kubernetes/kubernetes/pull/99494), [@enj](https://github.com/enj)) [SIG API Machinery, Apps, Auth, CLI, Instrumentation, Node, Security and Testing]
- The ServiceAppProtocol feature gate has been removed. It reached GA in Kubernetes 1.20. ([#103190](https://github.com/kubernetes/kubernetes/pull/103190), [@robscott](https://github.com/robscott)) [SIG Network]

### Feature

- (alpha feature) If the CSI driver supports the NodeServiceCapability `VOLUME_MOUNT_GROUP` and the `DelegateFSGroupToCSIDriver` feature gate is enabled, kubelet will delegate applying FSGroup to the driver by passing it to NodeStageVolume and NodePublishVolume, regardless of what other FSGroup policies are set. ([#103244](https://github.com/kubernetes/kubernetes/pull/103244), [@verult](https://github.com/verult)) [SIG Apps and Storage]
- API Server tracing can now trace re-entrant api requests. ([#103218](https://github.com/kubernetes/kubernetes/pull/103218), [@dashpole](https://github.com/dashpole)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle and Instrumentation]
- Add unified map on CRI to support cgroup v2. Refer to https://github.com/opencontainers/runtime-spec/blob/master/config-linux.md#unified. ([#102578](https://github.com/kubernetes/kubernetes/pull/102578), [@payall4u](https://github.com/payall4u)) [SIG Node]
- Added field-selector option for `kubectl top pod` ([#102155](https://github.com/kubernetes/kubernetes/pull/102155), [@lauchokyip](https://github.com/lauchokyip)) [SIG CLI]
- Added new metrics about API Priority and Fairness.  Each one has a label `priority_level`.  The last two also have a label `bound` taking values `min` and `max.
  - apiserver_flowcontrol_current_r: R(the time of the last change in state of the queues)
  - apiserver_flowcontrol_dispatch_r: R(the time of the latest request dispatch)
  - apiserver_flowcontrol_latest_s: S(the request last dispatched) = R(when that request starts executing in the virtual world)
  - apiserver_flowcontrol_next_s_bounds: min and max next S among non-empty queues
  - apiserver_flowcontrol_next_discounted_s_bounds: min and max next S - (sum [over requests executing] width * estimatedDuration) among non-empty queues ([#102859](https://github.com/kubernetes/kubernetes/pull/102859), [@MikeSpreitzer](https://github.com/MikeSpreitzer)) [SIG API Machinery and Instrumentation]
- Adding `--restart-kubelet` flag on E2E Node test suite ([#97028](https://github.com/kubernetes/kubernetes/pull/97028), [@knabben](https://github.com/knabben)) [SIG Node and Testing]
- Adds feature gate `KubeletInUserNamespace` which enables support for running kubelet in a user namespace.
  
  The user namespace has to be created before running kubelet.
  All the node components such as CRI need to be running in the same user namespace.
  
  When the feature gate is enabled, kubelet ignores errors that happens during setting the following sysctl values: `vm.overcommit_memory`, `vm.panic_on_oom`, `kernel.panic`, `kernel.panic_on_oops`, `kernel.keys.root_maxkeys`, `kernel.keys.root_maxbytes`. (These sysctl values for the host, not for the containers)
  
  kubelet also ignores an error during opening `/dev/kmsg`.
  This feature gate also allows kube-proxy to ignore an error during setting `RLIMIT_NOFILE`.
  
  This feature gate is especially useful for running Kubernetes inside Rootless Docker/Podman with `kind` or `minikube`. ([#92863](https://github.com/kubernetes/kubernetes/pull/92863), [@AkihiroSuda](https://github.com/AkihiroSuda)) [SIG Network, Node and Testing]
- Adds metrics for the delegated authenticator used by extension APIs that delegate authentication logic to the Kube API server. ([#99364](https://github.com/kubernetes/kubernetes/pull/99364), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery, Auth, Instrumentation and Node]
- Adds metrics for the delegated authorizer used by extension APIs that delegate authorization logic to the Kube API server. ([#100339](https://github.com/kubernetes/kubernetes/pull/100339), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery, Auth, Instrumentation and Node]
- Client-go credential plugins are now GA ([#102890](https://github.com/kubernetes/kubernetes/pull/102890), [@ankeesler](https://github.com/ankeesler)) [SIG API Machinery, Auth and Testing]
- Deprecated warning message for igonre-errors flag. ([#102677](https://github.com/kubernetes/kubernetes/pull/102677), [@yuzhiquan](https://github.com/yuzhiquan)) [SIG CLI]
- Endpoints that have more than 1000 endpoints will be truncated and the `endpoints.kubernetes.io/over-capacity` annotation on the Endpoints resource will be set to `truncated`. ([#103520](https://github.com/kubernetes/kubernetes/pull/103520), [@swetharepakula](https://github.com/swetharepakula)) [SIG Apps and Network]
- Expose `/debug/flags/v` to allow dynamically setting log level for kube-proxy. ([#98306](https://github.com/kubernetes/kubernetes/pull/98306), [@borgerli](https://github.com/borgerli)) [SIG Network]
- Expose container start time as container_start_time_seconds in kubelet /metrics/resource endpoint ([#102444](https://github.com/kubernetes/kubernetes/pull/102444), [@sanwishe](https://github.com/sanwishe)) [SIG Instrumentation, Node and Testing]
- Extended resources defined in LeastAllocated, MostAllocated and RequestedToCapacityRatio plugin argument are bypassed by the scheduler if the incoming Pod doesn't request them in the pod spec. ([#103169](https://github.com/kubernetes/kubernetes/pull/103169), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling]
- Features gates `EndpointSliceProxying` & `WindowsEndpointSliceProxying` are now GA and cannot be disabled. Kube-proxy will use EndpointSlices for endpoint information. ([#103451](https://github.com/kubernetes/kubernetes/pull/103451), [@swetharepakula](https://github.com/swetharepakula)) [SIG Network]
- IngressClassNamespacedParams feature gate has graduated to beta and is enabled by default. This means IngressClass resource will now have two new fields - `spec.paramters.namespace` and `spec.parameters.scope`. ([#101711](https://github.com/kubernetes/kubernetes/pull/101711), [@hbagdi](https://github.com/hbagdi)) [SIG Network]
- JSON logging format is no longer available by default in non-core Kubernetes Components and require owners to opt in. ([#102869](https://github.com/kubernetes/kubernetes/pull/102869), [@mengjiao-liu](https://github.com/mengjiao-liu)) [SIG API Machinery, Cluster Lifecycle and Instrumentation]
- Kube-apiserver: the alpha PodSecurity feature can be enabled by passing `--feature-gates=PodSecurity=true`, and enables controlling allowed pods using namespace labels. See https://git.k8s.io/enhancements/keps/sig-auth/2579-psp-replacement for more details. ([#103099](https://github.com/kubernetes/kubernetes/pull/103099), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Auth, Instrumentation, Release, Security and Testing]
- Kube-proxy uses V1 EndpointSlices. ([#103306](https://github.com/kubernetes/kubernetes/pull/103306), [@swetharepakula](https://github.com/swetharepakula)) [SIG Network]
- LogarithmicScaleDown is now Beta and enabled by default ([#101767](https://github.com/kubernetes/kubernetes/pull/101767), [@damemi](https://github.com/damemi)) [SIG Apps and Instrumentation]
- Make kubectl command headers default for beta. ([#103238](https://github.com/kubernetes/kubernetes/pull/103238), [@seans3](https://github.com/seans3)) [SIG CLI]
- Mark net.ipv4.ip_unprivileged_port_start as safe sysctl ([#103326](https://github.com/kubernetes/kubernetes/pull/103326), [@pacoxu](https://github.com/pacoxu)) [SIG Auth, Network, Node and Security]
- Move the memory manager feature to the beta that will make it enabled by default with the None policy. ([#101947](https://github.com/kubernetes/kubernetes/pull/101947), [@cynepco3hahue](https://github.com/cynepco3hahue)) [SIG Node and Testing]
- Promote the ServiceInternalTrafficPolicy feature to beta, which enables the internalTrafficPolicy field of Service by default. ([#103462](https://github.com/kubernetes/kubernetes/pull/103462), [@andrewsykim](https://github.com/andrewsykim)) [SIG Apps and Network]
- Promote the feature ServiceLBNodePortControl to Beta ([#100412](https://github.com/kubernetes/kubernetes/pull/100412), [@hanlins](https://github.com/hanlins)) [SIG API Machinery, Apps, Architecture, Network and Testing]
- Scheduler nows provides an option for plugin developers to move Pods to activeQ proactively. ([#103383](https://github.com/kubernetes/kubernetes/pull/103383), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling and Testing]
- Services with `externalTrafficPolicy: Local` now support graceful termination when using
  the iptables or ipvs mode of kube-proxy with EndpointSlices enabled. Specifically, if a
  connection for such a service arrives on a node when there are no "Ready" endpoints
  for the service, but there is at least one Terminating pod for that service on the node,
  then kube-proxy will send the traffic to the Terminating pod rather than dropping it. This
  patches up a race condition between when a pod is killed and when the external load
  balancer notices that it has been killed. ([#97238](https://github.com/kubernetes/kubernetes/pull/97238), [@andrewsykim](https://github.com/andrewsykim)) [SIG Network, Scalability and Windows]
- Shell completion has been migrated to Cobra's go solution.  `kubectl` is now smarter about disabling file completion when it does not apply.  Furthermore, completion for the `cp` command does not show all files unless the user has started typing something. ([#96087](https://github.com/kubernetes/kubernetes/pull/96087), [@marckhouzam](https://github.com/marckhouzam)) [SIG CLI]
- The `HugePageStorageMediumSize` feature has been promoted to GA, allowing unconditional usage of multiple sizes huge page resources on a container level. ([#99144](https://github.com/kubernetes/kubernetes/pull/99144), [@bart0sh](https://github.com/bart0sh)) [SIG Apps, Node and Storage]
- Update the system-validators library to v1.5.0. Includes validation for seccomp and fixes a stdout/stderr problem in the Docker validator. ([#103390](https://github.com/kubernetes/kubernetes/pull/103390), [@ironyman](https://github.com/ironyman)) [SIG Cluster Lifecycle]
- Updates the following images to pick up CVE fixes:
  - `debian` to v1.8.0
  - `debian-iptables` to v1.6.5
  - `setcap` to v2.0.3 ([#103235](https://github.com/kubernetes/kubernetes/pull/103235), [@thejoycekung](https://github.com/thejoycekung)) [SIG API Machinery, Release and Testing]
- `kubectl debug` is able to create ephemeral containers in pre-1.22 clusters with the `EphemeralContainers` feature enabled. Note that versions of kubectl prior to 1.22 are unable to create ephemeral containers in clusters version 1.22 and greater due to an API change. ([#103292](https://github.com/kubernetes/kubernetes/pull/103292), [@verb](https://github.com/verb)) [SIG CLI]

### Documentation

- `kubectl debug` will now print a warning message when using the `--target` option since many container runtimes do not support this yet. ([#101074](https://github.com/kubernetes/kubernetes/pull/101074), [@verb](https://github.com/verb)) [SIG CLI]

### Bug or Regression

- Cut strings at 100 chars or first new line in kubectl get ([#103514](https://github.com/kubernetes/kubernetes/pull/103514), [@soltysh](https://github.com/soltysh)) [SIG CLI]
- Disruption controller shouldn't error while syncing for unmanaged pods. ([#103414](https://github.com/kubernetes/kubernetes/pull/103414), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Apps and Testing]
- Fix kubectl alpha debug node does not work on tainted(NoExecute) nodes and tolerate everything. ([#98431](https://github.com/kubernetes/kubernetes/pull/98431), [@wawa0210](https://github.com/wawa0210)) [SIG CLI]
- Fix performance regression for update and apply operations on large CRDs. ([#103318](https://github.com/kubernetes/kubernetes/pull/103318), [@jpbetz](https://github.com/jpbetz)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Storage]
- Fixed SELinux relabeling of CSI volumes after CSI driver failure. ([#103154](https://github.com/kubernetes/kubernetes/pull/103154), [@jsafrane](https://github.com/jsafrane)) [SIG Node and Storage]
- Kubeadm: fix a bug where "kubeadm join" for control plane nodes would download certificates and keys from the cluster, but would not write publicly readable certificates and public keys with mode 0644 and instead use mode 0600. ([#103313](https://github.com/kubernetes/kubernetes/pull/103313), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubelet: the returned value for PodIPs is the same in the Downward API and in the pod.status.PodIPs field ([#103307](https://github.com/kubernetes/kubernetes/pull/103307), [@aojea](https://github.com/aojea)) [SIG Network and Node]
- Metrics Server will use Addon Manager 1.8.3 ([#103541](https://github.com/kubernetes/kubernetes/pull/103541), [@jbartosik](https://github.com/jbartosik)) [SIG Cloud Provider and Instrumentation]
- Pods that are known to the kubelet to have previously been Running should not revert to Pending state; the kubelet will now infer a termination. ([#102821](https://github.com/kubernetes/kubernetes/pull/102821), [@ehashman](https://github.com/ehashman)) [SIG Node]
- Switch scheduler to generate the merge patch on pod status instead of the full pod ([#103133](https://github.com/kubernetes/kubernetes/pull/103133), [@marwanad](https://github.com/marwanad)) [SIG Scheduling]
- The current behavior for Services that `IPFamilyPolicy` set as `PreferDualstack`. The current behavior when the cluster is upgraded to dual-stack is:
  - Services that have been set to IPFamilyPolicy = PreferDualstack will be upgraded when the service object is updated. e.g., when a user change a label.
  
  This behavior will change to:
  - Services that have been set  IPFamilyPolicy = PreferDualstack will not be upgraded when the service object is updated. User can still change policy, type etc and existing behaviors remain the same. ([#102898](https://github.com/kubernetes/kubernetes/pull/102898), [@khenidak](https://github.com/khenidak)) [SIG Network and Testing]
- Treat VSphere "File (vmdk path here) was not found" errors as success during volume deletion ([#92372](https://github.com/kubernetes/kubernetes/pull/92372), [@breunigs](https://github.com/breunigs)) [SIG Cloud Provider and Storage]
- Upgrades functionality of `kubectl kustomize` as described at
  https://github.com/kubernetes-sigs/kustomize/releases/tag/kustomize%2Fv4.2.0 ([#103419](https://github.com/kubernetes/kubernetes/pull/103419), [@natasha41575](https://github.com/natasha41575)) [SIG CLI]

### Other (Cleanup or Flake)

- Featuregate ServiceLoadBalancerClass will be promoted to beta, this enables the Service loadBalancerClass field by default. ([#103129](https://github.com/kubernetes/kubernetes/pull/103129), [@XudongLiuHarold](https://github.com/XudongLiuHarold)) [SIG Testing]
- Increased CSINodeIDMaxLength from 128  bytes to 192 bytes. Prepare to increase the length limit to 256 bytes in 1.23 release. ([#101256](https://github.com/kubernetes/kubernetes/pull/101256), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG Storage]
- Kubeadm: move the BootstrapToken* API and related utilities from v1beta3 to a separate API group/version - bootstraptoken/v1. ([#102964](https://github.com/kubernetes/kubernetes/pull/102964), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Remove balanced attached node volumes feature ([#102443](https://github.com/kubernetes/kubernetes/pull/102443), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Scheduling]
- Replace go-bindata with //go:embed ([#99829](https://github.com/kubernetes/kubernetes/pull/99829), [@palnabarun](https://github.com/palnabarun)) [SIG Architecture, CLI, Node and Testing]
- The deprecated RunAsGroup feature-gate has been removed, since the RunAsGroup feature graduated to GA in 1.21 and was unconditionally enabled. ([#101581](https://github.com/kubernetes/kubernetes/pull/101581), [@carlory](https://github.com/carlory)) [SIG API Machinery and Node]
- Update golang.org/x/net to v0.0.0-20210520170846-37e1c6afe023 ([#103176](https://github.com/kubernetes/kubernetes/pull/103176), [@CaoDonghui123](https://github.com/CaoDonghui123)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Node and Storage]
- Updated command descriptions and examples for grammar and punctuation consistency. ([#103524](https://github.com/kubernetes/kubernetes/pull/103524), [@bergerhoffer](https://github.com/bergerhoffer)) [SIG Auth and CLI]

## Dependencies

### Added
- github.com/felixge/httpsnoop: [v1.0.1](https://github.com/felixge/httpsnoop/tree/v1.0.1)
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.20.0

### Changed
- github.com/auth0/go-jwt-middleware: [5493cab → v1.0.1](https://github.com/auth0/go-jwt-middleware/compare/5493cab...v1.0.1)
- github.com/gopherjs/gopherjs: [0766667 → fce0ec3](https://github.com/gopherjs/gopherjs/compare/0766667...fce0ec3)
- github.com/heketi/heketi: [v10.2.0+incompatible → v10.3.0+incompatible](https://github.com/heketi/heketi/compare/v10.2.0...v10.3.0)
- github.com/mattn/go-runewidth: [v0.0.9 → v0.0.7](https://github.com/mattn/go-runewidth/compare/v0.0.9...v0.0.7)
- github.com/olekukonko/tablewriter: [v0.0.5 → v0.0.4](https://github.com/olekukonko/tablewriter/compare/v0.0.5...v0.0.4)
- github.com/smartystreets/assertions: [b2de0cb → v1.1.0](https://github.com/smartystreets/assertions/compare/b2de0cb...v1.1.0)
- golang.org/x/net: a5a99cb → 37e1c6a
- k8s.io/system-validators: v1.4.0 → v1.5.0
- k8s.io/utils: da69540 → 4b05e18
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.20 → v0.0.21
- sigs.k8s.io/kustomize/api: v0.8.10 → v0.8.11
- sigs.k8s.io/kustomize/cmd/config: v0.9.12 → v0.9.13
- sigs.k8s.io/kustomize/kustomize/v4: v4.1.3 → v4.2.0
- sigs.k8s.io/kustomize/kyaml: v0.10.20 → v0.11.0
- sigs.k8s.io/structured-merge-diff/v4: v4.1.1 → v4.1.2

### Removed
- github.com/etcd-io/gofail: [ad7f989](https://github.com/etcd-io/gofail/tree/ad7f989)
- github.com/go-bindata/go-bindata: [v3.1.1+incompatible](https://github.com/go-bindata/go-bindata/tree/v3.1.1)
- go.etcd.io/etcd/etcdutl/v3: v3.5.0
- go.etcd.io/etcd/tests/v3: v3.5.0



# v1.22.0-beta.0


## Downloads for v1.22.0-beta.0

### Source Code

filename | sha512 hash
-------- | -----------
[kubernetes.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes.tar.gz) | 0ef8d9ff8a70f46adf295de2557dc5ba16eb50c6834d94fa29af3d6db97960d2d514c2b6bf40914ee7bb03ed1c30bbb581a7b311eb28c391215a8b96c14f2f34
[kubernetes-src.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-src.tar.gz) | 81f93ac8a1dbba3f8b86182a26eca81d7801f6d3b0ad0a4bdeb88c5782d569ceb7d7720fd7922e118b479645d1915c9fc0d8be3874e68f58c011c8a5a2e7649b

### Client Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-client-darwin-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-client-darwin-amd64.tar.gz) | 6e14453938a4ac3bf5c8fa8fafa46335ff8764123d9e49305fd167315556dcf3dbd6e8a8daabf94556a2a3089d86db7a79cf873d91f1a9c34cb97bd3d2870bd4
[kubernetes-client-darwin-arm64.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-client-darwin-arm64.tar.gz) | d6aaa33d6b9a0e8492522f2830422ca365877a5cd3f6321a1343dc7c0732767b03a744c4c1c3452b5bfdac6defbcfdfc0bb9f85e900c4c4836f2183839985c1e
[kubernetes-client-linux-386.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-client-linux-386.tar.gz) | 81a1eebd9473d818292fdde9aa883a88ed25af4996ad69cd0f0650ff1bf63362cd9ae4e5ff7b5d3efe72e48ac9318b49a1a62694ca90154aa6cd2bf86c624556
[kubernetes-client-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-client-linux-amd64.tar.gz) | f55f798baebb36c3bc59209ca9e36b5b916970ca3be02bf992cbbebda8817b6fe111b4e8f75b9411e4ddfc25e68be981829cae5c16ced0b5c57ededd718fd367
[kubernetes-client-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-client-linux-arm.tar.gz) | 741ee724dabfd06eb6de932baff48ad7eae16c518f47d501a83a33b6fdfd6900f96681e7d10f163e657eaeaf4c8bc10dff3f7a958f022513f9b85d1a014ddcf5
[kubernetes-client-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-client-linux-arm64.tar.gz) | a42e04a8ad3d2ab8a41a0629cce043951ace195286d03263079993f78e4a1827be2ab0126cf506ceb53f07edf116093ad41cfe5b2779b4d3150667a6be0c187a
[kubernetes-client-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-client-linux-ppc64le.tar.gz) | 830a530cc5dd1f5269238b4d6c662288e5410a8122f46f8d149aa4f1487e6dad8766277ad936df6fbd3f57272c3a4a77b4bd16bb8021db9a567db7984844f0d7
[kubernetes-client-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-client-linux-s390x.tar.gz) | 01a050312946b43295736b97d0deaef88adb73082ce46a89db9ff2a06ab3cd0e5c336fb5068081b5cd57bdab1d45044bae687c8247abe4e6aeea84c2f9db33ee
[kubernetes-client-windows-386.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-client-windows-386.tar.gz) | 4fee154708baf77a7c18c63767f25f0a8b1d33fa526a48ba794abefee11844cd99a73cad8bd316639f954850db038534544b26e0548ed56c2ba9c9c5c53491c9
[kubernetes-client-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-client-windows-amd64.tar.gz) | 10c075dc41cd5a11687f9e4c893a5c05ab350de78d10ded3db66e544b137b5737f9523aa30469887e588079356a0755aab92a3dbfb4c0fe7f3c431515b14d0e4

### Server Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-server-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-server-linux-amd64.tar.gz) | 36e9dc75a3b2663c38d9e144a8d3b402d9ae79aa958f334d3dba771eed452c6e427b91feada87cf952f3b438cca37ff8a003a90af94008225496d798f29ecbe8
[kubernetes-server-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-server-linux-arm.tar.gz) | f21f9d831fa052b2560c536e345e9432861dc028d85eec10c03e25329e1747f35621f4412e839e048fc036c0c20ae5c343666d720f5274068968d5ee83057567
[kubernetes-server-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-server-linux-arm64.tar.gz) | 2647235317883cfa9e194cb760ff7c0fe60d22cb86a09e1ac0494da29b5ff46d2b73193212f2cc3665ca755ef9e542f1f926132657ca6938470fd52b84deb5bf
[kubernetes-server-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-server-linux-ppc64le.tar.gz) | bf1126dfdb7822664b64e46598ee806e86a8bf23dc8f7c1725ce88e9fefec770b0b10ac093b97483f95d70810ca9c2051e69d1efbdbfac83b3f2e3e488489e33
[kubernetes-server-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-server-linux-s390x.tar.gz) | 5919cd20f62aca48ed7fdbea1e411004ac079f9a413c769db8949f177ed377a2fcd3f310ef1b75514c7ae5f18ce68762b19c884fe93b14788d9707f977921721

### Node Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-node-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-node-linux-amd64.tar.gz) | 00012a87c2ccd919a1f73aa7f47f2b82ff7c256e3d3c2f5c0d3b7c48bb7011c49047203dbaf99743825f3f156f78ac60522620962fa29fca084b9a2ec89b588d
[kubernetes-node-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-node-linux-arm.tar.gz) | 70ad434c4d0999e3cf6f764a4e03393b320f6560c4f8e08414fedfa62a45cfedd9249738e77f12e4fd2101d017c748369ebb46d4057b14d6a49d4bf88e154bcb
[kubernetes-node-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-node-linux-arm64.tar.gz) | d242bedd238cc6f400f9177bf90a5f36926e9eb3f27c9cf71b5519be72e215546925b8d36c1cce708ad620009c7f8f0d1f955781e3bcda4b8c304afeda53373c
[kubernetes-node-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-node-linux-ppc64le.tar.gz) | 386ffac107d59246dfcb965c640c4c425d4a47405366779c7b43409194cccbb95fe9a798a1383c636d6d150fc9203a4c7c42b737f41d95bdac94d9df39c40608
[kubernetes-node-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-node-linux-s390x.tar.gz) | 6046b7c9343429fc6af4e850222ae3a6883d5f8237668d0905c48c7b5001cb8c4a2f27536b0b265c410ed746a9bef4789c0e1d9bc6e8c3e99587103367a318a8
[kubernetes-node-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-beta.0/kubernetes-node-windows-amd64.tar.gz) | e332ee42a9fb8ec6d7bfb8b6509447276dec0d5b04d0385a0ba8c05359cd02305afe2c9d3f3dfe6e5814afd78ca3635bee6a49209e598e0612969bce28b0cd8a

## Changelog since v1.22.0-alpha.3

## Urgent Upgrade Notes

### (No, really, you MUST read this before you upgrade)

 - Kubeadm: the flag --cri-socket is no longer allowed in a mixture with the flag --config. Please use the kubeadm configuration for setting the CRI socket for a node using {Init|Join}Configuration.nodeRegistration.criSocket. ([#101600](https://github.com/kubernetes/kubernetes/pull/101600), [@KofClubs](https://github.com/KofClubs)) [SIG Cluster Lifecycle]
 
## Changes by Kind

### Deprecation

- Deprecate --record flag in kubectl. The --record flag is being replaced with the mechanism from https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/859-kubectl-headers which annotates HTTP requests with kubectl command details. ([#102873](https://github.com/kubernetes/kubernetes/pull/102873), [@soltysh](https://github.com/soltysh)) [SIG CLI]
- E2e.test: removed the  `--viper-config` flag.  If you were previously using this to pass flags to `e2e.test` via a file, you will need to pass them directly on the command line, e.g. `e2e.test --e2e-output-dir` ([#102598](https://github.com/kubernetes/kubernetes/pull/102598), [@dims](https://github.com/dims)) [SIG Testing]
- Newly provisioned PVs by vSphere in-tree plugin will no longer have the beta FailureDomain label. vSphere volume plugin will start to have GA topology label instead. ([#102414](https://github.com/kubernetes/kubernetes/pull/102414), [@divyenpatel](https://github.com/divyenpatel)) [SIG Cloud Provider, Storage and Testing]
- The deprecated flag --hard-pod-affinity-symmetric-weight and --scheduler-name have been removed from kube-scheduler. Use ComponentConfig instead to configure those parameters. ([#102805](https://github.com/kubernetes/kubernetes/pull/102805), [@ahg-g](https://github.com/ahg-g)) [SIG Scheduling]

### API Change

- Because of the implementation logic of time.Format in golang, the displayed time zone is not consistent ([#102366](https://github.com/kubernetes/kubernetes/pull/102366), [@cndoit18](https://github.com/cndoit18)) [SIG Apps, Auth, Autoscaling, CLI, Cluster Lifecycle, Instrumentation, Network, Node and Testing]
- Endpoint slices mirroring controller no longer mirrors the last-applied-configuration annotation created by kubectl to updated endpoint slices ([#102731](https://github.com/kubernetes/kubernetes/pull/102731), [@sharmarajdaksh](https://github.com/sharmarajdaksh)) [SIG API Machinery, Apps, Cloud Provider, Network, Release, Scheduling, Storage and Testing]
- Exec plugin authors can override default handling of standard input via new interactiveMode kubeconfig field ([#99310](https://github.com/kubernetes/kubernetes/pull/99310), [@ankeesler](https://github.com/ankeesler)) [SIG API Machinery, Auth, CLI and Testing]
- Kube-scheduler component config v1beta2 API available
  Three scheduler plugins deprecated (NodeLabel, ServiceAffinity, NodePreferAvoidPods) ([#99597](https://github.com/kubernetes/kubernetes/pull/99597), [@adtac](https://github.com/adtac)) [SIG Scheduling]
- Network Policy EndPort is graduated to beta and is enabled by default ([#102834](https://github.com/kubernetes/kubernetes/pull/102834), [@rikatz](https://github.com/rikatz)) [SIG Network]

### Feature

- A new metric apiserver_flowcontrol_request_concurrency_in_use that shows the number of
  seats (concurrency) occupied by the currently executing requests in the API Priority and Fairness system ([#102795](https://github.com/kubernetes/kubernetes/pull/102795), [@tkashem](https://github.com/tkashem)) [SIG API Machinery and Instrumentation]
- Add a namespace label for all apiserver_admission_* metrics.
  - Expand the histogram range to 0-10s for all apiserver_admission_*_duration_seconds metrics. ([#101208](https://github.com/kubernetes/kubernetes/pull/101208), [@voutcn](https://github.com/voutcn)) [SIG API Machinery and Instrumentation]
- Add unified map on CRI to support cgroup v2. Refer to https://github.com/opencontainers/runtime-spec/blob/master/config-linux.md#unified. ([#102578](https://github.com/kubernetes/kubernetes/pull/102578), [@payall4u](https://github.com/payall4u)) [SIG Node]
- Base-images: Update to debian-base:buster-v1.7.1 ([#102594](https://github.com/kubernetes/kubernetes/pull/102594), [@mengjiao-liu](https://github.com/mengjiao-liu)) [SIG API Machinery and Release]
- Implement minReadySeconds for the StatefulSets. ([#101316](https://github.com/kubernetes/kubernetes/pull/101316), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Apps and Testing]
- Kubeadm: add the "--dry-run" flag to the "control-plane" phase of "kubeadm init" ([#102722](https://github.com/kubernetes/kubernetes/pull/102722), [@vinayakankugoyal](https://github.com/vinayakankugoyal)) [SIG Cluster Lifecycle]
- Kubeadm: add the "imagePullPolicy" field in the "nodeRegistration" section of InitConfiguration and JoinConfiguration in v1beta3. This allows the user to specify the image pull policy during "kubeadm init" and "kubeadm join". The value of this field must be one of "Always", "IfNotPresent" or "Never". The default behavior continues to be "IfNotPresent". ([#102901](https://github.com/kubernetes/kubernetes/pull/102901), [@wangyysde](https://github.com/wangyysde)) [SIG Cluster Lifecycle]
- Kubeadm: update CoreDNS to 1.8.4. Grant CoreDNS permissions to "list" and "watch" EndpointSlice objects to accommodate dual-stack support. ([#102466](https://github.com/kubernetes/kubernetes/pull/102466), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle]
- Kubernetes is now built with Golang 1.16.5 ([#102689](https://github.com/kubernetes/kubernetes/pull/102689), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Leader Migration now support a wildcard component name and the default value. ([#102711](https://github.com/kubernetes/kubernetes/pull/102711), [@jiahuif](https://github.com/jiahuif)) [SIG API Machinery and Cloud Provider]
- SetHostnameAsFQDN has been graduated to GA, which means feature cannot be disabled using a feature gate. ([#101294](https://github.com/kubernetes/kubernetes/pull/101294), [@javidiaz](https://github.com/javidiaz)) [SIG Node]
- The pod resource API will provide memory manager metrics in the case when the memory manager feature gate enabled and the memory manager policy is static. ([#101030](https://github.com/kubernetes/kubernetes/pull/101030), [@cynepco3hahue](https://github.com/cynepco3hahue)) [SIG Node and Testing]
- Update etcd version to 3.5.0-beta.3 ([#102062](https://github.com/kubernetes/kubernetes/pull/102062), [@serathius](https://github.com/serathius)) [SIG API Machinery, Cloud Provider, Cluster Lifecycle and Testing]

### Bug or Regression

- After DBus restarts, make GracefulNodeShutdown work again ([#100369](https://github.com/kubernetes/kubernetes/pull/100369), [@wzshiming](https://github.com/wzshiming)) [SIG Node and Testing]
- Aggregate write permissions on events to edit and admin role ([#102858](https://github.com/kubernetes/kubernetes/pull/102858), [@tumido](https://github.com/tumido)) [SIG Auth]
- Changed the Graceful Node Shutdown Pod termination reason and message.
  Changed the Graceful Node Shutdown Pod rejection reason and message. ([#102840](https://github.com/kubernetes/kubernetes/pull/102840), [@Kissy](https://github.com/Kissy)) [SIG Node]
- Fix a bug on the endpoint slices mirroring controller where endpoint NotReadyAddresses were mirrored as Ready to the corresponding EndpointSlice ([#102683](https://github.com/kubernetes/kubernetes/pull/102683), [@aojea](https://github.com/aojea)) [SIG Apps and Network]
- Fix scoring for NodeResourcesMostAllocated and NodeResourcesBalancedAllocation plugins when nodes have containers with no requests. This was leaving to under-utilization of small nodes. ([#102925](https://github.com/kubernetes/kubernetes/pull/102925), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling]
- Fix: cleanup outdated routes ([#102935](https://github.com/kubernetes/kubernetes/pull/102935), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fixed a bug that scheduler extenders are not called on preemptions ([#103019](https://github.com/kubernetes/kubernetes/pull/103019), [@ordovicia](https://github.com/ordovicia)) [SIG Scheduling]
- Fixed an issue with kubectl on certain older version of Windows or when legacy console mode is enabled on Windows 8 which causes kubectl exec to crash ([#102825](https://github.com/kubernetes/kubernetes/pull/102825), [@n4j](https://github.com/n4j)) [SIG API Machinery, CLI, Cloud Provider and Cluster Lifecycle]
- Fixed starting new pods after previous pod timed out unmounting its volumes. ([#100183](https://github.com/kubernetes/kubernetes/pull/100183), [@jsafrane](https://github.com/jsafrane)) [SIG Node, Storage and Testing]
- Fixes a bug due to which the controller was not populating the lastSuccessfulTime field added to batch/v1 cronjob.status ([#102642](https://github.com/kubernetes/kubernetes/pull/102642), [@alaypatel07](https://github.com/alaypatel07)) [SIG Apps]
- Kube-apiserver now reports the synthetic verb when logging requests, better explaining the user intent and matching what is reported in the metrics. ([#102934](https://github.com/kubernetes/kubernetes/pull/102934), [@lavalamp](https://github.com/lavalamp)) [SIG API Machinery and Instrumentation]
- Kubeadm: allow passing the flag --log-file if --config is passed. If you wish to log to a file you must also pass --logtostderr=false or --alsologtostderr=true. Alternatively you can pipe to a file using "kubeadm ... | tee ...". ([#101449](https://github.com/kubernetes/kubernetes/pull/101449), [@CaoDonghui123](https://github.com/CaoDonghui123)) [SIG Cluster Lifecycle]
- Kubeadm: remove the "ephemeral_storage" request from the etcd static pod that kubeadm deploys on stacked etcd control plane nodes. This request has caused sporadic failures on some setups due to a problem in the kubelet with cadvisor and the LocalStorageCapacityIsolation feature gate. See this issue for more details: https://github.com/kubernetes/kubernetes/issues/99305 ([#102673](https://github.com/kubernetes/kubernetes/pull/102673), [@jackfrancis](https://github.com/jackfrancis)) [SIG Cluster Lifecycle]
- Updates hostprocess tests to specify user ([#102965](https://github.com/kubernetes/kubernetes/pull/102965), [@jsturtevant](https://github.com/jsturtevant)) [SIG Testing and Windows]
- VSphere: Fix regression during attach disk if datastore is within a storage folder or datastore cluster. ([#102892](https://github.com/kubernetes/kubernetes/pull/102892), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider and Storage]

### Other (Cleanup or Flake)

- Allow CSI drivers to just run offline expansion tests ([#102665](https://github.com/kubernetes/kubernetes/pull/102665), [@gnufied](https://github.com/gnufied)) [SIG Storage and Testing]
- Client-go: reduce verbosity of "Starting/Stopping reflector" messages to 3 again ([#102788](https://github.com/kubernetes/kubernetes/pull/102788), [@pohly](https://github.com/pohly)) [SIG API Machinery]
- JSON logging now supports having information about source code location in the logging format, source code information is available under the key "caller" ([#102437](https://github.com/kubernetes/kubernetes/pull/102437), [@MadhavJivrajani](https://github.com/MadhavJivrajani)) [SIG API Machinery, Cluster Lifecycle and Instrumentation]
- Migrate staging/src/k8s.io/apiserver/pkg/registry logs to structured logging. ([#98287](https://github.com/kubernetes/kubernetes/pull/98287), [@lala123912](https://github.com/lala123912)) [SIG API Machinery]
- Remove balanced attached node volumes feature ([#102443](https://github.com/kubernetes/kubernetes/pull/102443), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Scheduling]
- The etcd client has been updated to 3.5.0; github.com/golang/protobuf, google.golang.org/protobuf, and google.golang.org/grpc have been updated to current versions. ([#100488](https://github.com/kubernetes/kubernetes/pull/100488), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Storage and Testing]
- Update Azure Go SDK to v55.0.0 ([#102441](https://github.com/kubernetes/kubernetes/pull/102441), [@feiskyer](https://github.com/feiskyer)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node and Storage]
- Update Calico to v3.19.1 ([#102386](https://github.com/kubernetes/kubernetes/pull/102386), [@JornShen](https://github.com/JornShen)) [SIG Cloud Provider]
- Update dep google/gnostic and google/go-cmp to v0.5.5 and updating transitive dependencies protobuf ([#102783](https://github.com/kubernetes/kubernetes/pull/102783), [@mcbenjemaa](https://github.com/mcbenjemaa)) [SIG Architecture and Testing]

## Dependencies

### Added
- github.com/antihax/optional: [v1.0.0](https://github.com/antihax/optional/tree/v1.0.0)
- github.com/benbjohnson/clock: [v1.0.3](https://github.com/benbjohnson/clock/tree/v1.0.3)
- github.com/certifi/gocertifi: [2c3bb06](https://github.com/certifi/gocertifi/tree/2c3bb06)
- github.com/cncf/udpa/go: [5459f2c](https://github.com/cncf/udpa/go/tree/5459f2c)
- github.com/cockroachdb/errors: [v1.2.4](https://github.com/cockroachdb/errors/tree/v1.2.4)
- github.com/cockroachdb/logtags: [eb05cc2](https://github.com/cockroachdb/logtags/tree/eb05cc2)
- github.com/coredns/caddy: [v1.1.0](https://github.com/coredns/caddy/tree/v1.1.0)
- github.com/etcd-io/gofail: [ad7f989](https://github.com/etcd-io/gofail/tree/ad7f989)
- github.com/getsentry/raven-go: [v0.2.0](https://github.com/getsentry/raven-go/tree/v0.2.0)
- github.com/go-kit/kit: [v0.9.0](https://github.com/go-kit/kit/tree/v0.9.0)
- github.com/josharian/intern: [v1.0.0](https://github.com/josharian/intern/tree/v1.0.0)
- github.com/opentracing/opentracing-go: [v1.1.0](https://github.com/opentracing/opentracing-go/tree/v1.1.0)
- github.com/robfig/cron/v3: [v3.0.1](https://github.com/robfig/cron/v3/tree/v3.0.1)
- go.etcd.io/etcd/api/v3: v3.5.0
- go.etcd.io/etcd/client/pkg/v3: v3.5.0
- go.etcd.io/etcd/client/v2: v2.305.0
- go.etcd.io/etcd/client/v3: v3.5.0
- go.etcd.io/etcd/etcdutl/v3: v3.5.0
- go.etcd.io/etcd/pkg/v3: v3.5.0
- go.etcd.io/etcd/raft/v3: v3.5.0
- go.etcd.io/etcd/server/v3: v3.5.0
- go.etcd.io/etcd/tests/v3: v3.5.0
- go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc: v0.20.0
- go.opentelemetry.io/contrib: v0.20.0
- go.opentelemetry.io/otel/exporters/otlp: v0.20.0
- go.opentelemetry.io/otel/metric: v0.20.0
- go.opentelemetry.io/otel/oteltest: v0.20.0
- go.opentelemetry.io/otel/sdk/export/metric: v0.20.0
- go.opentelemetry.io/otel/sdk/metric: v0.20.0
- go.opentelemetry.io/otel/sdk: v0.20.0
- go.opentelemetry.io/otel/trace: v0.20.0
- go.opentelemetry.io/otel: v0.20.0
- go.opentelemetry.io/proto/otlp: v0.7.0
- go.uber.org/goleak: v1.1.10

### Changed
- github.com/Azure/azure-sdk-for-go: [v53.1.0+incompatible → v55.0.0+incompatible](https://github.com/Azure/azure-sdk-for-go/compare/v53.1.0...v55.0.0)
- github.com/Azure/go-ansiterm: [d6e3b33 → d185dfc](https://github.com/Azure/go-ansiterm/compare/d6e3b33...d185dfc)
- github.com/Azure/go-autorest/autorest/adal: [v0.9.10 → v0.9.13](https://github.com/Azure/go-autorest/autorest/adal/compare/v0.9.10...v0.9.13)
- github.com/Azure/go-autorest/autorest/to: [v0.2.0 → v0.4.0](https://github.com/Azure/go-autorest/autorest/to/compare/v0.2.0...v0.4.0)
- github.com/Azure/go-autorest/autorest: [v0.11.17 → v0.11.18](https://github.com/Azure/go-autorest/autorest/compare/v0.11.17...v0.11.18)
- github.com/Azure/go-autorest/logger: [v0.2.0 → v0.2.1](https://github.com/Azure/go-autorest/logger/compare/v0.2.0...v0.2.1)
- github.com/aws/aws-sdk-go: [v1.35.24 → v1.38.49](https://github.com/aws/aws-sdk-go/compare/v1.35.24...v1.38.49)
- github.com/cockroachdb/datadriven: [80d97fb → bf6692d](https://github.com/cockroachdb/datadriven/compare/80d97fb...bf6692d)
- github.com/container-storage-interface/spec: [v1.3.0 → v1.5.0](https://github.com/container-storage-interface/spec/compare/v1.3.0...v1.5.0)
- github.com/coredns/corefile-migration: [v1.0.11 → v1.0.12](https://github.com/coredns/corefile-migration/compare/v1.0.11...v1.0.12)
- github.com/coreos/go-systemd/v22: [v22.3.1 → v22.3.2](https://github.com/coreos/go-systemd/v22/compare/v22.3.1...v22.3.2)
- github.com/envoyproxy/go-control-plane: [5f8ba28 → 668b12f](https://github.com/envoyproxy/go-control-plane/compare/5f8ba28...668b12f)
- github.com/form3tech-oss/jwt-go: [v3.2.2+incompatible → v3.2.3+incompatible](https://github.com/form3tech-oss/jwt-go/compare/v3.2.2...v3.2.3)
- github.com/go-openapi/jsonpointer: [v0.19.3 → v0.19.5](https://github.com/go-openapi/jsonpointer/compare/v0.19.3...v0.19.5)
- github.com/go-openapi/jsonreference: [v0.19.3 → v0.19.5](https://github.com/go-openapi/jsonreference/compare/v0.19.3...v0.19.5)
- github.com/go-openapi/swag: [v0.19.5 → v0.19.14](https://github.com/go-openapi/swag/compare/v0.19.5...v0.19.14)
- github.com/golang/groupcache: [8c9f03a → 41bb18b](https://github.com/golang/groupcache/compare/8c9f03a...41bb18b)
- github.com/golang/mock: [v1.4.3 → v1.4.4](https://github.com/golang/mock/compare/v1.4.3...v1.4.4)
- github.com/golang/protobuf: [v1.4.3 → v1.5.2](https://github.com/golang/protobuf/compare/v1.4.3...v1.5.2)
- github.com/google/btree: [v1.0.0 → v1.0.1](https://github.com/google/btree/compare/v1.0.0...v1.0.1)
- github.com/google/go-cmp: [v0.5.4 → v0.5.5](https://github.com/google/go-cmp/compare/v0.5.4...v0.5.5)
- github.com/googleapis/gnostic: [v0.5.1 → v0.5.5](https://github.com/googleapis/gnostic/compare/v0.5.1...v0.5.5)
- github.com/grpc-ecosystem/go-grpc-middleware: [f849b54 → v1.3.0](https://github.com/grpc-ecosystem/go-grpc-middleware/compare/f849b54...v1.3.0)
- github.com/grpc-ecosystem/grpc-gateway: [v1.9.5 → v1.16.0](https://github.com/grpc-ecosystem/grpc-gateway/compare/v1.9.5...v1.16.0)
- github.com/jonboulle/clockwork: [v0.1.0 → v0.2.2](https://github.com/jonboulle/clockwork/compare/v0.1.0...v0.2.2)
- github.com/json-iterator/go: [v1.1.10 → v1.1.11](https://github.com/json-iterator/go/compare/v1.1.10...v1.1.11)
- github.com/mailru/easyjson: [v0.7.0 → v0.7.6](https://github.com/mailru/easyjson/compare/v0.7.0...v0.7.6)
- github.com/mattn/go-isatty: [v0.0.4 → v0.0.3](https://github.com/mattn/go-isatty/compare/v0.0.4...v0.0.3)
- github.com/mattn/go-runewidth: [v0.0.7 → v0.0.9](https://github.com/mattn/go-runewidth/compare/v0.0.7...v0.0.9)
- github.com/miekg/dns: [v1.1.35 → v1.0.14](https://github.com/miekg/dns/compare/v1.1.35...v1.0.14)
- github.com/moby/term: [df9cb8a → 9d4ed18](https://github.com/moby/term/compare/df9cb8a...9d4ed18)
- github.com/olekukonko/tablewriter: [v0.0.4 → v0.0.5](https://github.com/olekukonko/tablewriter/compare/v0.0.4...v0.0.5)
- github.com/prometheus/client_golang: [v1.7.1 → v1.11.0](https://github.com/prometheus/client_golang/compare/v1.7.1...v1.11.0)
- github.com/prometheus/procfs: [v0.2.0 → v0.6.0](https://github.com/prometheus/procfs/compare/v0.2.0...v0.6.0)
- github.com/rogpeppe/fastuuid: [6724a57 → v1.2.0](https://github.com/rogpeppe/fastuuid/compare/6724a57...v1.2.0)
- github.com/soheilhy/cmux: [v0.1.4 → v0.1.5](https://github.com/soheilhy/cmux/compare/v0.1.4...v0.1.5)
- github.com/spf13/cobra: [v1.1.1 → v1.1.3](https://github.com/spf13/cobra/compare/v1.1.1...v1.1.3)
- github.com/spf13/jwalterweatherman: [v1.1.0 → v1.0.0](https://github.com/spf13/jwalterweatherman/compare/v1.1.0...v1.0.0)
- github.com/tmc/grpc-websocket-proxy: [0ad062e → e5319fd](https://github.com/tmc/grpc-websocket-proxy/compare/0ad062e...e5319fd)
- github.com/yuin/goldmark: [v1.2.1 → v1.3.5](https://github.com/yuin/goldmark/compare/v1.2.1...v1.3.5)
- go.etcd.io/bbolt: v1.3.5 → v1.3.6
- go.uber.org/atomic: v1.6.0 → v1.7.0
- go.uber.org/multierr: v1.5.0 → v1.6.0
- go.uber.org/zap: v1.16.0 → v1.17.0
- golang.org/x/lint: 738671d → 6edffad
- golang.org/x/mod: ce943fd → v0.4.2
- golang.org/x/net: 3d97a24 → a5a99cb
- golang.org/x/sync: 67f06af → 036812b
- golang.org/x/sys: d19ff85 → 59db8d7
- golang.org/x/tools: v0.1.0 → v0.1.2
- google.golang.org/genproto: 8816d57 → f16073e
- google.golang.org/grpc: v1.27.1 → v1.38.0
- google.golang.org/protobuf: v1.25.0 → v1.26.0
- gopkg.in/yaml.v3: eeeca48 → 496545a
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.19 → v0.0.20

### Removed
- github.com/bifurcation/mint: [93c51c6](https://github.com/bifurcation/mint/tree/93c51c6)
- github.com/caddyserver/caddy: [v1.0.3](https://github.com/caddyserver/caddy/tree/v1.0.3)
- github.com/cenkalti/backoff: [v2.1.1+incompatible](https://github.com/cenkalti/backoff/tree/v2.1.1)
- github.com/cheekybits/genny: [9127e81](https://github.com/cheekybits/genny/tree/9127e81)
- github.com/go-acme/lego: [v2.5.0+incompatible](https://github.com/go-acme/lego/tree/v2.5.0)
- github.com/jimstudt/http-authentication: [3eca13d](https://github.com/jimstudt/http-authentication/tree/3eca13d)
- github.com/klauspost/cpuid: [v1.2.0](https://github.com/klauspost/cpuid/tree/v1.2.0)
- github.com/kylelemons/godebug: [d65d576](https://github.com/kylelemons/godebug/tree/d65d576)
- github.com/lucas-clemente/aes12: [cd47fb3](https://github.com/lucas-clemente/aes12/tree/cd47fb3)
- github.com/lucas-clemente/quic-clients: [v0.1.0](https://github.com/lucas-clemente/quic-clients/tree/v0.1.0)
- github.com/lucas-clemente/quic-go-certificates: [d2f8652](https://github.com/lucas-clemente/quic-go-certificates/tree/d2f8652)
- github.com/lucas-clemente/quic-go: [v0.10.2](https://github.com/lucas-clemente/quic-go/tree/v0.10.2)
- github.com/marten-seemann/qtls: [v0.2.3](https://github.com/marten-seemann/qtls/tree/v0.2.3)
- github.com/mholt/certmagic: [6a42ef9](https://github.com/mholt/certmagic/tree/6a42ef9)
- github.com/naoina/go-stringutil: [v0.1.0](https://github.com/naoina/go-stringutil/tree/v0.1.0)
- github.com/naoina/toml: [v0.1.1](https://github.com/naoina/toml/tree/v0.1.1)
- github.com/robfig/cron: [v1.1.0](https://github.com/robfig/cron/tree/v1.1.0)
- go.etcd.io/etcd: dd1b699
- go.uber.org/tools: 2cfd321
- gopkg.in/cheggaaa/pb.v1: v1.0.25
- gopkg.in/mcuadros/go-syslog.v2: v2.2.1
- gopkg.in/resty.v1: v1.12.0
- rsc.io/quote/v3: v3.1.0
- rsc.io/sampler: v1.3.0



# v1.22.0-alpha.3


## Downloads for v1.22.0-alpha.3

### Source Code

filename | sha512 hash
-------- | -----------
[kubernetes.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes.tar.gz) | d45d66834b05aa2cb0ad1b0b7e0a00a9f91b992c74fe50f7a7ad396dae07c5af4855bd6e0191b9a424192c4351ba3e269effa6bca4f7405346c2aae6950c1efa
[kubernetes-src.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-src.tar.gz) | 0393a37d11babd35b918413e39880fbc9bcc4a3437944a741fc830c54672a5917e38298a2430cf775dacbc97f9251674a3f8f5e58de77e7e15d8b9512cefffd7

### Client Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-client-darwin-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-client-darwin-amd64.tar.gz) | a210d70e90e96f669b74a09752afe8ac118d7922296216298266ddb5413b2d00b5384545099d9fa12102932b95be82f5a21e9fc7fd0512bb974182d4a348d35e
[kubernetes-client-darwin-arm64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-client-darwin-arm64.tar.gz) | da38e928622f813d83fa33048446d4f2204b96884dd6d8596889bf3d6a26ddbcd1bcb7741978233e31dd5216d1c0fa3e81908d57f7be74ad54cbae35e9d19be1
[kubernetes-client-linux-386.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-client-linux-386.tar.gz) | 7fb77598532a0f42a6e052c403933f26c387bf6cf93ce7df88de334bbe25812da52940be13ade26ec841cb80c08273eddb5a051e8eb1bfd6ec14190acad0e827
[kubernetes-client-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-client-linux-amd64.tar.gz) | eabba5509ae1b5f770573e5fd1cffa39a86e9b8b97c3c2e930a197600ad704ed22925246b697b1070ab826a7680a1cd4338b31fcbcb9e24fe7df56f426587465
[kubernetes-client-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-client-linux-arm.tar.gz) | b94023510bd349d48c6242c1fd09ef74560667ef4a1362f75c2f07f0ad10b2e6a0f2aab36346a0db77488382192ea1ba0f044a2bc9416288f7a368030b39f5f0
[kubernetes-client-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-client-linux-arm64.tar.gz) | 3a577bf292d25450babbd1a15e1921f9068a0c7c06bd54a5d88a3a97982fb2ec55ebaeecb77a502b1179473d45bf99b553a88dd85c5df7daa973c14d6905d7a6
[kubernetes-client-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-client-linux-ppc64le.tar.gz) | b43e844dbd57c75616d20bafe1af8a08245644648a11b067b45048abf4da7f6f5c0ce3b2cd0dd09dd6ca81cebb881cee623d78a4dd9e963ef69e1d27e7f38471
[kubernetes-client-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-client-linux-s390x.tar.gz) | 1c472acb904ce7980b78b3716c6d9660add5161d100e4c6dd0cee2c4fc9dd3ff43ef26c3f48d5467c9e29d2101e80d60713ac3ddd7b50018be114a799edb4600
[kubernetes-client-windows-386.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-client-windows-386.tar.gz) | d224a9eeebbcf8fda304dd78ae491d74b513c484938a8ee7d66ef67d6821718f0665ce3a6cd44e628de7fab58e8a487ae431521857a92be8f1c07840a1b11472
[kubernetes-client-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-client-windows-amd64.tar.gz) | 6d497e4be0a11c427d16b774c61a6b32e09e97ea085c5a961cf155440ccf6220f5be97722cbe957b14722ca32d4c3d18a66559859a3ca810dcb881c8529bf098

### Server Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-server-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-server-linux-amd64.tar.gz) | 297d5b7d4161b5374a17bb1709874fdae02bb368e81a45c82dfb90508e5796e6871512f2f5f555697ea14597a26195ae6bbbc3d71ca685e6c888cb5ced689ae8
[kubernetes-server-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-server-linux-arm.tar.gz) | 305bac3735d716432c76c7911ea618f7198de3e6e1d1ef46e574ed882a307ae2bd16e2bd489dad24e5c3c76cfeaed572513b2990d621e157857f70784ae0cf96
[kubernetes-server-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-server-linux-arm64.tar.gz) | 88c634a4aeca7a55c9179a9f394096acece1aa2362c6354a2a7fedc4db7add7bfb7b9954010e65c430d3ada790e6e3ecf19713b1368491548b5cbdc2db2baf07
[kubernetes-server-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-server-linux-ppc64le.tar.gz) | e04de27a7c855ee7f0c50e37498a634ff5885e96311447ae0e136a7c8321cdb3a37897e89f03a4c0033b2112880b22787c96f21717d7c028dbebc696fc98faee
[kubernetes-server-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-server-linux-s390x.tar.gz) | 8f0a37542dc82402a542af6cb11e2831f0dbf16a9fa0594037d5e99efa7196c772352a44f99ba9ec4a78d7bfb350e6cbcc5404756a993bc5a5f6544b17e068ca

### Node Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-node-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-node-linux-amd64.tar.gz) | cb02d70d26484d858f0bb515bcf6e24c7d36d7dfa896358c06f1b260736bff359423c867df192f407d0f1b8e5d829be11e20f729de629506b0319b9a421872c7
[kubernetes-node-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-node-linux-arm.tar.gz) | c6cff49b74227c035bbbcc730c1e0f165ad6c6db992f1e0a02a0a581741ca03a638a2b2012f5eba01a847dbaf26692f23f935055e5a28354bd2b13935871c262
[kubernetes-node-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-node-linux-arm64.tar.gz) | 26d91660351ab07be1cb52ddb99087dccc8df31659d214634d81667a08eeb69b8c2552481cac0ab1a3f9c83152b6384ccaffe2d0b728cf7e2acbed272d166ac4
[kubernetes-node-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-node-linux-ppc64le.tar.gz) | c31c8533ff319d7d3a88207ecbc3f318cf54a810d5f3b1092534bf6e5f16e00ca635573f2f2c9cfa1d461b8e545e09821bf512f0ccae4e2c8b99232695b92766
[kubernetes-node-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-node-linux-s390x.tar.gz) | ae7b9726078e650527f4f861b4fff4a0d84e9fcba406944ee0ade44b4dfe803c37223bb4b8074277d82695654d2b1a4b83ac7376ecf7138b8b5718228012d162
[kubernetes-node-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.3/kubernetes-node-windows-amd64.tar.gz) | 21cb251d63bbcda04b0c6ea74badb2c54a8a8b79824440a20543ab300631d463693255c1367eeab0afe583248ad49a43459d405d70ce04e95014a272101fedf4

## Changelog since v1.22.0-alpha.2

## Urgent Upgrade Notes

### (No, really, you MUST read this before you upgrade)

 - CSI migration of AWS EBS volumes requires AWS EBS CSI driver ver. 1.0 that supports  `allowAutoIOPSPerGBIncrease` parameter in StorageClass. ([#101082](https://github.com/kubernetes/kubernetes/pull/101082), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
  - Conformance image is now built with Distroless. Users running Conformance testing should rely on container entrypoint instead of manual invocation to `/run_e2e.sh` or `/gorunner`, as they are now deprecated and will be removed in 1.25 release. Invoking `ginkgo` and `e2e.test` are still supported through overriding entrypoint (docker) or defining container `spec.command` (kubernetes) ([#99178](https://github.com/kubernetes/kubernetes/pull/99178), [@wilsonehusin](https://github.com/wilsonehusin)) [SIG Release and Testing]
  - Default StreamingProxyRedirects to disabled. If there is a >= 2 version skew between master
  and nodes, and the old nodes were enabling --redirect-container-streaming, this will break
  them. In this case, the StreamingProxyRedirects can still be manually enabled. ([#101647](https://github.com/kubernetes/kubernetes/pull/101647), [@pacoxu](https://github.com/pacoxu)) [SIG API Machinery and Node]
  - Kubeadm: remove the automatic detection and matching of cgroup drivers for Docker. For new clusters if you have not configured the cgroup driver explicitly you might get a failure in the kubelet on driver mismatch (kubeadm clusters should be using the "systemd" driver). Also remove the "IsDockerSystemdCheck" preflight check (warning) that checks if the Docker cgroup driver is set to "systemd". Ideally such detection / coordination should be on the side of CRI implementers and the kubelet (tracked here https://github.com/kubernetes/kubernetes/issues/99808). Please see the following page on how to configure cgroup drivers with kubeadm manually: https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/configure-cgroup-driver/ ([#99647](https://github.com/kubernetes/kubernetes/pull/99647), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
  - The CSIMigrationVSphereComplete feature flag is removed. InTreePluginvSphereUnregister will be the way moving forward. ([#101272](https://github.com/kubernetes/kubernetes/pull/101272), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG API Machinery, Node and Storage]
 
## Changes by Kind

### Deprecation

- E2e.test: removed the  `--viper-config` flag.  If you were previously using this to pass flags to `e2e.test` via a file, you will need to pass them directly on the command line, e.g. `e2e.test --e2e-output-dir` ([#102598](https://github.com/kubernetes/kubernetes/pull/102598), [@dims](https://github.com/dims)) [SIG Testing]
- Remove support for the Service topologyKeys field (alpha) and the kube-proxy implementation of it.  This field was deprecated several cycles ago.  This functionality is replaced by the combination of automatic topology hints per-endpoint (alpha) and the Service internalTrafficPolicy field (alpha). ([#102412](https://github.com/kubernetes/kubernetes/pull/102412), [@andrewsykim](https://github.com/andrewsykim)) [SIG API Machinery, Apps and Network]

### API Change

- --ssh-user and --ssh-key options are removed.  They only functioned on GCE, and only in-tree.  Use the apiserver network proxy instead. ([#102297](https://github.com/kubernetes/kubernetes/pull/102297), [@deads2k](https://github.com/deads2k)) [SIG API Machinery, Cloud Provider and Testing]
- Enable MaxSurge for DS by default ([#101742](https://github.com/kubernetes/kubernetes/pull/101742), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Apps and Testing]
- Introduce minReadySeconds api to the StatefulSets. ([#100842](https://github.com/kubernetes/kubernetes/pull/100842), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG API Machinery, Apps and Testing]
- Kube-controller-manger: the `--horizontal-pod-autoscaler-use-rest-clients`  flag and Heapster support in the horizontal pod autoscaler, deprecated since 1.12, is removed. ([#90368](https://github.com/kubernetes/kubernetes/pull/90368), [@serathius](https://github.com/serathius)) [SIG API Machinery, Apps, Autoscaling, Cloud Provider and Instrumentation]
- The deprecated flag --algorithm-provider has been removed from kube-scheduler. Use instead ComponentConfig to configure the set of enabled plugins ([#102239](https://github.com/kubernetes/kubernetes/pull/102239), [@Haleygo](https://github.com/Haleygo)) [SIG Cloud Provider and Scheduling]

### Feature

- An audit log entry will be generated when a validating admission webhook is failing open. ([#92739](https://github.com/kubernetes/kubernetes/pull/92739), [@cnphil](https://github.com/cnphil)) [SIG API Machinery]
- BoundServiceAccountTokenVolume is GA. The feature gate is going to be removed in 1.23. ([#101992](https://github.com/kubernetes/kubernetes/pull/101992), [@zshihang](https://github.com/zshihang)) [SIG Auth, Cloud Provider and Testing]
- Graduate prefer nominated node to beta ([#102201](https://github.com/kubernetes/kubernetes/pull/102201), [@chendave](https://github.com/chendave)) [SIG Scheduling]
- Introduce a feature gate DisableCloudProviders allowing to disable cloud-provider initialization in KAPI, KCM and kubelet.
  
  DisableCloudProviders FeatureGate is currently in Alpha, which means is currently disabled by default. Once the FeatureGate moves to beta, in-tree cloud providers would be disabled by default, and a user won't be able to specify --cloud-provider=<aws|openstack|azure|gcp|vsphere> anymore to any of KCM, KAPI or kubelet. Only a --cloud-provider=external would be allowed. CCM would have to run out-of-tree with CSI. ([#100136](https://github.com/kubernetes/kubernetes/pull/100136), [@Danil-Grigorev](https://github.com/Danil-Grigorev)) [SIG API Machinery, Cloud Provider, Instrumentation and Node]
- Kubeadm: add a new field "skipPhases" to v1beta3 InitConfiguration and JoinConfiguration that can contain a list of phases to skip during "kubeadm init" and "kubeadm join". The flag "--skip-phases" takes precedence over this field. ([#101923](https://github.com/kubernetes/kubernetes/pull/101923), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: add the RootlessControlPlane kubeadm specific feature gate (Alpha in 1.22, disabled by default).
  It can be used to enable an experimental feature that makes the control plane component static Pod containers 
  for kube-apiserver, kube-controller-manager, kube-scheduler and etcd to run as a non-root users. ([#102158](https://github.com/kubernetes/kubernetes/pull/102158), [@vinayakankugoyal](https://github.com/vinayakankugoyal)) [SIG Cluster Lifecycle]
- Kubeadm: during "kubeadm init/join/upgrade", always default the "cgroupDriver" value in the KubeletConfiguration to "systemd", unless the user was explicit about the value. See https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/configure-cgroup-driver/ for more details. ([#102133](https://github.com/kubernetes/kubernetes/pull/102133), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle]
- Kubectl: add "LAST RESTART" column to "kubectl get pods" output ([#100142](https://github.com/kubernetes/kubernetes/pull/100142), [@Ethyling](https://github.com/Ethyling)) [SIG CLI]
- NetworkPolicy validation framework support for windows ([#98077](https://github.com/kubernetes/kubernetes/pull/98077), [@jayunit100](https://github.com/jayunit100)) [SIG Auth, Network and Testing]
- New feature gate "ExpandedDNSConfig" is now available. This feature allows Kubernetes to have expanded DNS configuration. ([#100651](https://github.com/kubernetes/kubernetes/pull/100651), [@gjkim42](https://github.com/gjkim42)) [SIG Apps, Network and Node]
- Promote CronJobControllerV2 flag to GA, with removal in 1.23 ([#102529](https://github.com/kubernetes/kubernetes/pull/102529), [@soltysh](https://github.com/soltysh)) [SIG Apps]
- Promote Cronjobs storage version to batch/v1 ([#102363](https://github.com/kubernetes/kubernetes/pull/102363), [@mengjiao-liu](https://github.com/mengjiao-liu)) [SIG API Machinery and Testing]
- Scheduler now registers event handlers dynamically. ([#101394](https://github.com/kubernetes/kubernetes/pull/101394), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling and Testing]
- Some of the in-tree storage drivers indicate support for the MetricsProvider interface, but fail to configure this for BlockMode volumes. With a recent change, Kubelet will call GetMetrics() for BlockMode volumes, and the in-tree drivers that miss the support cause a Go panic. Now the in-tree storage drivers that support BlockMode volumes, will return the Capacity of the volume in the GetMetrics() call. ([#101587](https://github.com/kubernetes/kubernetes/pull/101587), [@nixpanic](https://github.com/nixpanic)) [SIG Instrumentation, Node, Storage and Testing]
- Support FakeClientset match subresource ([#100939](https://github.com/kubernetes/kubernetes/pull/100939), [@wzshiming](https://github.com/wzshiming)) [SIG API Machinery and Testing]
- Update the Debian images to pick up CVE fixes in the base images:
  - Update the `debian-base` image to v1.7.0
  - Update the `debian-iptables` image to v1.6.1 ([#102302](https://github.com/kubernetes/kubernetes/pull/102302), [@xmudrii](https://github.com/xmudrii)) [SIG API Machinery, Release and Testing]
- Update the setcap image to buster-v2.0.1 ([#102377](https://github.com/kubernetes/kubernetes/pull/102377), [@xmudrii](https://github.com/xmudrii)) [SIG Release]
- Watch requests are now handled throttled by priority and fairness filter in kube-apiserver ([#102171](https://github.com/kubernetes/kubernetes/pull/102171), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery]
- [kubectl] Enable using protocol buffers to request Metrics API ([#102039](https://github.com/kubernetes/kubernetes/pull/102039), [@serathius](https://github.com/serathius)) [SIG CLI]

### Bug or Regression

- Aggregate errors when putting vmss ([#98350](https://github.com/kubernetes/kubernetes/pull/98350), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix a bug on the endpoint slices mirroring controller where endpoint NotReadyAddresses were mirrored as Ready to the corresponding EndpointSlice ([#102683](https://github.com/kubernetes/kubernetes/pull/102683), [@aojea](https://github.com/aojea)) [SIG Apps and Network]
- Fix a bug that a preemptor pod may exist as a phantom in the scheduler. ([#102498](https://github.com/kubernetes/kubernetes/pull/102498), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling]
- Fix nulls are handles of array's and objects in json patches. ([#102467](https://github.com/kubernetes/kubernetes/pull/102467), [@pacoxu](https://github.com/pacoxu)) [SIG API Machinery, Apps, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Storage and Testing]
- Fix resource enforcement when using systemd cgroup driver ([#102147](https://github.com/kubernetes/kubernetes/pull/102147), [@kolyshkin](https://github.com/kolyshkin)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Storage and Testing]
- Fix runtime container status for post start hook error ([#100608](https://github.com/kubernetes/kubernetes/pull/100608), [@pacoxu](https://github.com/pacoxu)) [SIG Node]
- Fixed an issue blocking azure auth to prompt to device code authentication flow when refresh token expires. ([#102063](https://github.com/kubernetes/kubernetes/pull/102063), [@tdihp](https://github.com/tdihp)) [SIG API Machinery and Auth]
- Fixed garbage collection of dangling VolumeAttachments for PersistentVolumes migrated to CSI on startup of kube-controller-manager. ([#102176](https://github.com/kubernetes/kubernetes/pull/102176), [@timebertt](https://github.com/timebertt)) [SIG Apps and Storage]
- Fixed kubelet runtime panic for timed-out portforward streams. ([#102489](https://github.com/kubernetes/kubernetes/pull/102489), [@saschagrunert](https://github.com/saschagrunert)) [SIG API Machinery and Node]
- Fixed very rare volume corruption when a pod is deleted while kubelet is offline. ([#102059](https://github.com/kubernetes/kubernetes/pull/102059), [@jsafrane](https://github.com/jsafrane)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Storage]
- Fixes issue with websocket-based watches of Service objects not closing correctly on timeout ([#102539](https://github.com/kubernetes/kubernetes/pull/102539), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]
- Kubeadm: when using a custom image repository for CoreDNS kubeadm now will append the "coredns"  image name instead of  "coredns/coredns", thus restoring the behaviour existing before the v1.21 release. Users who rely on nested folder for the coredns image should set the "clusterConfiguration.dns.imageRepository" value including the nested path name (e.g using "registry.company.xyz/coredns" will force kubeadm to use "registry.company.xyz/coredns/coredns" image). No action is needed if using the default registry (k8s.gcr.io). ([#102502](https://github.com/kubernetes/kubernetes/pull/102502), [@ykakarap](https://github.com/ykakarap)) [SIG Cluster Lifecycle]
- Retry FibreChannel devices cleanup after error to ensure FC device is detached before it can be used on another node. ([#101862](https://github.com/kubernetes/kubernetes/pull/101862), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- ServiceOwnsFrontendIP shouldn't report error when the public IP doesn't match ([#102516](https://github.com/kubernetes/kubernetes/pull/102516), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Support correct sorting for cpu, memory, storage, ephemeral-storage, hugepages, and attachable-volumes ([#100435](https://github.com/kubernetes/kubernetes/pull/100435), [@lauchokyip](https://github.com/lauchokyip)) [SIG CLI and Testing]
- Update klog v2.9.0 ([#102332](https://github.com/kubernetes/kubernetes/pull/102332), [@pacoxu](https://github.com/pacoxu)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Storage]
- Update kube-proxy base image `debian-iptables` to v1.6.2 to pickup https://github.com/kubernetes/release/pull/2106
  - debian-iptables: select nft mode if ntf lines > legacy lines, matching https://github.com/kubernetes-sigs/iptables-wrappers/ ([#102590](https://github.com/kubernetes/kubernetes/pull/102590), [@BenTheElder](https://github.com/BenTheElder)) [SIG Network and Testing]
- We no longer allow the cluster operator to delete any "suggested" priority & fairness bootstrap configuration object, 
  If a cluster operator removes a suggested configuration, it will be restored by the apiserver. ([#102067](https://github.com/kubernetes/kubernetes/pull/102067), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]

### Other (Cleanup or Flake)

- Allow CSI drivers to just run offline expansion tests ([#102665](https://github.com/kubernetes/kubernetes/pull/102665), [@gnufied](https://github.com/gnufied)) [SIG Storage and Testing]
- Changed buildmode of non static Kubernetes binaries to produce position independent executables (PIE). ([#102323](https://github.com/kubernetes/kubernetes/pull/102323), [@saschagrunert](https://github.com/saschagrunert)) [SIG Release and Security]
- Clarified the description of a test in the e2e suite that mentions "SCTP" but is
  actually intended to be testing the behavior of network plugins that *don't*
  implement SCTP. ([#102509](https://github.com/kubernetes/kubernetes/pull/102509), [@danwinship](https://github.com/danwinship)) [SIG Network and Testing]
- Disable log sampling when using json logging format ([#102620](https://github.com/kubernetes/kubernetes/pull/102620), [@serathius](https://github.com/serathius)) [SIG API Machinery, Cluster Lifecycle and Instrumentation]
- Improve func ToSelectableFields' performance for event ([#102461](https://github.com/kubernetes/kubernetes/pull/102461), [@goodluckbot](https://github.com/goodluckbot)) [SIG API Machinery and Node]
- Migrate some log messages to structured logging in pkg/volume/plugins.go. ([#101510](https://github.com/kubernetes/kubernetes/pull/101510), [@huchengze](https://github.com/huchengze)) [SIG Storage]
- Update CNI plugins to v0.9.1 ([#102328](https://github.com/kubernetes/kubernetes/pull/102328), [@lentzi90](https://github.com/lentzi90)) [SIG Cloud Provider, Network, Node and Testing]
- Updated pause image to version 3.5, which now runs per default as pseudo user and group `65535:65535`.
  This does not have any effect on remote container runtimes like CRI-O and containerd, which setup the pod
  sandbox user and group on their own. ([#100292](https://github.com/kubernetes/kubernetes/pull/100292), [@saschagrunert](https://github.com/saschagrunert)) [SIG CLI, Cloud Provider, Cluster Lifecycle, Node and Testing]
- Upgrade functionality of `kubectl kustomize` as described at https://github.com/kubernetes-sigs/kustomize/releases/tag/kustomize%2Fv4.1.3 ([#102193](https://github.com/kubernetes/kubernetes/pull/102193), [@gautierdelorme](https://github.com/gautierdelorme)) [SIG CLI]

## Dependencies

### Added
- github.com/checkpoint-restore/go-criu/v5: [v5.0.0](https://github.com/checkpoint-restore/go-criu/v5/tree/v5.0.0)
- github.com/frankban/quicktest: [v1.11.3](https://github.com/frankban/quicktest/tree/v1.11.3)
- github.com/go-kit/log: [v0.1.0](https://github.com/go-kit/log/tree/v0.1.0)
- github.com/jpillora/backoff: [v1.0.0](https://github.com/jpillora/backoff/tree/v1.0.0)

### Changed
- github.com/alecthomas/units: [c3de453 → f65c72e](https://github.com/alecthomas/units/compare/c3de453...f65c72e)
- github.com/cilium/ebpf: [v0.2.0 → v0.5.0](https://github.com/cilium/ebpf/compare/v0.2.0...v0.5.0)
- github.com/containerd/console: [v1.0.1 → v1.0.2](https://github.com/containerd/console/compare/v1.0.1...v1.0.2)
- github.com/coreos/go-systemd/v22: [v22.1.0 → v22.3.1](https://github.com/coreos/go-systemd/v22/compare/v22.1.0...v22.3.1)
- github.com/evanphx/json-patch: [v4.9.0+incompatible → v4.11.0+incompatible](https://github.com/evanphx/json-patch/compare/v4.9.0...v4.11.0)
- github.com/go-logfmt/logfmt: [v0.4.0 → v0.5.0](https://github.com/go-logfmt/logfmt/compare/v0.4.0...v0.5.0)
- github.com/godbus/dbus/v5: [v5.0.3 → v5.0.4](https://github.com/godbus/dbus/v5/compare/v5.0.3...v5.0.4)
- github.com/google/cadvisor: [v0.39.0 → v0.39.2](https://github.com/google/cadvisor/compare/v0.39.0...v0.39.2)
- github.com/google/go-cmp: [v0.5.2 → v0.5.4](https://github.com/google/go-cmp/compare/v0.5.2...v0.5.4)
- github.com/julienschmidt/httprouter: [v1.2.0 → v1.3.0](https://github.com/julienschmidt/httprouter/compare/v1.2.0...v1.3.0)
- github.com/kr/pretty: [v0.2.0 → v0.2.1](https://github.com/kr/pretty/compare/v0.2.0...v0.2.1)
- github.com/moby/sys/mountinfo: [v0.4.0 → v0.4.1](https://github.com/moby/sys/mountinfo/compare/v0.4.0...v0.4.1)
- github.com/mwitkow/go-conntrack: [cc309e4 → 2f06839](https://github.com/mwitkow/go-conntrack/compare/cc309e4...2f06839)
- github.com/opencontainers/runc: [v1.0.0-rc93 → v1.0.0-rc95](https://github.com/opencontainers/runc/compare/v1.0.0-rc93...v1.0.0-rc95)
- github.com/opencontainers/runtime-spec: [e6143ca → 1c3f411](https://github.com/opencontainers/runtime-spec/compare/e6143ca...1c3f411)
- github.com/prometheus/common: [v0.10.0 → v0.26.0](https://github.com/prometheus/common/compare/v0.10.0...v0.26.0)
- github.com/spf13/jwalterweatherman: [v1.1.0 → v1.0.0](https://github.com/spf13/jwalterweatherman/compare/v1.1.0...v1.0.0)
- golang.org/x/sys: a50acf3 → d19ff85
- golang.org/x/text: v0.3.4 → v0.3.6
- k8s.io/klog/v2: v2.8.0 → v2.9.0
- k8s.io/utils: 67b214c → da69540
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.15 → v0.0.19
- sigs.k8s.io/kustomize/api: v0.8.8 → v0.8.10
- sigs.k8s.io/kustomize/cmd/config: v0.9.10 → v0.9.12
- sigs.k8s.io/kustomize/kustomize/v4: v4.1.2 → v4.1.3
- sigs.k8s.io/kustomize/kyaml: v0.10.17 → v0.10.20

### Removed
- github.com/agnivade/levenshtein: [v1.0.1](https://github.com/agnivade/levenshtein/tree/v1.0.1)
- github.com/alecthomas/template: [fb15b89](https://github.com/alecthomas/template/tree/fb15b89)
- github.com/andreyvit/diff: [c7f18ee](https://github.com/andreyvit/diff/tree/c7f18ee)
- github.com/checkpoint-restore/go-criu/v4: [v4.1.0](https://github.com/checkpoint-restore/go-criu/v4/tree/v4.1.0)
- github.com/go-kit/kit: [v0.9.0](https://github.com/go-kit/kit/tree/v0.9.0)
- github.com/go-openapi/analysis: [v0.19.5](https://github.com/go-openapi/analysis/tree/v0.19.5)
- github.com/go-openapi/errors: [v0.19.2](https://github.com/go-openapi/errors/tree/v0.19.2)
- github.com/go-openapi/loads: [v0.19.4](https://github.com/go-openapi/loads/tree/v0.19.4)
- github.com/go-openapi/runtime: [v0.19.4](https://github.com/go-openapi/runtime/tree/v0.19.4)
- github.com/go-openapi/spec: [v0.19.5](https://github.com/go-openapi/spec/tree/v0.19.5)
- github.com/go-openapi/strfmt: [v0.19.5](https://github.com/go-openapi/strfmt/tree/v0.19.5)
- github.com/go-openapi/validate: [v0.19.8](https://github.com/go-openapi/validate/tree/v0.19.8)
- github.com/gobuffalo/here: [v0.6.0](https://github.com/gobuffalo/here/tree/v0.6.0)
- github.com/kr/logfmt: [b84e30a](https://github.com/kr/logfmt/tree/b84e30a)
- github.com/markbates/pkger: [v0.17.1](https://github.com/markbates/pkger/tree/v0.17.1)
- github.com/tidwall/pretty: [v1.0.0](https://github.com/tidwall/pretty/tree/v1.0.0)
- github.com/vektah/gqlparser: [v1.1.2](https://github.com/vektah/gqlparser/tree/v1.1.2)
- go.mongodb.org/mongo-driver: v1.1.2
- k8s.io/heapster: v1.2.0-beta.1



# v1.22.0-alpha.2


## Downloads for v1.22.0-alpha.2

### Source Code

filename | sha512 hash
-------- | -----------
[kubernetes.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes.tar.gz) | 39d5177271e744058585c4b924ff91e4df654db81257a4710b77a055ac6033c8d6414772a4c42e3ec7f568ac5c9691c53225a13a68610aa0b07c3bcaf252fe4c
[kubernetes-src.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-src.tar.gz) | d9832ab5ba568f89ffb7e9bfef3dd0baee69c5a29bc34e2f8f83fef08f13575e4982409dba422b912245655b326565f9e71e523bcbd391b97fd385ae7e4debaa

### Client Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-client-darwin-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-client-darwin-amd64.tar.gz) | d007c403a9586a0047db4abc8766845aa501798524a259902a3a3e5d43928a819b9857ef4b49632384139e12d3b0e0c0cbf2966a5067e9e29496d4bf14a2ea24
[kubernetes-client-darwin-arm64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-client-darwin-arm64.tar.gz) | 75ce76788e5bebcd6c06a8cc804c39edccaf42941dfd35cea331eb86393918fb6addef2bf507b78d9dac6eb3627568c281404a5fd899fb396052ff9658dc3f70
[kubernetes-client-linux-386.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-client-linux-386.tar.gz) | 9d590915534c1fe3d69c1e0df7b16c6668e52be32a3649214a0a4940a8ed1565efe2c300a1c7aa02c9605be8e829fd9a75229d2b0a9a0765f3ce16b6ad68f4b1
[kubernetes-client-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-client-linux-amd64.tar.gz) | 5e44d189b32a61b3f060a5ec13207cea526c7fedbc42967915e6b50f106ac862c13560bb15066bed3134407621ae506c18297c7b3ea2f561fb20a97ac02215cc
[kubernetes-client-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-client-linux-arm.tar.gz) | d8d446133b14f11da9f33a20c6700d23d2616b4d6cf750e8074526b8442b4e0e437b20444fc583f4097c8b064966a4a1e52fb2e01096e2c94ec4e05ef2d4b48e
[kubernetes-client-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-client-linux-arm64.tar.gz) | 20db0fee191a027885b9a12615732b40e88c148f04343f56e67dfa5a12e08a51238c6e93aed05685afd6b203dc3f1961c6db4096ba867caf299d5d0a190a91d6
[kubernetes-client-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-client-linux-ppc64le.tar.gz) | b642a040ac656c609be3191af20f5b3142d20b1d39846e3052402e99bb5fca9211e4225cb775a9ec19b9cf7e47754ece813a7d367d9c911be18a1ea5584cf178
[kubernetes-client-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-client-linux-s390x.tar.gz) | 161756afb0b040dd5134d91f4982dd0233f3e4fec31375dd6b2f515c12f6fc0c7237a0c8283bf2a83e147df69403b35c3d9bbe7b872779dd5b2e43ef5c8693f8
[kubernetes-client-windows-386.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-client-windows-386.tar.gz) | 3956a25f75a29f23a559a335a0629299a083143db1ccad6db2ff76c27ead72ad25a5db81b558225a530749ffc58749342079f80c5af4f0134553b6de05f60a5b
[kubernetes-client-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-client-windows-amd64.tar.gz) | 29ba410e9d600b92ec02f284a545045d4e3b1e6c247fc5db64c2a8536108456389986efbdb762faba6509b1b50e9bbc3638d2dca19577de79b0de34ad749e410

### Server Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-server-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-server-linux-amd64.tar.gz) | 0fad8691f0c72e0b4b09e5da9b806353f1a5c48c3b38c90674d44e673daa77ed85e727434ba9cbe2717ca65005059af17fb7b7db4d452aa67fef7cf2395da738
[kubernetes-server-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-server-linux-arm.tar.gz) | e6d8e4dbcc5e5790114834b8aa5a9bfc5b1c18c4b16cb043f3fd409c22c8b2ccfdb165357e584f650321c5c07ca5aae405f70da65efa32f5dffbeb25ebc22c42
[kubernetes-server-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-server-linux-arm64.tar.gz) | 4f0c46ad6a504ea0b7607175603e61530496d29759f27e6e9dac8b7bb923f8920ed6dd2afb5d709f2f96850145252d4dd702bb77254791639cfb33648f3b1f04
[kubernetes-server-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-server-linux-ppc64le.tar.gz) | 58cb325bd7470972df7d286a7c160f732f261ce4858882f99cc5ab91ba43f86d1cf1294651f61ba1416c17ff91abfe178dfbb1c264716029d58b94f595dc734d
[kubernetes-server-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-server-linux-s390x.tar.gz) | 993806a9f7404365ab6e8a4017a5c5dece028c9f8c376498c196dda9bb885ecbceaad5498f43bea8d1309707216ae4173dc8aa69151ad304e5f1993be1f7f6dc

### Node Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-node-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-node-linux-amd64.tar.gz) | 77c1fd98ac2e8a665bffae60bdc66f1b5fc29482d29f58b4d5705b43478fc536885e6634ffd2e8a18ff0ea589a15a2df67ba86ede2025a697019030bd7893bbc
[kubernetes-node-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-node-linux-arm.tar.gz) | 839ec50c2438279fae2b52efa985556a7c4ba090c8296d56ae8623b3b7123cb6c4b0a656083cc43463e57fdb3d8bae2609196879061aa806aac3a65562c02e40
[kubernetes-node-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-node-linux-arm64.tar.gz) | 3568f19f2cde2da63e5897a8f206a475f86c41f273dac4eb1e31416945d112c6d00ce74e4159732d3805cbe093c94ec53c573227f41ab873c6698023b473b2f5
[kubernetes-node-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-node-linux-ppc64le.tar.gz) | f9de072d40b9a354785ae1dbc182ddcb431e0c4e00fe8f4c56e2b5ff2062845e0c740e7efd4b9697bca9848b0808ad01f20817f84d5c5d5c9c78e52be7962243
[kubernetes-node-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-node-linux-s390x.tar.gz) | fc4b31c68367e938991a5ca4d9df9c38950939ada6f6c0dd6a827d43d5f003b20fdea25d34a213853d53d0a933e4715425f37668de7e110ee0722cb866fa94bd
[kubernetes-node-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.2/kubernetes-node-windows-amd64.tar.gz) | 8dbee9ebf915c645ce199d7f190323ffd71b810f2cec2e1dee8d35948994aa7d08ebc9a82ab083f1eb83476ae104d4a63b4bc258ecbcc9ab3f158d56f179d7a8

## Changelog since v1.22.0-alpha.1

## Urgent Upgrade Notes

### (No, really, you MUST read this before you upgrade)

 - Intree volume plugin scaleio support is been completely removed from Kubernetes. ([#101685](https://github.com/kubernetes/kubernetes/pull/101685), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG API Machinery, Node and Storage]
  - Newly provisioned PVs by Azure disk will no longer have the beta FailureDomain label. Azure disk volume plugin will start to have GA topology label instead. ([#101534](https://github.com/kubernetes/kubernetes/pull/101534), [@kassarl](https://github.com/kassarl)) [SIG Cloud Provider and Storage]
  - Scheduler's CycleState now embeds internal read/write locking inside its Read() and Write() functions. Meanwhile, Lock() and Unlock() function are removed.
  
  scheduler plugin developers are now required to remove CycleState#Lock() and CycleState#Unlock(). Just simply use Read() and Write() as they're natively thread-safe now. ([#101542](https://github.com/kubernetes/kubernetes/pull/101542), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling and Storage]
 
## Changes by Kind

### Deprecation

- Controller-manager: the following flags have no effect and would be removed in v1.24:
  - `--port`
  - `--address`
  The insecure port flags `--port` may only be set to 0 now.
  
  In addtion, please be careful that:
  - controller-manager MUST start with `--authorization-kubeconfig` and `--authentication-kubeconfig` correctly set to get authentication/authorization working.
  - liveness/readiness probes to controller-manager MUST use HTTPS now, and the default port has been changed to 10257.
  - Applications that fetch metrics from controller-manager should use a dedicated service account which is allowed to access nonResourceURLs `/metrics`. ([#96216](https://github.com/kubernetes/kubernetes/pull/96216), [@knight42](https://github.com/knight42)) [SIG API Machinery, Cloud Provider, Instrumentation and Testing]
- Ingress v1beta1 has been deprecated ([#102030](https://github.com/kubernetes/kubernetes/pull/102030), [@aojea](https://github.com/aojea)) [SIG CLI, Network and Testing]
- Kubeadm: remove the deprecated `--csr-only` and `--csr-dir` flags from "kubeadm init phase certs". Deprecate the same flags under "kubeadm certs renew". In both cases the command "kubeadm certs generate-csr" should be used instead. ([#102108](https://github.com/kubernetes/kubernetes/pull/102108), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: remove the ClusterStatus API from v1beta3 and its management in the kube-system/kubeadm-config ConfigMap. This method of keeping track of what API endpoints exists in the cluster was replaced (in a prior release) by a method to annotate the etcd Pods that kubeadm creates in "stacked etcd" clusters. The following CLI sub-phases are deprecated and are now a NO-OP: for " kubeadm join": "control-plane-join/update-status", for "kubeadm reset": "update-cluster-status". Unless you are using these phases explicitly, you should not be affected. ([#101915](https://github.com/kubernetes/kubernetes/pull/101915), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: remove the deprecated command "kubeadm alpha kubeconfig". Please use "kubeadm kubeconfig" instead. ([#101938](https://github.com/kubernetes/kubernetes/pull/101938), [@knight42](https://github.com/knight42)) [SIG Cluster Lifecycle]
- Kubeadm: remove the deprecated command 'kubeadm config view'. A replacement for this command is 'kubectl get cm -n kube-system kubeadm-config -o=jsonpath="{.data.ClusterConfiguration}"' ([#102071](https://github.com/kubernetes/kubernetes/pull/102071), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: remove the deprecated flag `--image-pull-timeout` for 'kubeadm upgrade apply' command ([#102093](https://github.com/kubernetes/kubernetes/pull/102093), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Kubeadm: remove the deprecated flag `--insecure-port` from the kube-apiserver manifest that kubeadm manages. The flag had no effect since 1.20, since the insecure serving of the component was disabled in the same version. ([#102121](https://github.com/kubernetes/kubernetes/pull/102121), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle]
- Kubeadm: remove the deprecated hyperkube image support in v1beta3. This implies removal of ClusterConfiguration.UseHyperKubeImage. ([#101537](https://github.com/kubernetes/kubernetes/pull/101537), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: remove the field ClusterConfiguration.DNS.Type in v1beta3 since CoreDNS is the only supported DNS type. ([#101547](https://github.com/kubernetes/kubernetes/pull/101547), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- `storageos`,`quobyte` and `flocker` storage volume plugins are deprecated and will be removed in a later release. ([#101773](https://github.com/kubernetes/kubernetes/pull/101773), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG Storage]

### API Change

- Add alpha support for HostProcess containers on Windows ([#99576](https://github.com/kubernetes/kubernetes/pull/99576), [@marosset](https://github.com/marosset)) [SIG API Machinery, Apps, Node, Testing and Windows]
- Add three metrics to job controller to monitor if Job works in a healthy condition.
  IndexedJob promoted to Beta ([#101292](https://github.com/kubernetes/kubernetes/pull/101292), [@AliceZhang2016](https://github.com/AliceZhang2016)) [SIG Apps, Instrumentation and Testing]
- Corrected the documentation for escaping dollar signs in a container's env, command and args property. ([#101916](https://github.com/kubernetes/kubernetes/pull/101916), [@MartinKanters](https://github.com/MartinKanters)) [SIG Apps]
- Omit comparison with boolean constant ([#101523](https://github.com/kubernetes/kubernetes/pull/101523), [@GreenApple10](https://github.com/GreenApple10)) [SIG CLI and Cloud Provider]
- Pod Affinity NamespaceSelector and the associated CrossNamespaceAffinity quota scope graduated to beta ([#101496](https://github.com/kubernetes/kubernetes/pull/101496), [@ahg-g](https://github.com/ahg-g)) [SIG API Machinery, Apps and Testing]
- V1.Node .status.images[].names is now optional ([#102159](https://github.com/kubernetes/kubernetes/pull/102159), [@roycaihw](https://github.com/roycaihw)) [SIG Apps and Node]

### Feature

- Added BinaryData description to kubectl describe ([#100568](https://github.com/kubernetes/kubernetes/pull/100568), [@lauchokyip](https://github.com/lauchokyip)) [SIG CLI]
- Feat: change parittion style to GPT on Windows ([#101412](https://github.com/kubernetes/kubernetes/pull/101412), [@andyzhangx](https://github.com/andyzhangx)) [SIG Storage and Windows]
- Improve logging of APIService availability changes in kube-apiserver. ([#101420](https://github.com/kubernetes/kubernetes/pull/101420), [@sttts](https://github.com/sttts)) [SIG API Machinery]
- Kubeadm: add the RootlessControlPlane kubeadm specific feature gate (Alpha in 1.22, disabled by default).
  It can be used to enable an experimental feature that makes the control plane component static Pod containers 
  for kube-apiserver, kube-controller-manager, kube-scheduler and etcd to run as a non-root users. ([#102158](https://github.com/kubernetes/kubernetes/pull/102158), [@vinayakankugoyal](https://github.com/vinayakankugoyal)) [SIG Cluster Lifecycle]
- Kubeadm: set the seccompProfile to runtime/default in the PodSecurityContext of the  control-plane components that run as static Pods. ([#100234](https://github.com/kubernetes/kubernetes/pull/100234), [@vinayakankugoyal](https://github.com/vinayakankugoyal)) [SIG Cluster Lifecycle]
- Kubernetes is now built with Golang 1.16.4 ([#101809](https://github.com/kubernetes/kubernetes/pull/101809), [@justaugustus](https://github.com/justaugustus)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Metrics server nanny has now poll period set to 30s (previously 5 minutes) to allow faster scaling of metrics server. ([#101869](https://github.com/kubernetes/kubernetes/pull/101869), [@olagacek](https://github.com/olagacek)) [SIG Cloud Provider and Instrumentation]
- New metrics: `apiserver_kube_aggregator_x509_missing_san_total` and `apiserver_webhooks_x509_missing_san_total`. This metric measures a number of connections to webhooks/aggregated API servers that use certificates without Subject Alternative Names. It being non-zero is a warning sign that these connections will stop functioning in the future since Golang is going to deprecate x509 certificate subject Common Names for server hostname verification. ([#95396](https://github.com/kubernetes/kubernetes/pull/95396), [@stlaz](https://github.com/stlaz)) [SIG API Machinery, Auth and Instrumentation]
- Node Problem Detector is now available for GCE Windows nodes. ([#101539](https://github.com/kubernetes/kubernetes/pull/101539), [@jeremyje](https://github.com/jeremyje)) [SIG Cloud Provider, Node and Windows]
- Secret values are now masked by default in kubectl diff output. ([#96084](https://github.com/kubernetes/kubernetes/pull/96084), [@loozhengyuan](https://github.com/loozhengyuan)) [SIG CLI]
- The `WarningHeader` feature is now GA and is unconditionally enabled. The `apiserver_requested_deprecated_apis` metric has graduated to stable status. The `WarningHeader` feature-gate is no longer operative and will be removed in v1.24. ([#100754](https://github.com/kubernetes/kubernetes/pull/100754), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Instrumentation and Testing]
- Warnings for use of deprecated and known-bad values in pod specs are now sent ([#101688](https://github.com/kubernetes/kubernetes/pull/101688), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Auth]
- You can use this Builder function to create events Field Selector ([#101817](https://github.com/kubernetes/kubernetes/pull/101817), [@cndoit18](https://github.com/cndoit18)) [SIG API Machinery and Scalability]

### Failing Test

- Fixes the `should receive events on concurrent watches in same order` conformance test to work properly on clusters that auto-create additional configmaps in namespaces ([#101950](https://github.com/kubernetes/kubernetes/pull/101950), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]
- Resolves an issue with the "ServiceAccountIssuerDiscovery should support OIDC discovery" conformance test failing on clusters which are configured with issuers outside the cluster ([#101589](https://github.com/kubernetes/kubernetes/pull/101589), [@mtaufen](https://github.com/mtaufen)) [SIG Auth and Testing]

### Bug or Regression

- Added jitter factor to lease controller that better smears load on kube-apiserver over time. ([#101652](https://github.com/kubernetes/kubernetes/pull/101652), [@marseel](https://github.com/marseel)) [SIG API Machinery and Scalability]
- Avoid caching the Azure VMSS instances whose network profile is nil ([#100948](https://github.com/kubernetes/kubernetes/pull/100948), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Azure: avoid setting cached Sku when updating VMSS and VMSS instances ([#102005](https://github.com/kubernetes/kubernetes/pull/102005), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix Azure node public IP fetching issues from instance metadata service when the node is part of standard load balancer backend pool. ([#100690](https://github.com/kubernetes/kubernetes/pull/100690), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix EndpointSlice describe panic when an Endpoint doesn't have zone ([#101025](https://github.com/kubernetes/kubernetes/pull/101025), [@tnqn](https://github.com/tnqn)) [SIG CLI]
- Fix kubectl set env or resources not working for initcontainers ([#101669](https://github.com/kubernetes/kubernetes/pull/101669), [@carlory](https://github.com/carlory)) [SIG CLI]
- Fix resource enforcement when using systemd cgroup driver ([#102147](https://github.com/kubernetes/kubernetes/pull/102147), [@kolyshkin](https://github.com/kolyshkin)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Storage and Testing]
- Fix: avoid nil-pointer panic when checking the frontend IP configuration ([#101739](https://github.com/kubernetes/kubernetes/pull/101739), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: delete non existing disk issue ([#102083](https://github.com/kubernetes/kubernetes/pull/102083), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: not tagging static public IP ([#101752](https://github.com/kubernetes/kubernetes/pull/101752), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fixed a bug that `kubectl create configmap` always returns zero exit code when failed. ([#101780](https://github.com/kubernetes/kubernetes/pull/101780), [@nak3](https://github.com/nak3)) [SIG CLI]
- Fixed false-positive uncertain volume attachments, which led to unexpected detachment of CSI migrated volumes ([#101737](https://github.com/kubernetes/kubernetes/pull/101737), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG Apps and Storage]
- Fixed mounting of NFS volumes when IPv6 address is used as a server. ([#101067](https://github.com/kubernetes/kubernetes/pull/101067), [@Elbehery](https://github.com/Elbehery)) [SIG Storage]
- GCE Windows will no longer install Docker on containerd nodes. ([#101747](https://github.com/kubernetes/kubernetes/pull/101747), [@jeremyje](https://github.com/jeremyje)) [SIG Cloud Provider and Windows]
- Kube-proxy log now shows the "Skipping topology aware endpoint filtering since no hints were provided for zone" warning under the right conditions ([#101857](https://github.com/kubernetes/kubernetes/pull/101857), [@dervoeti](https://github.com/dervoeti)) [SIG Network]
- Kubeadm upgrade etcd to 3.4.13-3 ([#100612](https://github.com/kubernetes/kubernetes/pull/100612), [@pacoxu](https://github.com/pacoxu)) [SIG API Machinery, Cloud Provider and Cluster Lifecycle]
- Kubeadm: fix the bug that kubeadm only uses the first hash in caCertHashes to verify the root CA ([#101977](https://github.com/kubernetes/kubernetes/pull/101977), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Kubectl create service now respects namespace flag ([#101005](https://github.com/kubernetes/kubernetes/pull/101005), [@zxh326](https://github.com/zxh326)) [SIG CLI]
- Kubectl wait --for=delete ignores not found error correctly now. ([#96702](https://github.com/kubernetes/kubernetes/pull/96702), [@lingsamuel](https://github.com/lingsamuel)) [SIG CLI and Testing]
- Parsing of cpuset information now properly detects more invalid input such as "1--3" or "10-6" ([#100565](https://github.com/kubernetes/kubernetes/pull/100565), [@lack](https://github.com/lack)) [SIG Node]
- Register/Deregister Targets in chunks for AWS TargetGroup ([#101592](https://github.com/kubernetes/kubernetes/pull/101592), [@M00nF1sh](https://github.com/M00nF1sh)) [SIG Cloud Provider]
- Respect annotation size limit for server-side apply updates to the client-side apply annotation. Also, fix opt-out of this behavior by setting the client-side apply annotation to the empty string. ([#102105](https://github.com/kubernetes/kubernetes/pull/102105), [@julianvmodesto](https://github.com/julianvmodesto)) [SIG API Machinery]
- The conformance tests:
  - Services should serve multiport endpoints from pods
  - Services should serve a basic endpoint from pods
  were only validating the API objects, not performing any validation on the actual Services implementation.
  Those tests now validate that the Services under test are able to forward traffic to the endpoints. ([#101709](https://github.com/kubernetes/kubernetes/pull/101709), [@aojea](https://github.com/aojea)) [SIG Network and Testing]
- When `DisableAcceleratorUsageMetrics` is set, do not collect accelerator metrics using cAdvisor. ([#101712](https://github.com/kubernetes/kubernetes/pull/101712), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev)) [SIG Instrumentation and Node]

### Other (Cleanup or Flake)

- Fake clients now implement a `FakeClient` interface ([#100940](https://github.com/kubernetes/kubernetes/pull/100940), [@markusthoemmes](https://github.com/markusthoemmes)) [SIG API Machinery and Instrumentation]
- Kubeadm: the `CriticalAddonsOnly` toleration has been removed from `kube-proxy` DaemonSet ([#101966](https://github.com/kubernetes/kubernetes/pull/101966), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Metrics Server updated to use 0.4.4 image that doesn't depend on deprecated authorization.k8s.io/v1beta1 subjectaccessreviews API version. ([#101477](https://github.com/kubernetes/kubernetes/pull/101477), [@x13n](https://github.com/x13n)) [SIG Cloud Provider and Instrumentation]
- Migrate proxy/ipvs/proxier.go logs to structured logging ([#97796](https://github.com/kubernetes/kubernetes/pull/97796), [@JornShen](https://github.com/JornShen)) [SIG Network]
- Remove duplicate packet import ([#101187](https://github.com/kubernetes/kubernetes/pull/101187), [@GreenApple10](https://github.com/GreenApple10)) [SIG API Machinery]
- The `VolumeSnapshotDataSource` feature gate that is GA since v1.20 is unconditionally enabled, and can no longer be specified via the `--feature-gates` argument. ([#101531](https://github.com/kubernetes/kubernetes/pull/101531), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG Storage]
- The deprecated CRIContainerLogRotation feature-gate has been removed, since the CRIContainerLogRotation feature graduated to GA in 1.21 and was unconditionally enabled. ([#101578](https://github.com/kubernetes/kubernetes/pull/101578), [@carlory](https://github.com/carlory)) [SIG Node]
- The deprecated RootCAConfigMap feature-gate has been removed, since the RootCAConfigMap feature graduated to GA in 1.21 and was unconditionally enabled. ([#101579](https://github.com/kubernetes/kubernetes/pull/101579), [@carlory](https://github.com/carlory)) [SIG Auth]

## Dependencies

### Added
- github.com/nxadm/tail: [v1.4.4](https://github.com/nxadm/tail/tree/v1.4.4)
- rsc.io/quote/v3: v3.1.0
- rsc.io/sampler: v1.3.0

### Changed
- github.com/containernetworking/cni: [v0.8.0 → v0.8.1](https://github.com/containernetworking/cni/compare/v0.8.0...v0.8.1)
- github.com/golang/mock: [v1.4.4 → v1.4.3](https://github.com/golang/mock/compare/v1.4.4...v1.4.3)
- github.com/onsi/ginkgo: [v1.11.0 → v1.14.0](https://github.com/onsi/ginkgo/compare/v1.11.0...v1.14.0)
- github.com/onsi/gomega: [v1.7.0 → v1.10.1](https://github.com/onsi/gomega/compare/v1.7.0...v1.10.1)
- github.com/stretchr/testify: [v1.6.1 → v1.7.0](https://github.com/stretchr/testify/compare/v1.6.1...v1.7.0)

### Removed
- github.com/hpcloud/tail: [v1.0.0](https://github.com/hpcloud/tail/tree/v1.0.0)
- github.com/thecodeteam/goscaleio: [v0.1.0](https://github.com/thecodeteam/goscaleio/tree/v0.1.0)
- gopkg.in/fsnotify.v1: v1.4.7



# v1.22.0-alpha.1


## Downloads for v1.22.0-alpha.1

### Source Code

filename | sha512 hash
-------- | -----------
[kubernetes.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes.tar.gz) | de3fb80c8fdcabe60f37e3dcb1c61e8733c95fc0d45840f6861eafde09a149c3880f3e0b434d33167ffa66bdfeb887696ac7bfd2b44b85c29f99ba12965305ed
[kubernetes-src.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-src.tar.gz) | 753b9022b3c487d4bc9f8b302de14b7b4ef52b7664ff6d6b8bca65b6896cbc5932038de551a02c412afdd3ac2d56a8141e0dcb1dac7d24102217bd4f2beff936

### Client binaries

filename | sha512 hash
-------- | -----------
[kubernetes-client-darwin-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-client-darwin-amd64.tar.gz) | 8ba8627419704285abad0d98d28555d4bf4ce624c6958d0cca5ca8f53f1c40bb514631980ef39d52e2a604aff93bc078b30256d307d8af9839df91f8493d9aa5
[kubernetes-client-darwin-arm64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-client-darwin-arm64.tar.gz) | a039181d9dbff3203e75f357c65eaaf1667ab0834167b9ac12ff76999e276b9cc077e843b6043388183bd7c350c42ea28ab2d7b074c4f1987e43298e918595e1
[kubernetes-client-linux-386.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-client-linux-386.tar.gz) | 3474282cfe7f8f2966fca742453c632294ba224126748b162d42bd68a715681f2845c740252400d0b7d21dd3a11440530a5b84e454225655c16e056ca413e9de
[kubernetes-client-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-client-linux-amd64.tar.gz) | 7bd1e8b21af6b72757cdef9a4d76ea0eda3dbd558f2f5a7bee8f24f2c9b05d1cf52cfebd2f5ea991811917c3c18f1ac3dbde7e5094d5cd8a73478077a797b801
[kubernetes-client-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-client-linux-arm.tar.gz) | 0505f0c8e3733584ad1fc22ad729aea9f2452c8452ab1ed5e735e53ff48a92c248ba7310e5e9fa76630fa06a600c4ce8ee1b2b2845f07dba795fddbff5b7e941
[kubernetes-client-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-client-linux-arm64.tar.gz) | f5cbb08845bc6519538325250a7826e65ede254e5cf700a3f9b9128fec205f8d90827639bc64146b7c44008acd6a708bba59a3fbcefec1ca8e0050f6e3330290
[kubernetes-client-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-client-linux-ppc64le.tar.gz) | a3d90dc2ca5970ef4029ad9e9ff678816048c4dc58e7ad0f17a9a873855d71fdb3d23f4f7c88465f2261ed72747e85b78c80006e221e456bab0f07dc91022f1c
[kubernetes-client-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-client-linux-s390x.tar.gz) | cfee985e127f9471da4cb538362e3150c4edf12e8c72c5415024244007c9bf46c8f4a7f19e9fa8afb3126e379efce837114f8d1cee0f78d1602fe5e807e24b06
[kubernetes-client-windows-386.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-client-windows-386.tar.gz) | 47811776c0d1569afb3c8a689bb8989b57e8d3da4291606da6fc8b481e79b8632ac333f5c011e2bfd4fe4677827b27f64bd15253c2d83fdb5c0ce40671322e82
[kubernetes-client-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-client-windows-amd64.tar.gz) | d009d8178f94bcd69a1ae5a6ff39438b9811204f4c4f3b11b6219bcbd7d80f86ed2d6486feb88128fa42383550e37af6b3a603f0cecae1fdb86b69725d0b331a

### Server binaries

filename | sha512 hash
-------- | -----------
[kubernetes-server-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-server-linux-amd64.tar.gz) | 9bec26661b3ca7a688da8cc6fbb6ba4bf5e9993599401dbc9f7d95a2805d8a5c319052c30f33236094ba0a3b984a2246173d5334457ce7453ce74c84f5012c01
[kubernetes-server-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-server-linux-arm.tar.gz) | 89737d178779c9c636c246995aca9447a8e22150c63ae57cc3f1360b905c654d0f1c47dd35f958262e26a5fe61212fad308778d2acc9dbd8baff563f4c9a3e48
[kubernetes-server-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-server-linux-arm64.tar.gz) | 9ddb37baa8d2589eb2f3611cea8df71be26f9f2e4d935d552a530e9c5815f20d20aec6069a476b77fb2b99b2701289def2565b27c772713fee4b0fde8b804b95
[kubernetes-server-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-server-linux-ppc64le.tar.gz) | 8db94c576b6845b52ec16fb009a158ef2d733733c8fca48b2fadaef085b371d24b5e5f68758df24ec72189ea7963a9c72cff82b6d6163d1e89ef73de7fd830bd
[kubernetes-server-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-server-linux-s390x.tar.gz) | 99e086b5b2e39fcc6610232493cf56548913fb5bde9323cf301834b707518e20a6ce5c6d4713f9cd304cc4b9190de077e6d935e359396fabba1c436e658cc8bc

### Node binaries

filename | sha512 hash
-------- | -----------
[kubernetes-node-linux-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-node-linux-amd64.tar.gz) | 45bed8e46bd18ff86346fe4c3a971411d973b69e5cfd0db58162972bdc37fdf3387642284e43b9436e3862d8f2ee51ad8b147ee13a260b8fc9f42cbca78a1209
[kubernetes-node-linux-arm.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-node-linux-arm.tar.gz) | 3bf9e33cf90cd87679027b63989f3110e486b101189a8f0f05d0d8bdb5d22479ab4f84697413219d54e3c503ad54c533ee985144a57b45f093899e926e5b37fd
[kubernetes-node-linux-arm64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-node-linux-arm64.tar.gz) | ae1c5f1a0b40585a42e62075f173cfa9c6bcf81ad16fb9f04bf16e5df9bb02f5526cbdd93fbf1a811cba2001598fd04a53fad731bf4b917d498f60c93124a526
[kubernetes-node-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-node-linux-ppc64le.tar.gz) | 3dc8197d953dfd873ecd5e7a2b04d5b8b82d972b774497873f935b2e3ba033f05317866b3b795df56bb06f80e34545f100a89af9083d4ad6e9334295bb5262db
[kubernetes-node-linux-s390x.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-node-linux-s390x.tar.gz) | ec8f013c3e1a6bb151c968461b3f6b03b2a08283f4d253ec52e83acda2c03ac73fbae1de771baf69dfa26eb3a92f894fd2486ca8323f3d4750640b5b38bd99c4
[kubernetes-node-windows-amd64.tar.gz](https://dl.k8s.io/v1.22.0-alpha.1/kubernetes-node-windows-amd64.tar.gz) | acc8e3352a8d8ed8640d0787f2fb0d51ab0dac6f84687ab00a05c4a5470f1eb4821c878004e16a829cfd134d38e6f63b4b7f165637085d82a0a638f37e3c081e

## Changelog since v1.21.0

## Urgent Upgrade Notes

### (No, really, you MUST read this before you upgrade)

 - Audit log files are now created with a mode of 0600. Existing file permissions will not be changed. If you need the audit file to be readable by a non-root user, you can pre-create the file with the desired permissions. ([#95387](https://github.com/kubernetes/kubernetes/pull/95387), [@JAORMX](https://github.com/JAORMX)) [SIG API Machinery and Auth]
 
## Changes by Kind

### Deprecation

- Kubeadm: remove the deprecated kubeadm API v1beta1. Introduce a new kubeadm API v1beta3. See https://pkg.go.dev/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta3 for a list of changes since v1beta2. Note that v1beta2 is not yet deprecated, but will be in a future release. ([#101129](https://github.com/kubernetes/kubernetes/pull/101129), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- PodUnknown phase is now deprecated. ([#95286](https://github.com/kubernetes/kubernetes/pull/95286), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev)) [SIG Apps, CLI, Network, Node, Storage and Testing]
- Removal of the CSI nodepublish path by the kubelet is deprecated. This must be done by the CSI plugin according to the CSI spec. ([#101441](https://github.com/kubernetes/kubernetes/pull/101441), [@dobsonj](https://github.com/dobsonj)) [SIG Storage]

### API Change

- "Auto" is now a valid value for the `service.kubernetes.io/topology-aware-hints` annotation. ([#100728](https://github.com/kubernetes/kubernetes/pull/100728), [@robscott](https://github.com/robscott)) [SIG Apps, Instrumentation and Network]
- Kube-apiserver: `--service-account-issuer` can be specified multiple times now, to enable non-disruptive change of issuer. ([#101155](https://github.com/kubernetes/kubernetes/pull/101155), [@zshihang](https://github.com/zshihang)) [SIG API Machinery, Auth, Node and Testing]
- New "node-high" priority-level has been added to Suggested API Priority and Fairness configuration. ([#101151](https://github.com/kubernetes/kubernetes/pull/101151), [@mborsz](https://github.com/mborsz)) [SIG API Machinery]
- PodDeletionCost promoted to Beta ([#101080](https://github.com/kubernetes/kubernetes/pull/101080), [@ahg-g](https://github.com/ahg-g)) [SIG Apps]
- SSA treats certain structs as atomic ([#100684](https://github.com/kubernetes/kubernetes/pull/100684), [@Jefftree](https://github.com/Jefftree)) [SIG API Machinery, Auth, Node and Storage]
- Server Side Apply now treats all <Some>Selector fields as atomic (meaning the entire selector is managed by a single writer and updated together), since they contain interrelated and inseparable fields that do not merge in intuitive ways. ([#97989](https://github.com/kubernetes/kubernetes/pull/97989), [@Danil-Grigorev](https://github.com/Danil-Grigorev)) [SIG API Machinery]
- The `pods/ephemeralcontainers` API now returns and expects a `Pod` object instead of `EphemeralContainers`. This is incompatible with the previous alpha-level API. ([#101034](https://github.com/kubernetes/kubernetes/pull/101034), [@verb](https://github.com/verb)) [SIG Apps, Auth, CLI and Testing]
- The pod/eviction subresource now accepts policy/v1 Eviction requests in addition to policy/v1beta1 Eviction requests ([#100724](https://github.com/kubernetes/kubernetes/pull/100724), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Apps, Architecture, Auth, CLI, Storage and Testing]
- Track ownership of scale subresource for all scalable resources i.e. Deployment, ReplicaSet, StatefulSet, ReplicationController, and Custom Resources. ([#98377](https://github.com/kubernetes/kubernetes/pull/98377), [@nodo](https://github.com/nodo)) [SIG API Machinery and Testing]
- We have added a new Priority & Fairness rule that exempts all probes (/readyz, /healthz, /livez) to prevent 
  restarting of "healthy" kube-apiserver instance(s) by kubelet. ([#100678](https://github.com/kubernetes/kubernetes/pull/100678), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]

### Feature

- Base image updates to mitigate kube-proxy and etcd container image CVEs
  - debian-base to buster-v1.6.0
  - debian-iptables to buster-v1.6.0 ([#100976](https://github.com/kubernetes/kubernetes/pull/100976), [@jindijamie](https://github.com/jindijamie)) [SIG Release and Testing]
- EmptyDir memory backed volumes are sized as the the minimum of pod allocatable memory on a host and an optional explicit user provided value. ([#101048](https://github.com/kubernetes/kubernetes/pull/101048), [@dims](https://github.com/dims)) [SIG Node]
- Fluentd: isolate logging resources in separate namespace ([#68004](https://github.com/kubernetes/kubernetes/pull/68004), [@saravanan30erd](https://github.com/saravanan30erd)) [SIG Cloud Provider and Instrumentation]
- It add two flags, `--max-pods` and `--extended-resources` ([#100267](https://github.com/kubernetes/kubernetes/pull/100267), [@Jeffwan](https://github.com/Jeffwan)) [SIG Node and Scalability]
- Kube config is now exposed in the scheduler framework handle. Out-of-tree plugins can leverage that to build CRD informers easily. ([#100644](https://github.com/kubernetes/kubernetes/pull/100644), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Apps, Scheduling and Testing]
- Kubeadm: add `--validity-period` flag for 'kubeadm kubeconfig user' command ([#100907](https://github.com/kubernetes/kubernetes/pull/100907), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Kubemark's hollow-node will now print flags before starting ([#101181](https://github.com/kubernetes/kubernetes/pull/101181), [@mm4tt](https://github.com/mm4tt)) [SIG Scalability]
- Kubernetes is now built with Golang 1.16.3 ([#101206](https://github.com/kubernetes/kubernetes/pull/101206), [@justaugustus](https://github.com/justaugustus)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Promote NamespaceDefaultLabelName to GA.  All Namespace API objects have a `kubernetes.io/metadata.name` label matching their metadata.name field to allow selecting any namespace by its name using a label selector. ([#101342](https://github.com/kubernetes/kubernetes/pull/101342), [@rosenhouse](https://github.com/rosenhouse)) [SIG API Machinery and Apps]
- Run etcd as non-root on GCE provider' ([#100635](https://github.com/kubernetes/kubernetes/pull/100635), [@cindy52](https://github.com/cindy52)) [SIG Cloud Provider]
- SSA is GA ([#100139](https://github.com/kubernetes/kubernetes/pull/100139), [@Jefftree](https://github.com/Jefftree)) [SIG API Machinery]
- System-cluster-critical pods should not get a low OOM Score. 
  
  As of now both system-node-critical and system-cluster-critical pods have -997 OOM score, making them one of the last processes to be OOMKilled. By definition system-cluster-critical pods can be scheduled elsewhere if there is a resource crunch on the node where as system-node-critical pods cannot be rescheduled. This was the reason for system-node-critical to have higher priority value than system-cluster-critical.  This change allows only system-node-critical priority class to have low OOMScore.
  
  action required
  If the user wants to have the pod to be OOMKilled last and the pod has system-cluster-critical priority class, it has to be changed to system-node-critical priority class to preserve the existing behavior ([#99729](https://github.com/kubernetes/kubernetes/pull/99729), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Node]
- The job controller removes running pods when the number of completions was achieved. ([#99963](https://github.com/kubernetes/kubernetes/pull/99963), [@alculquicondor](https://github.com/alculquicondor)) [SIG Apps]
- `kubectl describe` will by default fetch large lists of resources in chunks of up to 500 items rather than requesting all resources up front from the server. A new flag `--chunk-size=SIZE` may be used to alter the number of items or disable this feature when `0` is passed.  This is a beta feature. ([#101171](https://github.com/kubernetes/kubernetes/pull/101171), [@KnVerey](https://github.com/KnVerey)) [SIG CLI and Testing]
- `kubectl drain` will by default fetch large lists of resources in chunks of up to 500 items rather than requesting all resources up front from the server. A new flag `--chunk-size=SIZE` may be used to alter the number of items or disable this feature when `0` is passed.  This is a beta feature. ([#100148](https://github.com/kubernetes/kubernetes/pull/100148), [@KnVerey](https://github.com/KnVerey)) [SIG CLI and Testing]

### Failing Test

- Fixed generic ephemeal volumes with OwnerReferencesPermissionEnforcement admission plugin enabled. ([#101186](https://github.com/kubernetes/kubernetes/pull/101186), [@jsafrane](https://github.com/jsafrane)) [SIG Auth and Storage]
- Fixes kubectl drain --dry-run=server ([#100206](https://github.com/kubernetes/kubernetes/pull/100206), [@KnVerey](https://github.com/KnVerey)) [SIG CLI and Testing]

### Bug or Regression

- Added privileges for EndpointSlice to the default view & edit RBAC roles ([#101203](https://github.com/kubernetes/kubernetes/pull/101203), [@mtougeron](https://github.com/mtougeron)) [SIG Auth and Security]
- Chain the field manager creation calls in newDefaultFieldManager ([#101076](https://github.com/kubernetes/kubernetes/pull/101076), [@kevindelgado](https://github.com/kevindelgado)) [SIG API Machinery]
- EndpointSlice IP validation now matches Endpoints IP validation. ([#101084](https://github.com/kubernetes/kubernetes/pull/101084), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Ensure service deleted when the Azure resource group has been deleted ([#100944](https://github.com/kubernetes/kubernetes/pull/100944), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Expose `rest_client_rate_limiter_duration_seconds` metric to component-base to track client side rate limiter latency in seconds. Broken down by verb and URL. ([#100311](https://github.com/kubernetes/kubernetes/pull/100311), [@IonutBajescu](https://github.com/IonutBajescu)) [SIG API Machinery, Cluster Lifecycle and Instrumentation]
- Fire an event when failing to open NodePort ([#100599](https://github.com/kubernetes/kubernetes/pull/100599), [@masap](https://github.com/masap)) [SIG Network]
- Fix a bug in kube-proxy latency metrics to calculate only the latency value for the endpoints that are created after it starts running. This is needed because all the endpoints objects are processed on restarts, independently when they were generated. ([#100861](https://github.com/kubernetes/kubernetes/pull/100861), [@aojea](https://github.com/aojea)) [SIG Instrumentation and Network]
- Fix availability set cache in vmss cache ([#100110](https://github.com/kubernetes/kubernetes/pull/100110), [@CecileRobertMichon](https://github.com/CecileRobertMichon)) [SIG Cloud Provider]
- Fix display of Job completion mode in kubectl describe ([#101160](https://github.com/kubernetes/kubernetes/pull/101160), [@alculquicondor](https://github.com/alculquicondor)) [SIG Apps and CLI]
- Fix panic with kubectl create ingress annotation flag and empty value ([#101377](https://github.com/kubernetes/kubernetes/pull/101377), [@rikatz](https://github.com/rikatz)) [SIG CLI]
- Fix raw block mode CSI NodePublishVolume stage miss pod info ([#99069](https://github.com/kubernetes/kubernetes/pull/99069), [@phantooom](https://github.com/phantooom)) [SIG Storage]
- Fix rounding of volume storage requests ([#100100](https://github.com/kubernetes/kubernetes/pull/100100), [@maxlaverse](https://github.com/maxlaverse)) [SIG Cloud Provider and Storage]
- Fix: azure file inline volume namespace issue in csi migration translation ([#101235](https://github.com/kubernetes/kubernetes/pull/101235), [@andyzhangx](https://github.com/andyzhangx)) [SIG Apps, Cloud Provider, Node and Storage]
- Fix: not delete existing pip when service is deleted ([#100694](https://github.com/kubernetes/kubernetes/pull/100694), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: set "host is down" as corrupted mount ([#101398](https://github.com/kubernetes/kubernetes/pull/101398), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed a bug where startupProbe stopped working after a container's first restart ([#101093](https://github.com/kubernetes/kubernetes/pull/101093), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- Fixed port-forward memory leak for long-running and heavily used connections. ([#99839](https://github.com/kubernetes/kubernetes/pull/99839), [@saschagrunert](https://github.com/saschagrunert)) [SIG API Machinery and Node]
- Fixed using volume partitions on AWS Nitro systems. ([#100500](https://github.com/kubernetes/kubernetes/pull/100500), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Generated OpenAPI now correctly specifies 201 as a possible response code for PATCH operations ([#100141](https://github.com/kubernetes/kubernetes/pull/100141), [@brendandburns](https://github.com/brendandburns)) [SIG API Machinery]
- KCM sets the upper-bound timeout limit for outgoing requests to 70s. Previously no timeout was set. Requests without explicit timeout might potentially hang forever and lead to starvation of the application. ([#99358](https://github.com/kubernetes/kubernetes/pull/99358), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery]
- Kubeadm: enable `--experimental-patches` flag for 'kubeadm join phase control-plane-join all' command ([#101110](https://github.com/kubernetes/kubernetes/pull/101110), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Kubelet: improve the performance when waiting for a synchronization of the node list with the kube-apiserver ([#99336](https://github.com/kubernetes/kubernetes/pull/99336), [@neolit123](https://github.com/neolit123)) [SIG Node]
- Logging for GCE Windows clusters will be more accurate and complete when using Fluent-bit. ([#101271](https://github.com/kubernetes/kubernetes/pull/101271), [@jeremyje](https://github.com/jeremyje)) [SIG Cloud Provider and Windows]
- No support endpointslice in linux userpace mode ([#100913](https://github.com/kubernetes/kubernetes/pull/100913), [@JornShen](https://github.com/JornShen)) [SIG Network]
- Prevent Kubelet stuck in DiskPressure when imagefs minReclaim is set ([#99095](https://github.com/kubernetes/kubernetes/pull/99095), [@maxlaverse](https://github.com/maxlaverse)) [SIG Node]
- Reduce vSphere volume name to 63 characters ([#100404](https://github.com/kubernetes/kubernetes/pull/100404), [@gnufied](https://github.com/gnufied)) [SIG Storage]
- Reduces delay initializing on non-AWS platforms docker runtime. ([#93260](https://github.com/kubernetes/kubernetes/pull/93260), [@nckturner](https://github.com/nckturner)) [SIG Cloud Provider]
- Removed `/sbin/apparmor_parser` requirement for the AppArmor host validation.
  This allows using AppArmor on distributions which ship the binary in a different path. ([#97968](https://github.com/kubernetes/kubernetes/pull/97968), [@saschagrunert](https://github.com/saschagrunert)) [SIG Node and Testing]
- Renames the timeout field for the DelegatingAuthenticationOptions to TokenRequestTimeout and set the timeout only for the token review client. Previously the timeout was also applied to watches making them reconnecting every 10 seconds. ([#100959](https://github.com/kubernetes/kubernetes/pull/100959), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery, Auth and Cloud Provider]
- Reorganized iptables rules to reduce rules in KUBE-SERVICES and KUBE-NODEPORTS chains and improve performance ([#96959](https://github.com/kubernetes/kubernetes/pull/96959), [@tssurya](https://github.com/tssurya)) [SIG Network]
- Respect ExecProbeTimeout=false for dockershim ([#100200](https://github.com/kubernetes/kubernetes/pull/100200), [@jackfrancis](https://github.com/jackfrancis)) [SIG Node and Testing]
- Restore kind-specific output for `kubectl describe podsecuritypolicy` ([#101436](https://github.com/kubernetes/kubernetes/pull/101436), [@KnVerey](https://github.com/KnVerey)) [SIG CLI]
- The kubelet now reports distinguishes log messages about certificate rotation for its client cert and server cert separately to make debugging problems with one or the other easier. ([#101252](https://github.com/kubernetes/kubernetes/pull/101252), [@smarterclayton](https://github.com/smarterclayton)) [SIG API Machinery and Auth]
- Updates dependency sigs.k8s.io/structured-merge-diff to v4.1.1 ([#100784](https://github.com/kubernetes/kubernetes/pull/100784), [@kevindelgado](https://github.com/kevindelgado)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Storage]
- Upgrades functionality of `kubectl kustomize` as described at
  https://github.com/kubernetes-sigs/kustomize/releases/tag/kustomize%2Fv4.1.2 ([#101120](https://github.com/kubernetes/kubernetes/pull/101120), [@monopole](https://github.com/monopole)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle and Instrumentation]
- Use default timeout of 10s for Azure ACR credential provider. ([#100686](https://github.com/kubernetes/kubernetes/pull/100686), [@hasheddan](https://github.com/hasheddan)) [SIG Cloud Provider]
- [kubeadm] Support for custom imagetags for etcd images which contain build metadata, when imagetags are in the form of version_metadata. For instance, if the etcd version is v3.4.13+patch.0, the supported imagetag would be v3.4.13_patch.0 ([#100350](https://github.com/kubernetes/kubernetes/pull/100350), [@jr0d](https://github.com/jr0d)) [SIG Cluster Lifecycle]

### Other (Cleanup or Flake)

- After the deprecation period,now the Kubelet's  `--chaos-chance` flag  are removed. ([#101057](https://github.com/kubernetes/kubernetes/pull/101057), [@wangyysde](https://github.com/wangyysde)) [SIG Node]
- DynamicFakeClient now exposes its tracker via a `Tracker()` function ([#100085](https://github.com/kubernetes/kubernetes/pull/100085), [@markusthoemmes](https://github.com/markusthoemmes)) [SIG API Machinery]
- Exposes WithCustomRoundTripper method for specifying a middleware function for custom HTTP behaviour for the delegated auth clients. ([#99775](https://github.com/kubernetes/kubernetes/pull/99775), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery]
- Migrate some log messages to structured logging in pkg/volume/volume_linux.go. ([#99566](https://github.com/kubernetes/kubernetes/pull/99566), [@huchengze](https://github.com/huchengze)) [SIG Instrumentation and Storage]
- Official binaries now include the golang generated build ID (`buildid`) instead of an empty string. ([#101411](https://github.com/kubernetes/kubernetes/pull/101411), [@saschagrunert](https://github.com/saschagrunert)) [SIG Release]
- Remove deprecated --generator flag from kubectl autoscale ([#99900](https://github.com/kubernetes/kubernetes/pull/99900), [@MadhavJivrajani](https://github.com/MadhavJivrajani)) [SIG CLI]
- Remove the deprecated flag --generator from kubectl create deployment command ([#99915](https://github.com/kubernetes/kubernetes/pull/99915), [@BLasan](https://github.com/BLasan)) [SIG CLI]
- Update Azure Go SDK version to v53.1.0 ([#101357](https://github.com/kubernetes/kubernetes/pull/101357), [@feiskyer](https://github.com/feiskyer)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle and Instrumentation]
- Update cri-tools dependency to v1.21.0 ([#100956](https://github.com/kubernetes/kubernetes/pull/100956), [@saschagrunert](https://github.com/saschagrunert)) [SIG Cloud Provider and Node]

## Dependencies

### Added
- github.com/gofrs/uuid: [v4.0.0+incompatible](https://github.com/gofrs/uuid/tree/v4.0.0)
- github.com/stoewer/go-strcase: [v1.2.0](https://github.com/stoewer/go-strcase/tree/v1.2.0)
- go.uber.org/tools: 2cfd321

### Changed
- github.com/Azure/azure-sdk-for-go: [v43.0.0+incompatible → v53.1.0+incompatible](https://github.com/Azure/azure-sdk-for-go/compare/v43.0.0...v53.1.0)
- github.com/Azure/go-autorest/autorest/adal: [v0.9.5 → v0.9.10](https://github.com/Azure/go-autorest/autorest/adal/compare/v0.9.5...v0.9.10)
- github.com/Azure/go-autorest/autorest: [v0.11.12 → v0.11.17](https://github.com/Azure/go-autorest/autorest/compare/v0.11.12...v0.11.17)
- github.com/googleapis/gnostic: [v0.4.1 → v0.5.1](https://github.com/googleapis/gnostic/compare/v0.4.1...v0.5.1)
- go.uber.org/atomic: v1.4.0 → v1.6.0
- go.uber.org/multierr: v1.1.0 → v1.5.0
- go.uber.org/zap: v1.10.0 → v1.16.0
- gopkg.in/yaml.v3: 9f266ea → eeeca48
- k8s.io/kube-openapi: 591a79e → 9528897
- sigs.k8s.io/kustomize/api: v0.8.5 → v0.8.8
- sigs.k8s.io/kustomize/cmd/config: v0.9.7 → v0.9.10
- sigs.k8s.io/kustomize/kustomize/v4: v4.0.5 → v4.1.2
- sigs.k8s.io/kustomize/kyaml: v0.10.15 → v0.10.17
- sigs.k8s.io/structured-merge-diff/v4: v4.1.0 → v4.1.1

### Removed
- github.com/satori/go.uuid: [v1.2.0](https://github.com/satori/go.uuid/tree/v1.2.0)