package server

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"

	"github.com/RachaelLuo/kex/pkg/server/servererror"
	"github.com/RachaelLuo/kex/pkg/utils"
	"github.com/RachaelLuo/kex/pkg/zone/proxy"
)

// listPodOfDeployment get all pod on deployment
func (h *handler) listPodOfDeployment(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	client, err := proxy.GetClusterPorxyClientFromCode(c.Param("clusterCode"))
	if err != nil {
		servererror.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	deployment, err := client.AppsV1().Deployments(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		servererror.HandleError(c, http.StatusInternalServerError, err)
		return
	}
	resultList := corev1.PodList{Items: []corev1.Pod{}}
	if *deployment.Spec.Replicas != 0 {
		pods, err := utils.ListDeploymentPods(context.Background(), client, *deployment)
		if err != nil {
			servererror.HandleError(c, http.StatusInternalServerError, err)
			return
		}
		resultList = corev1.PodList{
			Items: pods,
		}
	}

	resultList.APIVersion = "v1"
	resultList.Kind = "List"
	c.JSON(http.StatusOK, resultList)
}

func (h *handler) restartDeployment(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	client, err := proxy.GetClusterPorxyClientFromCode(c.Param("clusterCode"))
	if err != nil {
		servererror.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	deployClient := client.AppsV1().Deployments(namespace)

	oldDeploy, err := deployClient.Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		servererror.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	restartTime := time.Now().UnixNano()
	oldDeploy.Spec.Template.Labels["kex.io/restart"] = strconv.FormatInt(restartTime, 10)

	deploy, err := deployClient.Update(context.Background(), oldDeploy, metav1.UpdateOptions{})
	if err != nil {
		servererror.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	if deploy.ResourceVersion == oldDeploy.ResourceVersion {
		servererror.HandleError(c, http.StatusInternalServerError, errors.New("restart deployment failed"))
		return
	}

	op := metav1.SingleObject(deploy.ObjectMeta)
	timeout := int64(5 * 60) // 5 min
	op.TimeoutSeconds = &timeout
	op.ResourceVersion = deploy.ResourceVersion
	op.Watch = true

	ready := false
	wch, err := deployClient.Watch(context.Background(), op)
	if err != nil {
		servererror.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	for event := range wch.ResultChan() {
		if status, err := utils.DeploymentsEventStatusFromRuntime(event.Object); err != nil {
			klog.Error(err)
			continue
		} else if ready = status == utils.StatusReady; ready {
			break
		}
	}

	if !ready {
		servererror.HandleError(c, http.StatusInternalServerError, errors.New("restart deployment failed"))
		return
	}

	c.JSON(200, nil)
}
