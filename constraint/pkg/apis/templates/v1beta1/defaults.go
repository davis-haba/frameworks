package v1beta1

import (
	"github.com/open-policy-agent/frameworks/constraint/pkg/apis/templates"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/schema/defaulting"
	"k8s.io/apimachinery/pkg/runtime"
)

const version = "v1beta1"

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_ConstraintTemplate(obj *ConstraintTemplate) { // nolint:revive // Required exact function name.
	// turn the CT into an unstructured
	un, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		panic("Failed to convert v1 ConstraintTemplate to Unstructured")
	}

	defaulting.Default(un, templates.ConstraintTemplateSchemas[version])

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(un, obj)
	if err != nil {
		panic("Failed to convert Unstructured to v1 ConstraintTemplate")
	}
}