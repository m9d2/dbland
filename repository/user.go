package repository

import (
	"dbland/model"
)

type UserRepository struct {
}

func (r UserRepository) FindByUsername(username string) *model.User {
	var user model.User
	DB.Where("username = ?", username).First(&user)
	return &user
}

func (r UserRepository) FindAll() (*[]model.User, error) {
	return nil, nil
}

func (r UserRepository) Save(user *model.User) {

}

func (r UserRepository) UpdateLastLoginTime(id *uint) error {
	return nil
}
