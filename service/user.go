package service

import (
	"dbland/model"
	"dbland/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (s UserService) Add(c *gin.Context) error {
	user := model.NewUser()
	err := c.ShouldBind(&user)
	if err != nil {
		return err
	}
	var pwd []byte
	pwd, err = bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	hashedPassword := string(pwd)
	user.Password = &hashedPassword

	tx := repository.DB.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	s.userRepository.Save(tx, user)
	return nil
}

// Login user login
func (s UserService) Login(c *gin.Context) error {
	var err error
	username := c.PostForm("username")
	password := c.PostForm("password")

	tx := repository.DB.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	var user *model.User
	user, err = s.userRepository.FindByUsername(tx, username)
	if err != nil {
		return err
	}

	if user == nil {
		slog.Error("User does not exist", "username", username)
		return fmt.Errorf("user does not exist")
	}

	err = s.userRepository.UpdateLastLoginTime(tx, user.Id)
	if err != nil {
		return err
	}

	// verify password
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		slog.Error("Password error", "username", username, "password", password)
		return fmt.Errorf("password error")
	}

	return err
}

func (s UserService) Logout(c *gin.Context) {
}
