package router

import (
	"Projects/code_breaker/api"

	"github.com/gin-gonic/gin"
)

// ChannelHandlers adds channel handlers to the router group
func APIHandlers(base *gin.RouterGroup) {
	apibase := base.Group("/")
	{
		apibase.POST("/product", api.Create)
		apibase.DELETE("/product/:code", api.Delete)
		apibase.PUT("/product/:code", api.Update)
	}
}
