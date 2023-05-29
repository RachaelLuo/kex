package utils

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func Status(u *unstructured.Unstructured) (string, error) {
	return status(u)
}

func StatusFromRuntime(o runtime.Object) (string, error) {
	u, err := RuntimeObjectToUnstructured(o)
	if err != nil {
		return StatusUnknown, err
	}
	return status(u)
}

func DeploymentsEventStatusFromRuntime(o runtime.Object) (string, error) {
	objOut, err := runtime.DefaultUnstructuredConverter.ToUnstructured(o)
	if err != nil {
		return StatusUnknown, err
	}
	u := &unstructured.Unstructured{Object: objOut}
	if err != nil {
		return StatusUnknown, err
	}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "apps",
		Version: "v1",
		Kind:    "Deployment",
	})
	return status(u)
}
