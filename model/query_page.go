package model

type QueryPage struct {
	Columns   []string                 `json:"columns"`
	TableName string                   `json:"table_name"`
	Rows      []map[string]interface{} `json:"rows"`
	Total     uint                     `json:"total"`
	PageNo    int                      `json:"page_no"`
}
