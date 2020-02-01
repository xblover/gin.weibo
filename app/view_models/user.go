package viewmodels

import (
	"gin.weibo/app/models"
)

//用户
type UserViewModel struct {
	ID     int
	Name   string
	Email  string
	Avatar string
}

//用户数据展示
func NewUserViewModelSerializer(u *models.User) UserViewModel {
	return UserViewModel{
		ID:     int(u.ID),
		Name:   u.Name,
		Email:  u.Email,
		Avatar: "http://www.gravatar.com/avatat/",
	}
}
