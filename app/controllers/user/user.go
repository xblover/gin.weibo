package user

import (
	"net/http"
	"strconv"

	"gin.weibo/app/models"
	viewmodels "gin.weibo/app/view_models"
	"github.com/gin-gonic/gin"
)

//用户列表
func Index(c *gin.Context) {

}

// 创建用户页面
func Create(c *gin.Context) {
	c.HTML(http.StatusOK, "user/create.html", gin.H{})
}

// 用户详情
func Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusOK, "参数错误 %v", err)
		return
	}

	m := &models.User{}
	user, err := m.Get(id)

	c.HTML(http.StatusOK, "user/show.html", viewmodels.NewUserViewModelsSerializer(user, 140))
}

// 保存用户
func Store(c *gin.Context) {

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
