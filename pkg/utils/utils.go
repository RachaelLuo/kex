package utils

import (
	"context"
	"os"

	jsoniter "github.com/json-iterator/go"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"

	"github.com/RachaelLuo/kex/pkg/apis/cluster"
	"github.com/RachaelLuo/kex/pkg/zone/clientset"
)

var Std2Jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary

func CreateConfigMapsFromLocal(localClusterInfos, namespace, clusterInfo string, client clientset.Interface) (*corev1.ConfigMap, error) {
	if len(localClusterInfos) == 0 {
		if cm, err := client.CoreV1().ConfigMaps(namespace).Create(context.TODO(), &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name: clusterInfo,
			},
		}, metav1.CreateOptions{}); err != nil {
			return nil, err
		} else {
			return cm, nil
		}
	}
	data, err := os.ReadFile(localClusterInfos)
	if err != nil {
		return nil, err
	}
	clusterInfos := make([]*cluster.ClusterInfo, 0)
	if err = Std2Jsoniter.Unmarshal(data, &clusterInfos); err != nil {
		return nil, err
	}

	if len(clusterInfos) == 0 {
		klog.Warning("no proxy cluster")
	}

	binaryData := make(map[string][]byte, len(clusterInfos))
	for i := range clusterInfos {
		if clusterData, err := Std2Jsoniter.Marshal(clusterInfos[i]); err != nil {
			return nil, err
		} else {
			binaryData[clusterInfos[i].Code] = clusterData
		}
	}

	if cm, err := client.CoreV1().ConfigMaps(namespace).Create(context.TODO(), &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: clusterInfo,
		},
		BinaryData: binaryData,
	}, metav1.CreateOptions{}); err != nil {
		return nil, err
	} else {
		return cm, nil
	}
}
