package models

import (
	"time"
)

type User struct {
	ID          int64     `json:"id" gorm:"id"`
	Account     string    `json:"account" gorm:"account"`
	UserName    string    `json:"user_name" gorm:"user_name"`
	Password    string    `json:"password" gorm:"password"`
	Mobile      string    `json:"mobile" gorm:"mobile"`
	HeadImg     string    `json:"head_img" gorm:"head_img"` // 头像
	Status      int8      `json:"status" gorm:"status"`     // 1;正常 2：禁用
	Mark        string    `json:"mark" gorm:"mark"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
	LastLoginAt time.Time `json:"last_login_at" gorm:"last_login_at"` // 最后登录时间
	Avatar      string    `json:"avatar" gorm:"avatar"`
}

// TableName 表名称
func (*User) TableName() string {
	return "dg_user"
}