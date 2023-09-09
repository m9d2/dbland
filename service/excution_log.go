package service

import (
	"dbland/repository"
	"github.com/gin-gonic/gin"
)

type ExecutionLogService struct {
	executionLogRepository repository.ExecutionLogRepository
}

func (r ExecutionLogService) Save(c *gin.Context) error {
	return nil
}
