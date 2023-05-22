package clusters

import (
	"github.com/RachaelLuo/kex/pkg/zone/clientset"

	"github.com/gin-gonic/gin"
)

type handler struct {
	namespace         string
	clusterInfos      string
	localClusterInfos string
	client            clientset.Interface
}

func InstallHandlers(routerGroup *gin.RouterGroup, namespace, clusterInfos, localClusterInfos string, client clientset.Interface) {
	h := &handler{
		namespace:         namespace,
		clusterInfos:      clusterInfos,
		localClusterInfos: localClusterInfos,
		client:            client,
	}
	routerGroup.POST("/apis/clusters/:clusterCode", h.addCluster)
	routerGroup.DELETE("/apis/clusters/:clusterCode", h.removeCluster)
	routerGroup.PUT("/apis/clusters/:clusterCode", h.updateCluster)
	routerGroup.GET("/apis/clusters/:clusterCode", h.getCluster)
	routerGroup.GET("/apis/clusters", h.getClusters)
	routerGroup.PATCH("/apis/clusters/:clusterCode", h.applyCluster)
	routerGroup.Any("/apis/proxy/v1/clusters/:clusterCode/*urlPath", h.proxyCluster)
}
