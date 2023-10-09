package repository

import "dbland/model"

type ChartRepository struct {
}

func (r ChartRepository) List() (*[]model.Chart, error) {
	var charts []model.Chart
	sql := "SELECT * FROM chart"
	err := DB.Select(&charts, sql)
	return &charts, err
}

func (r ChartRepository) Save(chart *model.Chart) error {
	sql := "INSERT INTO chart(id, title, type, sql, created_time) VALUES ($1, $2, $3, $4)"
	_, err := DB.Exec(sql, chart.Id, chart.Title, chart.Type, chart.Sql, chart.CreatedTime)
	if err != nil {
		return err
	}
	return nil
}

func (r ChartRepository) Update(chart *model.Chart) error {
	sql := "UPDATE chart SET title = $1, type = $2, sql = $3 where id = $3"
	_, err := DB.Exec(sql, chart.Title, chart.Type, chart.Sql, chart.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r ChartRepository) DeleteById(id int) error {
	sql := "DELETE from chart where id = $1"
	_, err := DB.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
