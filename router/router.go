package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thalkz/promo_code/endpoints"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Get user value
	r.PUT("verify", endpoints.HandleVerify)
	r.POST("add", endpoints.HandleAdd)

	return r
}
