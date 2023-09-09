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

func (s ConnectionConfigService) List(g *gin.Context) (*[]model.ConnectionConfig, error) {
	return s.connectionConfigRepository.List()
}

func (s ConnectionConfigService) Save(g *gin.Context) error {
	config := model.NewConnectionConfig()
	err := g.ShouldBind(&config)
	if err != nil {
		return err
	}
	tx := repository.DB.MustBegin()
	defer func() {
		if err != nil {
			err = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	s.connectionConfigRepository.Save(tx, config)
	return nil
}

func (s ConnectionConfigService) Update(g *gin.Context) error {
	config := model.NewConnectionConfig()
	err := g.ShouldBind(&config)
	if err != nil {
		return err
	}
	tx := repository.DB.MustBegin()
	defer func() {
		if err != nil {
			err = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	s.connectionConfigRepository.Update(tx, config)
	return nil
}

func (s ConnectionConfigService) Delete(g *gin.Context) error {
	idStr := g.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	tx := repository.DB.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	s.connectionConfigRepository.DeleteById(tx, uint(id))
	return nil
}
