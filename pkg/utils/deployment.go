package utils

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ListDeploymentPods returns a set of pods controlled by given deployment.
func ListDeploymentPods(oldCtx context.Context, client kubernetes.Interface, deployment appsv1.Deployment) ([]corev1.Pod, error) {
	selector, err := metav1.LabelSelectorAsSelector(deployment.Spec.Selector)
	if err != nil {
		return nil, err
	}
	allRS, err := client.AppsV1().ReplicaSets(deployment.Namespace).List(oldCtx, metav1.ListOptions{
		LabelSelector: selector.String(),
	})
	if err != nil {
		return nil, err
	}
	allPods, err := client.CoreV1().Pods(deployment.Namespace).List(oldCtx, metav1.ListOptions{
		LabelSelector: selector.String(),
	})
	if err != nil {
		return nil, err
	}
	return FilterDeploymentPodsByOwnerReference(deployment, allRS.Items, allPods.Items), nil
}

// FilterDeploymentPodsByOwnerReference returns a subset of pods controlled by given deployment.
func FilterDeploymentPodsByOwnerReference(deployment appsv1.Deployment, allRS []appsv1.ReplicaSet,
	allPods []corev1.Pod) []corev1.Pod {
	matchingPods := make([]corev1.Pod, 0)
	for rsk := range allRS {
		if len(allPods) == 0 {
			break
		}
		if metav1.IsControlledBy(&allRS[rsk], &deployment) {
			matchingPods = append(matchingPods, FilterPodsByControllerRef(&allRS[rsk], &allPods)...)
		}
	}
	return matchingPods
}
