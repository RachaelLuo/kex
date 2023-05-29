package server

import (
	"github.com/gin-gonic/gin"

	"github.com/RachaelLuo/kex/pkg/zone/clientset"
)

type handler struct {
	client clientset.Interface
}

func InstallHandlers(routerGroup *gin.RouterGroup, client clientset.Interface) {
	h := &handler{
		client: client,
	}

	// /apis/server/v1/
	routerGroupV1 := routerGroup.Group("/v1")
	{
		// Proxy cluster for all native api
		routerGroupV1.Any("/proxy/cluster/:clusterCode/*urlPath", h.proxyCluster)

		// node
		routerGroupV1.GET("/cluster/:clusterCode/node/:name/namespace", h.listNamespaceOfNode)

		// deployment
		routerGroupV1.GET("/cluster/:clusterCode/namespace/:namespace/deployments/:name/pods", h.listPodOfDeployment)
		routerGroupV1.POST("/cluster/:clusterCode/namespace/:namespace/deployments/:name/restart", h.restartDeployment)
	}
}
