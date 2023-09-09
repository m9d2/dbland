package api

import (
	"dbland/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func (h UserHandler) InitRouter(g *gin.RouterGroup) {
	g.POST("login", h.login)
	g.POST("user", h.add)
	g.GET("logout", h.logout)
}

// login user login
func (h UserHandler) login(c *gin.Context) {
	err := h.userService.Login(c)
	JSON(c, err)
}

// login add user
func (h UserHandler) add(c *gin.Context) {
	err := h.userService.Add(c)
	JSON(c, err)
}

func (h UserHandler) logout(c *gin.Context) {
	h.userService.Logout(c)
	JSON(c, nil)
}
