// +build tools

// Official workaround to track tool dependencies with go modules:
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package tools

import (
	// Need this for code generation
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"

	// Need the whole vendored repo for applying manifests
	_ "github.com/uccps-samples/machine-api-operator"
)
