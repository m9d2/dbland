package repository

import (
	"dbland/model"
	"github.com/jmoiron/sqlx"
)

type ConnectionConfigRepository struct {
}

func (r ConnectionConfigRepository) List() (*[]model.ConnectionConfig, error) {
	var configs []model.ConnectionConfig
	sql := "SELECT * FROM config"
	err := DB.Select(&configs, sql)
	return &configs, err
}

func (r ConnectionConfigRepository) GetById(id uint) (*model.ConnectionConfig, error) {
	var config model.ConnectionConfig
	sql := "SELECT id, type, name, host, port, username, password, db_file FROM config WHERE id = $1"
	err := DB.Get(&config, sql, id)
	return &config, err
}

func (r ConnectionConfigRepository) Save(tx *sqlx.Tx, config *model.ConnectionConfig) {
	sql := "INSERT INTO config(id, type, name, host, port, username, password, db_file) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	tx.MustExec(sql, config.Id, config.Type, config.Name, config.Host, config.Port, config.Username, config.Password, config.DbFile)
}

func (r ConnectionConfigRepository) Update(tx *sqlx.Tx, config *model.ConnectionConfig) {
	sql := "UPDATE config SET type = $1, name = $2, host = $3, port = $4, username = $5, password = $6, db_file = $7 WHERE id = $8"
	tx.MustExec(sql, config.Type, config.Name, config.Host, config.Port, config.Username, config.Password, config.DbFile, config.Id)
}

func (r ConnectionConfigRepository) DeleteById(tx *sqlx.Tx, id uint) {
	sql := "DELETE FROM config WHERE id = $1"
	tx.MustExec(sql, id)
}
