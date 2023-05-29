package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/klog"

	"github.com/RachaelLuo/kex/pkg/server/servererror"
	"github.com/RachaelLuo/kex/pkg/zone/proxy"
)

func (h *handler) proxyCluster(c *gin.Context) {
	client, err := proxy.GetClusterPorxyClientFromCode(c.Param("clusterCode"))
	if err != nil {
		servererror.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	request := client.RESTClient().Verb(c.Request.Method).AbsPath(c.Param("urlPath")).Body(c.Request.Body)
	for k, v := range c.Request.URL.Query() {
		var s string
		for i := range v {
			if i != 0 {
				s += ","
			}
			s += v[i]
		}
		request.Param(k, s)
	}
	for k, v := range c.Request.Header {
		request.SetHeader(k, v...)
	}

	var statusCode int
	data, err := request.Do(c).StatusCode(&statusCode).Raw()
	if err != nil {
		klog.Error(err)
		if serr, ok := err.(*apierrors.StatusError); ok {
			c.JSON(int(serr.ErrStatus.Code), err)
		} else {
			c.JSON(http.StatusInternalServerError, err)
		}
	}
	c.Data(statusCode, "application/json, text/plain, */*", data)
}
