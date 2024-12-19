package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int            `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"column:name;size:255" json:"name"`
	Email     string         `gorm:"column:email;size:255" json:"email"`
	Password  string         `gorm:"column:password;size:255" json:"password"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at"`
}

func (User) TableName() string {
	return "users" // 自定义的表名
}
