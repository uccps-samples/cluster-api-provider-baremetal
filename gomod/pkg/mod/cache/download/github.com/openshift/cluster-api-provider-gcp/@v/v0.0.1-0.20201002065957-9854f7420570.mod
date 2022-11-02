module github.com/openshift/cluster-api-provider-gcp

go 1.13

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/go-logr/logr v0.2.1
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/uccps-samples/api v0.0.0-20200901182017-7ac89ba6b971
	github.com/uccps-samples/machine-api-operator v0.2.1-0.20200926044412-b7d860f8074c
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6
	google.golang.org/api v0.15.0

	// kube 1.18
	k8s.io/api v0.19.0
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
	k8s.io/klog v1.0.0
	k8s.io/klog/v2 v2.3.0
	sigs.k8s.io/controller-runtime v0.6.2
	sigs.k8s.io/controller-tools v0.3.0
	sigs.k8s.io/yaml v1.2.0
)

replace (
	sigs.k8s.io/cluster-api-provider-aws => github.com/uccps-samples/cluster-api-provider-aws v0.2.1-0.20200618031251-e16dd65fdd85
	sigs.k8s.io/cluster-api-provider-azure => github.com/uccps-samples/cluster-api-provider-azure v0.1.0-alpha.3.0.20200618001858-af08a66b92de
)