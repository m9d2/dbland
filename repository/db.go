package repository

import (
	"dbland/config"
	"dbland/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
)

var DB *gorm.DB

func Initialize() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.Conf.Sqlite.DataPath), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		return
	}
	err = DB.AutoMigrate(&model.Chart{}, &model.Config{}, &model.Log{}, &model.User{})
	if err != nil {
		slog.Error(err.Error())
		return
	}
}
