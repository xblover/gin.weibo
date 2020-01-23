package routes

import (
	"gin.weibo/app/controllers/home"
	"github.com/gin-gonic/gin"
)

func registerWeb(g *gin.Engine) {
	//root
	g.GET("/", home.Index)
	g.GET("/2", home.Index2)
}
