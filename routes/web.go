package routes

import (
	"os/user"

	"github.com/gin-gonic/gin"

	staticpage "gin.weibo/app/controllers/static_page"
	"gin.weibo/app/controllers/user"
)

func registerWeb(g *gin.Engine) {
	//static page
	g.GET("/", staticpage.Home)
	g.GET("/help", staticpage.Help)
	g.GET("about", staticpage.About)

	//user
	g.GET("/signup", user.Create)
	userRouter := g.Group("/users")
	{
		userRouter.GET("", user.Index)
		userRouter.GET("create", user.Create)
		userRouter.GET("/show/:id", user.Show)
		userRouter.POST("", user.Store)
		userRouter.GET("/edit/:id", user.Edit)
		userRouter.PATCH("/:id", user.Update)
		userRouter.DELETE("/:id", user.Destroy)
	}

}
