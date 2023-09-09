package connectors

type Column struct {
	ColumnName    string `db:"COLUMN_NAME" json:"column_name"`
	ColumnType    string `db:"COLUMN_TYPE" json:"column_type"`
	ColumnComment string `db:"COLUMN_COMMENT" json:"column_comment"`
}

type Query struct {
	Columns     []string         `json:"columns"`
	TableName   string           `json:"table_name"`
	Rows        []map[string]any `json:"rows"`
	Total       int              `json:"total"`
	ElapsedTime float64          `json:"elapsed_time"`
}

type Table struct {
	Name       string `json:"name,omitempty"`
	Rows       string `json:"rows,omitempty"`
	Collation  string `json:"collation,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
}
