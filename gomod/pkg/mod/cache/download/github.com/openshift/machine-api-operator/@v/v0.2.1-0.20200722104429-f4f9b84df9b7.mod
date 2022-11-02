module github.com/uccps-samples/machine-api-operator

go 1.13

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/evanphx/json-patch v4.5.0+incompatible
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/google/gofuzz v1.1.0
	github.com/google/uuid v1.1.1
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.8.1
	github.com/uccps-samples/api v0.0.0-20200424083944-0422dc17083e
	github.com/uccps-samples/client-go v0.0.0-20200326155132-2a6cd50aedd0
	github.com/openshift/cluster-api-provider-gcp v0.0.1-0.20200713133651-5c8a640669ac
	github.com/uccps-samples/library-go v0.0.0-20200512120242-21a1ff978534
	github.com/operator-framework/operator-sdk v0.5.1-0.20190301204940-c2efe6f74e7b
	github.com/prometheus/client_golang v1.1.0
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.4.0
	github.com/vmware/govmomi v0.22.2
	golang.org/x/net v0.0.0-20200421231249-e086a090c8fd
	gopkg.in/gcfg.v1 v1.2.3
	gopkg.in/warnings.v0 v0.1.2 // indirect
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v0.18.2
	k8s.io/code-generator v0.18.2
	k8s.io/klog v1.0.0
	k8s.io/kubectl v0.18.0-rc.1
	k8s.io/utils v0.0.0-20200327001022-6496210b90e8
	sigs.k8s.io/cluster-api-provider-aws v0.0.0-00010101000000-000000000000
	sigs.k8s.io/cluster-api-provider-azure v0.0.0-00010101000000-000000000000
	sigs.k8s.io/controller-runtime v0.6.0
	sigs.k8s.io/controller-tools v0.3.0
	sigs.k8s.io/yaml v1.2.0
)

replace sigs.k8s.io/cluster-api-provider-aws => github.com/uccps-samples/cluster-api-provider-aws v0.2.1-0.20200713194048-b3f366f88cc5

replace sigs.k8s.io/cluster-api-provider-azure => github.com/uccps-samples/cluster-api-provider-azure v0.1.0-alpha.3.0.20200714223321-88a538beaf71
