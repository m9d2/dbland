package model

import "time"

type Chart struct {
	Id          *uint      `db:"id" json:"id"`
	Title       *string    `db:"title" json:"title"`
	Sql         *string    `db:"sql" json:"sql"`
	Type        *string    `db:"type" json:"type"`
	CreatedTime *time.Time `db:"created_time" json:"created_time"`
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
