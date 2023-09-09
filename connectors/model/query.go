package model

type Query struct {
	Columns   []string         `json:"columns"`
	TableName string           `json:"table_name"`
	Rows      []map[string]any `json:"rows"`
	Total     int              `json:"total"`
}
