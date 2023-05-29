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

	// /apis/cluster/v1/
	routerGroupV1 := routerGroup.Group("/v1")
	{
		routerGroupV1.POST("/code/:clusterCode", h.addCluster)
		routerGroupV1.DELETE("/code/:clusterCode", h.removeCluster)
		routerGroupV1.PUT("/code/:clusterCode", h.updateCluster)
		routerGroupV1.GET("/code/:clusterCode", h.getCluster)
		routerGroupV1.GET("/", h.getClusters)
		routerGroupV1.PATCH("/code/:clusterCode", h.applyCluster)
	}

}
