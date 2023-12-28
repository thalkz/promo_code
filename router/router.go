package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thalkz/promo_code/endpoints"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.PUT("add", endpoints.HandleAdd)
	r.GET("verify", endpoints.HandleVerify)

	return r
}
