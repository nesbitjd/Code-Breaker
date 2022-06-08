package router

import (
	"Projects/hangle_server/api"

	"github.com/gin-gonic/gin"
)

// ChannelHandlers adds channel handlers to the router group
func APIHandlers(base *gin.RouterGroup) {
	apibase := base.Group("/")
	{
		apibase.POST("/hangman", api.CreateWord)
		apibase.DELETE("/hangman/:word", api.Delete)
		apibase.PUT("/hangman/:id", api.Update)
		apibase.GET("/hangman/:id", api.Read)
	}
}
