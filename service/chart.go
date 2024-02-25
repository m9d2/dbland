package service

import (
	"dbland/model"
	"dbland/repository"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ChartService struct {
	repository repository.ChartRepository
}

func (s ChartService) List(g *gin.Context) *[]model.Chart {
	return s.repository.List()
}

func (s ChartService) Save(g *gin.Context) error {
	chart := model.NewChart()
	err := g.ShouldBind(&chart)
	if err != nil {
		return err
	}
	return s.repository.Save(chart)
}

func (s ChartService) Update(g *gin.Context) error {
	chart := model.NewChart()
	err := g.ShouldBind(&chart)
	if err != nil {
		return err
	}
	return s.repository.Update(chart)
}

func (s ChartService) Delete(g *gin.Context) error {
	idStr := g.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	return s.repository.DeleteById(id)
}
