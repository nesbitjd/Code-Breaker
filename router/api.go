package router

import (
	"Projects/code_breaker/api"

	"github.com/gin-gonic/gin"
)

// ChannelHandlers adds channel handlers to the router group
func APIHandlers(base *gin.RouterGroup) {
	apibase := base.Group("/")
	{
		apibase.POST("/create", api.Create)
		apibase.DELETE("/delete/:code", api.Delete)
	}
}
