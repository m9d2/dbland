package config

import (
	"dbland/connectors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var Conf = Config{}

type Config struct {
	Datasource string
	Sqlite     struct {
		DBfile string `yaml:"db-file"`
	}
	Mysql struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
	Server struct {
		Environment string `yaml:"environment"`
		Path        string `yaml:"path"`
	}
	InitSchema bool `yaml:"init-schema"`
}

func InitConfig() {
	file, _ := os.ReadFile("./config.yaml")

	err := yaml.Unmarshal(file, &Conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	fillDefaultValues()
}

func fillDefaultValues() {
	if Conf.Datasource == "" {
		Conf.Datasource = "sqlite"
	}
	if Conf.Sqlite.DBfile == "" {
		Conf.Sqlite.DBfile = "dbland.db"
	}
	if Conf.Mysql.Database == "" {
		Conf.Mysql.Database = connectors.SQLite
	}
	if Conf.Server.Environment == "" {
		Conf.Server.Environment = gin.ReleaseMode
	}
	if Conf.Server.Path == "" {
		Conf.Server.Path = "v1"
	}
}
