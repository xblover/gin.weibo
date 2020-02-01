package models

import "time"

// BaseModel model 基类
type BaseModel struct {
	ID uint `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	// MySQL 的DATE/DATETIME 类型可以对应Golang的time.Time
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// 有DeletedAt(类型需要是 *time.Time) 即支持gorm软删除
	DeleteAt *time.Time `gorm:"column:deleted_at" sql:"index"`
}
