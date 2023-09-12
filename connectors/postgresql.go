package connectors

import (
	"database/sql"
	"dbland/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreSQLConnector struct {
	DefaultConnector
}

func (c PostgreSQLConnector) Ping(config *model.ConnectionConfig) error {
	db, err := sqlx.Open("postgres", c.dsn(config))
	if err != nil {
		return err
	}
	err = db.Ping()
	return err
}

func (c PostgreSQLConnector) Connect(config *model.ConnectionConfig) (*sqlx.DB, error) {
	return sqlx.Open("postgres", c.dsn(config))
}

func (c PostgreSQLConnector) ShowDatabases(db *sqlx.DB) ([]string, error) {
	var result []string
	sqlStr := "SELECT datname FROM pg_database WHERE datistemplate = false"
	rows, err := db.Queryx(sqlStr)
	defer func() {
		_ = rows.Close()
		_ = db.Close()
	}()
	for rows.Next() {
		var row string
		err = rows.Scan(&row)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, err
}

func (c PostgreSQLConnector) ShowTables(db *sqlx.DB, database string) (*[]Table, error) {
	var tables []Table
	sqlStr := "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'"
	rows, err := db.Queryx(sqlStr, database)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
		_ = db.Close()
	}()

	for rows.Next() {
		var table Table
		var tableName sql.NullString
		var err = rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}

		if tableName.Valid {
			table.Name = tableName.String
		}

		tables = append(tables, table)
	}
	return &tables, err
}

func (c PostgreSQLConnector) dsn(database *model.ConnectionConfig) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable ", *database.Username, *database.Password, *database.Host, *database.Port)
}
