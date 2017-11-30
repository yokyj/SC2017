package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	UID        int        `orm:"id,auto-inc" gorm:"primary_key;AUTO_INCREMENT;column:uid"`
	UserName   string     `gorm:"column:username"`
	DepartName string     `gorm:"column:departname"`
	CreateAt   *time.Time `gorm:"column:created"`
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName shold not null!")
	}
	if u.CreateAt == nil {
		t := time.Now()
		u.CreateAt = &t
	}
	return &u
}

func (u *UserInfo) TableName() string {
	return "userinfo"
}