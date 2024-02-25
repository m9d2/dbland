package repository

import (
	"dbland/model"
)

type ExecutionLogRepository struct {
}

func (r ExecutionLogRepository) Save(log *model.Log) {
	DB.Save(log)
}
