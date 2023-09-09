package repository

import (
	"dbland/config"
	"dbland/connectors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
	"os"
)

var DB *sqlx.DB

func Initialize() {
	var err error
	switch config.Conf.Datasource {
	case connectors.SQLite:
		DB, err = sqlx.Open("sqlite3", config.Conf.Sqlite.DBfile)
		break
	case connectors.Mysql:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Conf.Mysql.Username,
			config.Conf.Mysql.Password,
			config.Conf.Mysql.Host,
			config.Conf.Mysql.Port,
			config.Conf.Mysql.Database)
		DB, err = sqlx.Open("mysql", dsn)
		break
	}

	if err != nil {
		slog.Error(err.Error())
		return
	}
	if config.Conf.InitSchema {
		err = executeSQLScript(DB, getSchema())
		if err != nil {
			slog.Error(err.Error())
			return
		}
	}
}

func executeSQLScript(db *sqlx.DB, scriptPath string) error {
	script, err := os.ReadFile(scriptPath)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(script))
	if err != nil {
		return err
	}
	return nil
}

func getSchema() string {
	if config.Conf.Datasource == connectors.SQLite {
		return "schema-sqlite.sql"
	}
	if config.Conf.Datasource == connectors.Mysql {
		return "schema-mysql.sql"
	}
	return ""
}
