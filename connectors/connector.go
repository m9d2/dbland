package connectors

import (
	"dbland/model"
	"dbland/repository"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	Mysql      = "mysql"
	Oracle     = "oracle"
	SQLite     = "sqlite"
	PostgreSQL = "postgresql"
)

type Connector interface {
	Ping(config *model.Config) error
	Connect(config *model.Config) (*sqlx.DB, error)
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
	Execute(db *sqlx.DB, sqlStr string) (int, error)
}

type ConnectorFactory struct {
	connectionConfigRepository repository.ConnectionConfigRepository
	mysqlConnector             MysqlConnector
	oracleConnector            OracleConnector
	sqLiteConnector            SQLiteConnector
	postgreSQLConnector        PostgreSQLConnector
}

func (c ConnectorFactory) GetInstance(dbType string) (Connector, error) {
	var connector Connector
	switch dbType {
	case Mysql:
		connector = c.mysqlConnector
	case Oracle:
		connector = c.oracleConnector
	case SQLite:
		connector = c.sqLiteConnector
	case PostgreSQL:
		connector = c.postgreSQLConnector

	default:
		err := fmt.Errorf("database not support")
		return nil, err
	}
	return connector, nil
}
