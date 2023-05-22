package auth

import (
	"github.com/gin-gonic/gin"
)

func MultiAuth(accounts gin.Accounts) gin.HandlerFunc {
	return func(c *gin.Context) {
		gin.BasicAuthForRealm(accounts, "")(c)
	}
}
