package config

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	Conf       = Config{}
	dbFile     = "dbland.db"
	dataPath   = "./data"
	serverPath = "v1"
)

type Config struct {
	Sqlite struct {
		DataPath string `yaml:"data-path"`
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
	if Conf.Sqlite.DataPath == "" {
		Conf.Sqlite.DataPath = dataPath
	}
	_ = os.MkdirAll(Conf.Sqlite.DataPath, os.ModePerm)
	Conf.Sqlite.DataPath += "/" + dbFile

	if Conf.Server.Environment == "" {
		Conf.Server.Environment = gin.ReleaseMode
	}
	if Conf.Server.Path == "" {
		Conf.Server.Path = serverPath
	}
}
