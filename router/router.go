package router

import (
	"apiserver/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/yufenghui/apiserver/handler/sd"
	"net/http"
)

// load middlewares, routes, handlers
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// Middlewares
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not found!")
	})

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
	}

	return g
}
