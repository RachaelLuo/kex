package utils

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FilterPodsByControllerRef returns a subset of pods controlled by given controller resource, excluding deployments.
func FilterPodsByControllerRef(owner metav1.Object, remainingPods *[]corev1.Pod) []corev1.Pod {
	matchingPods := make([]corev1.Pod, 0)
	for podInx := 0; podInx < len(*remainingPods); podInx++ {
		if metav1.IsControlledBy(&(*remainingPods)[podInx], owner) {
			matchingPods = append(matchingPods, (*remainingPods)[podInx])
			(*remainingPods) = append((*remainingPods)[:podInx], (*remainingPods)[podInx+1:]...)
			podInx--
		}
	}
	return matchingPods
}
