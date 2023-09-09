package repository

import (
	"dbland/model"
	"github.com/jmoiron/sqlx"
	"time"
)

type UserRepository struct {
}

func (r UserRepository) FindByUsername(tx *sqlx.Tx, username string) (*model.User, error) {
	var user model.User
	sql := "SELECT * FROM user WHERE username = $1"
	err := tx.Get(&user,
		sql,
		username)
	return &user, err
}

func (r UserRepository) FindAll(tx *sqlx.Tx) (*[]model.User, error) {
	return nil, nil
}

func (r UserRepository) Save(tx *sqlx.Tx, user *model.User) {
	sql := "INSERT INTO user(name, password, email, username, avatar, status, last_login_ip, last_login_time) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"
	tx.MustExec(sql,
		user.Name,
		user.Password,
		user.Email,
		user.Username,
		user.Avatar,
		user.Status,
		user.LastLoginIp,
		user.LastLoginTime)
}

func (r UserRepository) UpdateLastLoginTime(tx *sqlx.Tx, id *uint) error {
	sql := "UPDATE user SET last_login_time = $1 WHERE id = $2"
	_, err := tx.Exec(sql, time.Now(), id)
	return err
}
