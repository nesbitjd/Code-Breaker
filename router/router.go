package router

import (
	"Projects/hangle_server/api"

	"github.com/gin-gonic/gin"
)

const (
	base = "/api/v1"
)

// Load is a server function that returns the engine for processing web requests
// on the host it's running on.
func Load(options ...gin.HandlerFunc) *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/health", api.Health)

	baseAPI := r.Group(base)
	{
		APIHandlers(baseAPI)
	}

	return r
}
