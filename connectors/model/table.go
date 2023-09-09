package model

type Table struct {
	Name       string `json:"name,omitempty"`
	Rows       string `json:"rows,omitempty"`
	Collation  string `json:"collation,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
}
