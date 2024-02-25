package model

import "time"

type Chart struct {
	Id          *uint      `json:"id"`
	Title       *string    `json:"title"`
	Sql         *string    `json:"sql"`
	Type        *string    `json:"type"`
	CreatedTime *time.Time `json:"created_time"`
}

func NewChart() *Chart {
	return &Chart{
		Id:          nil,
		Title:       nil,
		Sql:         nil,
		Type:        nil,
		CreatedTime: nil,
	}
}

func (c Chart) TableName() string {
	return "sys_chart"
}
