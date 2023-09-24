package repository

import (
	"dbland/model"
	"github.com/jmoiron/sqlx"
)

type ExecutionLogRepository struct {
}

func (r ExecutionLogRepository) Save(tx *sqlx.Tx, log *model.ExecutionLog) {
	sql := "INSERT INTO execution_log (id, source, database, sql, status, cost, ip, user_agent, created_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
	tx.MustExec(sql, log.Id, log.Source, log.Database, log.Sql, log.Status, log.Cost, log.Ip, log.UserAgent, log.CreatedTime)
}
