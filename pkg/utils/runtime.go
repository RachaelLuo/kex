package utils

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func RuntimeObjectToUnstructured(o runtime.Object) (*unstructured.Unstructured, error) {
	objOut, err := runtime.DefaultUnstructuredConverter.ToUnstructured(o)
	if err != nil {
		return nil, err
	}
	u := &unstructured.Unstructured{Object: objOut}
	u.SetGroupVersionKind(o.GetObjectKind().GroupVersionKind())
	return u, nil
}
