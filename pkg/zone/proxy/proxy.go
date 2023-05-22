package proxy

import (
	"context"
	"fmt"
	"sync"
	"time"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/flowcontrol"
	"k8s.io/klog"
	api "k8s.io/kubernetes/pkg/apis/core"

	"github.com/RachaelLuo/kex/pkg/apis/cluster"
	"github.com/RachaelLuo/kex/pkg/utils"
	"github.com/RachaelLuo/kex/pkg/zone/clientset"
)

var (
	codeClusterClient sync.Map
	idClusterClient   sync.Map
	clusterWatch      watch.Interface
	retrySync         = make(chan struct{}, 1)
)

func InitProxy(clusterInfo, namespace, localClusterInfos string, kubeclient clientset.Interface) (err error) {
	if clusterWatch, err = kubeclient.CoreV1().ConfigMaps(namespace).Watch(context.TODO(), metav1.SingleObject(metav1.ObjectMeta{
		Name:      clusterInfo,
		Namespace: namespace,
	})); err != nil {
		return
	}

	go watchcm(clusterInfo, namespace, localClusterInfos, kubeclient)
	go watchConfig(clusterInfo, namespace, localClusterInfos, kubeclient)

	return
}

func watchcm(clusterInfo, namespace, localClusterInfos string, kubeclient clientset.Interface) {
	for {
		if _, err := kubeclient.CoreV1().ConfigMaps(namespace).Get(context.TODO(), clusterInfo, metav1.GetOptions{}); err != nil {
			if apierrors.IsNotFound(err) {
				if _, err = utils.CreateConfigMapsFromLocal(localClusterInfos, namespace, clusterInfo, kubeclient); err != nil {
					klog.Error(err)
				}
			} else {
				klog.Error(err)
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func watchConfig(clusterInfo, namespace, localClusterInfos string, kubeclient clientset.Interface) {
	retryNew := make(chan struct{}, 1)
	for {
		select {
		case _, ok := <-clusterWatch.ResultChan():
			if !ok {
				if err := newClustersAndWatch(clusterInfo, namespace, localClusterInfos, kubeclient); err != nil {
					klog.Error(err)
					retryNew <- struct{}{}
				}
			} else {
				if err := syncCodeClusterClient(clusterInfo, namespace, kubeclient); err != nil {
					klog.Error(err)
					retrySync <- struct{}{}
				}
			}
		case <-retryNew:
			if err := newClustersAndWatch(clusterInfo, namespace, localClusterInfos, kubeclient); err != nil {
				klog.Error(err)
				retryNew <- struct{}{}
			}
		case <-retrySync:
			if err := syncCodeClusterClient(clusterInfo, namespace, kubeclient); err != nil {
				klog.Error(err)
				retrySync <- struct{}{}
			}
		}
	}
}

func newClustersAndWatch(clusterInfo, namespace, localClusterInfos string, kubeclient clientset.Interface) error {
	var err error
	if _, err := kubeclient.CoreV1().ConfigMaps(namespace).Get(context.TODO(), clusterInfo, metav1.GetOptions{}); err != nil {
		if apierrors.IsNotFound(err) {
			if _, err = utils.CreateConfigMapsFromLocal(localClusterInfos, namespace, clusterInfo, kubeclient); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	clusterWatch, err = kubeclient.CoreV1().ConfigMaps(namespace).Watch(context.TODO(), metav1.SingleObject(metav1.ObjectMeta{
		Name:      clusterInfo,
		Namespace: namespace,
	}))

	return err
}

func syncCodeClusterClient(clusterInfo, namespace string, kubeclient clientset.Interface) error {
	cm, err := kubeclient.CoreV1().ConfigMaps(namespace).Get(context.TODO(), clusterInfo, metav1.GetOptions{})
	if err != nil {
		return err
	}
	idtmp := make(map[string]*clientset.Clientset, len(cm.BinaryData))
	codetmp := make(map[string]*clientset.Clientset, len(cm.BinaryData))
	codes := make([]string, 0, len(cm.BinaryData))
	for code, data := range cm.BinaryData {
		client, err := newClient(data)
		if err != nil {
			return err
		}

		ns, err := client.CoreV1().Namespaces().Get(context.TODO(), api.NamespaceSystem, metav1.GetOptions{})
		if err != nil {
			return err
		}

		idtmp[string(ns.GetUID())] = client
		codetmp[code] = client
		codes = append(codes, code)
	}

	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	codeClusterClient.Range(func(key, _ any) bool {
		codeClusterClient.Delete(key)
		return true
	})

	idClusterClient.Range(func(key, _ any) bool {
		idClusterClient.Delete(key)
		return true
	})

	for k, v := range codetmp {
		codeClusterClient.Store(k, v)
	}

	for k, v := range idtmp {
		idClusterClient.Store(k, v)
	}
	klog.Infof("cluster %v proxy successfull", codes)
	return nil
}

func newClient(data []byte) (*clientset.Clientset, error) {
	clusterInfo := &cluster.ClusterInfo{}
	if err := utils.Std2Jsoniter.Unmarshal(data, clusterInfo); err != nil {
		return nil, err
	}

	restConfig, err := clientcmd.RESTConfigFromKubeConfig(clusterInfo.Kubeconfig)
	if err != nil {
		return nil, err
	}

	// set rateLimiter 1000
	restConfig.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(1000, 1000)
	return clientset.NewForConfig(restConfig)
}

func GetClusterPorxyClientFromCode(code string) (*clientset.Clientset, error) {
	client, ok := codeClusterClient.Load(code)
	if !ok {
		return nil, fmt.Errorf("cluster %v Not Found", code)
	}
	return client.(*clientset.Clientset), nil
}

func GetClusterPorxyClientFromID(id string) (*clientset.Clientset, error) {
	client, ok := idClusterClient.Load(id)
	if !ok {
		return nil, fmt.Errorf("cluster %v Not Found", id)
	}
	return client.(*clientset.Clientset), nil
}
