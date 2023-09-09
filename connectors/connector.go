package connectors

import (
	"dbland/model"
	"github.com/jmoiron/sqlx"
)

const (
	Mysql      = "mysql"
	Oracle     = "oracle"
	SQLite     = "sqlite"
	PostgreSQL = "postgreSQL"
)

type Connector interface {
	Ping(config *model.ConnectionConfig) error
	Connect(config *model.ConnectionConfig) (*sqlx.DB, error)
	ShowDatabases(db *sqlx.DB) ([]string, error)
	Use(db *sqlx.DB, database string) error
	ShowTables(db *sqlx.DB, database string) (*[]Table, error)
	// Column params order -
	// mysql: schema, table
	//
	// sqlite: table
	Column(db *sqlx.DB, params ...string) (*[]Column, error)
	Query(db *sqlx.DB, sqlStr string) (*Query, error)
	Count(db *sqlx.DB, sqlStr string) (int, error)
	Ddl(db *sqlx.DB, table string) (string, error)
}
