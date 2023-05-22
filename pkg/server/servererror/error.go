package servererror

import (
	"github.com/gin-gonic/gin"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/klog"
)

func HandleError(c *gin.Context, code int, err error) {
	if c.IsAborted() || c.Writer.Size() > 0 {
		return
	}

	if err != nil {
		klog.Errorf("err: %+v", err)
	}

	switch t := err.(type) {
	case apierrors.APIStatus:
		c.JSON(int(t.Status().Code), err)
		return
	}
	response := map[string]string{}
	response["message"] = err.Error()
	c.JSON(code, response)
}
