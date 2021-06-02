package admission

import (
	"github.com/anchore/kubernetes-admission-controller/cmd/kubernetes-admission-controller/anchore"
	"github.com/anchore/kubernetes-admission-controller/cmd/kubernetes-admission-controller/validation"
)

type PolicySelector struct {
	ResourceSelector ResourceSelector
	PolicyReference  anchore.ClientConfiguration
	Mode             validation.Mode
}
