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

	controllers.Render(c, "user/show.html", gin.H{
		"userData": viewmodels.NewUserViewModelSerializer(user, 140),
	})
}

// Store 保存用户
func Store(c *gin.Context) {
	//验证参数
	userForm := &requests.UserForm{
		Name:                 c.PostForm("name"),
		Email:                c.PostForm("email"),
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("password_confirmation"),
	}

	user, errors := userForm.ValidateAndSave()

	if len(errors) != 0 {
		flash.SaveValidateMessage(c, errors)
		controllers.Redirect(c, "/users/create")
		return
	}

	controllers.Redirect(c, "/users/create")
	flash.NewSuccessFlash(c, "欢迎，您将在这里开启一段新的旅程~")
	controllers.Redirect(c, "/users/show/"+strconv.Itoa(int(user.ID)))
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
