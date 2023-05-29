package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/RachaelLuo/kex/pkg/server/servererror"
	"github.com/RachaelLuo/kex/pkg/zone/proxy"
)

// listNamespaceOfNode get all namespace on node
func (h *handler) listNamespaceOfNode(c *gin.Context) {
	name := c.Param("name")
	client, err := proxy.GetClusterPorxyClientFromCode(c.Param("clusterCode"))
	if err != nil {
		servererror.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	podList, err := client.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{
		FieldSelector: "spec.nodeName=" + name,
	})
	if err != nil {
		servererror.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	namespacesMap := make(map[string]struct{})
	for i := range podList.Items {
		namespacesMap[podList.Items[i].Namespace] = struct{}{}
	}

	namespaces := make([]string, 0, len(namespacesMap))
	for namespace := range namespacesMap {
		namespaces = append(namespaces, namespace)
	}

	c.JSON(http.StatusOK, namespaces)
}
