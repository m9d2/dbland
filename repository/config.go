package repository

import (
	"dbland/model"
)

type ConnectionConfigRepository struct {
}

func (r ConnectionConfigRepository) List() *[]model.Config {
	var configs []model.Config
	DB.Find(&configs)
	return &configs
}

func (r ConnectionConfigRepository) GetById(id uint) *model.Config {
	var config model.Config
	DB.Where("id = ?", id).First(&config)
	return &config
}

func (r ConnectionConfigRepository) Save(config *model.Config) {
	DB.Save(config)
}

func (r ConnectionConfigRepository) Update(config *model.Config) {
	DB.Updates(config)
}

func (r ConnectionConfigRepository) DeleteById(id uint) {
	DB.Delete(model.Config{}, id)
}
