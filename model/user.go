package model

import "time"

type User struct {
	Id            *uint      `db:"id"`
	Username      *string    `db:"username"`
	Password      *string    `db:"password"`
	Email         *string    `db:"email"`
	Name          *string    `db:"name"`
	Avatar        *string    `db:"avatar"`
	Status        int        `db:"status"`
	LastLoginIp   *string    `db:"last_login_ip"`
	LastLoginTime *time.Time `db:"last_login_time"`
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
