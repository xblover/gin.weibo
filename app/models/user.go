package models

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	"gin.weibo/database"
)

// 用户模型
type User struct {
	BaseModel
	Name            string    `gorm:"column:name;type:varchar(255);not null"`
	Email           string    `gorm:"column:email;type:varchar(255);unique;not null"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at"`
	Password        string    `gorm:"column:password;type:varchar(255);not null"`
	RememberToken   string    `gorm:"column:remember_token;type:varchar(100)"`
	IsAdmin         uint      `gorm:"column:is_admin;type:tinyint(1)"`
	ActivationToken string    `gorm:"activation_token;type:varchar(255)"`
	Activated       uint      `gorm:"column:activated;type:tinyint(1);not null"`
}

// TableName 表名
func (User) TableName() string {
	return "users"
}

// Get 获取一个用户
func (User) Get(id int) (*User, error) {
	u := &User{}
	d := database.DB.First(&u, id)
	return u, d.Error
}

// GetByEmail 根据email 来获取用户
func (User) GetByEmail(email string) (*User, error) {
	u := &User{}
	d := database.DB.Where("email = ?", email).First(&u)
	return u, d.Error
}

// 生成用户头像
func (u *User) Gravatar(size int) string {
	hash := md5.Sum([]byte(u.Email))
	return "http://www.gravatar.com/avatar/" + hex.EncodeToString(hash[:]) + "?s=" + strconv.Itoa(size)
}
