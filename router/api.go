package router

import (
	"github.com/nesbitjd/hangle_server/api/record"
	"github.com/nesbitjd/hangle_server/api/user"
	"github.com/nesbitjd/hangle_server/api/word"

	"github.com/gin-gonic/gin"
)

// ChannelHandlers adds channel handlers to the router group
func APIHandlers(base *gin.RouterGroup) {
	apibase := base.Group("/")
	{
		// API endpoints for record table
		apibase.POST("/record", record.Create)
		apibase.DELETE("/record/:id", record.Delete)
		apibase.PUT("/record/:id", record.Update)
		apibase.GET("/record/:id", record.Read)

		// API endpoints for user table
		apibase.POST("/user", user.Create)
		apibase.DELETE("/user/:id", user.Delete)
		apibase.PUT("/user/:id", user.Update)
		apibase.GET("/user/:id", user.Read)

		// API endpoints for word table
		apibase.POST("/word", word.Create)
		apibase.DELETE("/word/:id", word.Delete)
		apibase.PUT("/word/:id", word.Update)
		apibase.GET("/word/:id", word.Read)
	}
}
