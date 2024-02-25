package model

import "time"

type User struct {
	Id            *uint      `json:"id"`
	Username      *string    `json:"username"`
	Password      *string    `json:"password"`
	Email         *string    `json:"email"`
	Name          *string    `json:"name"`
	Avatar        *string    `json:"avatar"`
	Status        int        `json:"status"`
	LastLoginIp   *string    `json:"last_login_ip"`
	LastLoginTime *time.Time `json:"last_login_time"`
}

func NewUser() *User {
	return &User{
		Id:            nil,
		Username:      nil,
		Password:      nil,
		Email:         nil,
		Name:          nil,
		Avatar:        nil,
		Status:        1,
		LastLoginIp:   nil,
		LastLoginTime: nil,
	}
}

func (u User) TableName() string {
	return "sys_user"
}
