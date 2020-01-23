package routes

import (
	"net/http"

	"gin.weibo/app/controllers/home"
	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine) *gin.Engine {
	// -------------------------------------注册全局中间件---------------------------------------------
	g.Use(gin.Recovery())
	g.Use(gin.Logger())

	// --------------------------------------注册路由---------------------------------------------
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	// root
	g.GET("/", home.Index)
	g.GET("/2", home.Index2)

	return g
}
