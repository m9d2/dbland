package repository

import "dbland/model"

type ChartRepository struct {
}

func (r ChartRepository) List() *[]model.Chart {
	var charts []model.Chart
	DB.Find(&charts)
	return &charts
}

func (r ChartRepository) Save(chart *model.Chart) error {
	return nil
}

func (r ChartRepository) Update(chart *model.Chart) error {
	return nil
}

func (r ChartRepository) DeleteById(id int) error {
	return nil
}
