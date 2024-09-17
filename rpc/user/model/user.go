package model

import (
	"gorm.io/gorm"
)

// Account 账号
type Account struct {
	gorm.Model
	UserId      string `json:"user_id"  column:"user_id"`
	UserAccount string `json:"user_account" column:"user_account"`
	Category    int64  `json:"category" column:"category"`
}

func (u *Account) TableName() string {
	return "user_info.account"
}

type User struct {
	gorm.Model
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	Avatar      []byte `json:"avatar"`
	PhoneNumber string `json:"phone_number"`
	Salt        string `json:"salt"`
	Password    string `json:"password"`
}

func (u *User) TableName() string {
	return "user_info.user"
}
