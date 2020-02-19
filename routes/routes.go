package routes

import (
	"net/http"

	"gin.weibo/middleware"

	"github.com/gin-gonic/gin"
	ginSessions "github.com/tommy351/gin-sessions"
)

var (
	sessionKeyPairs  = []byte("secret123")
	sessionStoreName = "my_session"
)

// Register 注册路由和中间件
func Register(g *gin.Engine) *gin.Engine {
	// -------------------------------------注册全局中间件---------------------------------------------
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	//session
	store := ginSessions.NewCookieStore(sessionKeyPairs)
	g.Use(ginSessions.Middleware(sessionStoreName, store))
	// 自定义全局中间件
	g.Use(middleware.OldValue()) // 记忆上次表单提交的内容，消费即消失

	// --------------------------------------注册路由---------------------------------------------
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	// web
	registerWeb(g)
	//api
	registerApi(g)

	return g
}
