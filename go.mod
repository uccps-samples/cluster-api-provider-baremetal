module github.com/openshift/cluster-api-provider-baremetal

go 1.16

require (
	github.com/metal3-io/baremetal-operator v0.0.0-00010101000000-000000000000
	github.com/openshift/machine-api-operator v0.2.1-0.20210513225032-5644b5803418
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	k8s.io/api v0.21.0
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/klog/v2 v2.8.0
	sigs.k8s.io/controller-runtime v0.9.0-beta.1.0.20210512131817-ce2f0c92d77e
	sigs.k8s.io/controller-tools v0.4.1
	sigs.k8s.io/yaml v1.2.0
)

replace (
	github.com/metal3-io/baremetal-operator => github.com/openshift/baremetal-operator v0.0.0-20220818132643-5492cf523099 // Use OpenShift fork
	k8s.io/client-go => k8s.io/client-go v0.21.0
	sigs.k8s.io/cluster-api-provider-aws => github.com/openshift/cluster-api-provider-aws v0.2.1-0.20220919090134-1a6124fe5dd8
	sigs.k8s.io/cluster-api-provider-azure => github.com/openshift/cluster-api-provider-azure v0.1.0-alpha.3.0.20220214133436-ad5852b1255c
)
