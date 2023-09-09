package api

import (
	"dbland/service"
	"github.com/gin-gonic/gin"
)

type ConnectionConfigHandler struct {
	connectionConfigService service.ConnectionConfigService
}

func (s ConnectionConfigHandler) InitRouter(g *gin.RouterGroup) {
	g.POST("config", s.add)
	g.PUT("config", s.update)
	g.GET("config", s.list)
	g.DELETE("config/:id", s.delete)
}

func (s ConnectionConfigHandler) list(c *gin.Context) {
	data, err := s.connectionConfigService.List(c)
	if err != nil {
		JSON(c, err)
	} else {
		JSON(c, data)
	}
}

func (s ConnectionConfigHandler) add(c *gin.Context) {
	err := s.connectionConfigService.Save(c)
	JSON(c, err)
}

func (s ConnectionConfigHandler) update(c *gin.Context) {
	err := s.connectionConfigService.Update(c)
	JSON(c, err)
}

func (s ConnectionConfigHandler) delete(c *gin.Context) {
	err := s.connectionConfigService.Delete(c)
	JSON(c, err)
}
