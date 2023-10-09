package api

import (
	"dbland/service"

	"github.com/gin-gonic/gin"
)

type ChartHandler struct {
	service service.ChartService
}

func (s ChartHandler) InitRouter(g *gin.RouterGroup) {
	g.POST("chart", s.add)
	g.PUT("chart", s.update)
	g.GET("chart", s.list)
	g.DELETE("chart/:id", s.delete)
}

func (s ChartHandler) list(c *gin.Context) {
	data, err := s.service.List(c)
	if err != nil {
		JSON(c, err)
	} else {
		JSON(c, data)
	}
}

func (s ChartHandler) add(c *gin.Context) {
	err := s.service.Save(c)
	JSON(c, err)
}

func (s ChartHandler) update(c *gin.Context) {
	err := s.service.Update(c)
	JSON(c, err)
}

func (s ChartHandler) delete(c *gin.Context) {
	err := s.service.Delete(c)
	JSON(c, err)
}
