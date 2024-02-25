package service

import (
	"dbland/model"
	"dbland/repository"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ConnectionConfigService struct {
	connectionConfigRepository repository.ConnectionConfigRepository
}

func (s ConnectionConfigService) List(g *gin.Context) *[]model.Config {
	return s.connectionConfigRepository.List()
}

func (s ConnectionConfigService) Save(g *gin.Context) error {
	config := model.NewConfig()
	err := g.ShouldBind(&config)
	if err != nil {
		return err
	}
	s.connectionConfigRepository.Save(config)
	return nil
}

func (s ConnectionConfigService) Update(g *gin.Context) error {
	config := model.NewConfig()
	err := g.ShouldBind(&config)
	if err != nil {
		return err
	}
	s.connectionConfigRepository.Update(config)
	return nil
}

func (s ConnectionConfigService) Delete(g *gin.Context) error {
	idStr := g.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	s.connectionConfigRepository.DeleteById(uint(id))
	return nil
}
