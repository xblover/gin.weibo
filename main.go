package main

import (
	"html/template"
	"log"
	"net/http"

	"gin.weibo/app/helpers"
	"gin.weibo/app/models"
	"gin.weibo/config"
	"gin.weibo/database"
	"gin.weibo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化配置
	config.InitConfig()

	//gin config
	// gin.SetMode(config.AppConfig.RunMode)
	g := gin.New()
	setupGin(g)

	//db config
	db := database.InitDB()
	db.AutoMigrate(
		&models.User{},
	)
	defer db.Close()

	//router config
	routes.Register(g)

	//启动
	log.Printf("Start to listening the incoming requests on http address: %s", config.AppConfig.Addr)
	log.Fatal(http.ListenAndServe(config.AppConfig.Addr, g).Error())
}

//配置 gin
func setupGin(g *gin.Engine) {
	//启动模式配置
	gin.SetMode(config.AppConfig.RunMode)

	// 项目静态文件配置
	g.Static("/"+config.ProjectConfig.PublicPath, config.ProjectConfig.PublicPath)
	g.StaticFile("/favicon.ico", config.ProjectConfig.PublicPath+"/favicon.ico")

	// 模板配置
	// 注册模板函数
	g.SetFuncMap(template.FuncMap{
		"Mix":       helpers.Mix,
		"CsrfField": helpers.CsrfField,
	})
	g.LoadHTMLGlob("resources/views/**/*")
}
