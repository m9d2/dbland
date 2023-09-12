package connectors

import (
	"dbland/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MysqlConnector struct {
	DefaultConnector
}

func (c MysqlConnector) Ping(config *model.ConnectionConfig) error {
	db, err := sqlx.Open("mysql", c.dsn(config))
	if err != nil {
		return err
	}
	err = db.Ping()
	return err
}

func (c MysqlConnector) Connect(config *model.ConnectionConfig) (*sqlx.DB, error) {
	return sqlx.Open("mysql", c.dsn(config))
}

func (c DefaultConnector) Use(db *sqlx.DB, database string) error {
	sqlStr := "use " + database
	_, err := db.Exec(sqlStr)
	return err
}

func (c MysqlConnector) dsn(database *model.ConnectionConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", *database.Username, *database.Password, *database.Host, *database.Port)
}
