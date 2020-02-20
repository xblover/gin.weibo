package user

import (
	"net/http"
	"strconv"

	"gin.weibo/app/models"
	viewmodels "gin.weibo/app/view_models"
	"github.com/gin-gonic/gin"

	"gin.weibo/app/controllers"
	"gin.weibo/app/requests"
	"gin.weibo/pkg/flash"
)

// Index 用户列表
func Index(c *gin.Context) {
	controllers.Render(c, "user/index.html", gin.H{
		"my": "user index",
	})
}

// Create 创建用户页面
func Create(c *gin.Context) {
	controllers.Render(c, "user/create.html", gin.H{})
}

// Show 用户详情
func Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusOK, "参数错误 %v", err)
		return
	}

	m := &models.User{}
	user, err := m.Get(id)

	c.HTML(http.StatusOK, "user/show.html", viewmodels.NewUserViewModelSerializer(user, 140))
}

// Store 保存用户
func Store(c *gin.Context) {
	userForm := &requests.UserForm{
		Name:                 c.PostForm("name"),
		Email:                c.PostForm("email"),
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("password_confirmation"),
	}

	errors := userForm.Validate()

	if len(errors) != 0 {
		controllers.Render(c, "user/create.html", gin.H{
			"errors": errors,
		})
	}

	flash.NewSuccessFlash(c, "啦啦啦啦写入 flash 成功啦")
	controllers.Redirect(c, "http://localhost:8888/users/create")
}

// 编辑用户界面
func Edit(c *gin.Context) {

}

// 编辑用户
func Update(c *gin.Context) {

}

// 删除用户
func Destroy(c *gin.Context) {

}
