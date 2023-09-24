package repository

import (
	"dbland/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
)

var DB *sqlx.DB

func Initialize() {
	var err error
	DB, err = sqlx.Open("sqlite3", config.Conf.Sqlite.DataPath)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	err = executeSQL()
	if err != nil {
		slog.Error(err.Error())
	}
}

func executeSQL() error {
	sqlStr := `
		CREATE TABLE IF NOT EXISTS "user" (
			"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"name" TEXT,
			"password" TEXT,
			"email" TEXT,
			"username" TEXT,
			"avatar" TEXT,
			"status" integer,
			"last_login_ip" TEXT,
			"last_login_time" DATE
		);
		CREATE TABLE IF NOT EXISTS "config" (
			"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"type" TEXT,
			"name" TEXT,
			"host" TEXT,
			"port" TEXT,
			"username" TEXT,
			"password" TEXT,
			"db_file" TEXT
		);
		CREATE TABLE IF NOT EXISTS "execution_log" (
			"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"source" TEXT,
			"database" TEXT,
			"sql" TEXT,
			"status" integer,
			"cost" integer,
			"ip" TEXT,
			"user_agent" TEXT,
			"created_time" DATE
		);
	`
	_, err := DB.Exec(sqlStr)

	if err != nil {
		return err
	}
	return nil
}
