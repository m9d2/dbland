package model

type Column struct {
	ColumnName    string `db:"COLUMN_NAME" json:"column_name"`
	ColumnType    string `db:"COLUMN_TYPE" json:"column_type"`
	ColumnComment string `db:"COLUMN_COMMENT" json:"column_comment"`
}
