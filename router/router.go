package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yufenghui/apiserver/handler/sd"
	"github.com/yufenghui/apiserver/handler/user"
	"github.com/yufenghui/apiserver/router/middleware"
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

	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)
	}

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
	}

	return g
}
