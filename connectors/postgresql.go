package connectors

import (
	"dbland/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PostgreSQLConnector struct {
}

func (c PostgreSQLConnector) Ping(config *model.ConnectionConfig) error {
	return nil
}

func (c PostgreSQLConnector) Connect(config *model.ConnectionConfig) (*sqlx.DB, error) {
	return sqlx.Open("mysql", c.dsn(config))
}

func (c PostgreSQLConnector) ShowDatabases(db *sqlx.DB) ([]string, error) {
	return nil, nil
}

func (c PostgreSQLConnector) Use(db *sqlx.DB, database string) error {
	return nil
}

func (c PostgreSQLConnector) ShowTables(db *sqlx.DB, database string) (*[]Table, error) {
	return nil, nil
}

func (c PostgreSQLConnector) Column(db *sqlx.DB, params ...string) (*[]Column, error) {
	return nil, nil
}

func (c PostgreSQLConnector) Query(db *sqlx.DB, sqlStr string) (*Query, error) {
	return nil, nil
}
func (c PostgreSQLConnector) Count(db *sqlx.DB, sqlStr string) (int, error) {
	rows, err := db.Queryx(sqlStr)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = rows.Close()
		_ = db.Close()
	}()
	for rows.Next() {
		var total int
		err = rows.Scan(total)
		if err != nil {
			return total, nil
		}
	}
	return 0, err
}

func (c PostgreSQLConnector) Ddl(db *sqlx.DB, table string) (string, error) {
	return "", nil
}

func (c PostgreSQLConnector) dsn(database *model.ConnectionConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", database.Username, database.Password, database.Host, database.Port)
}
