package connectors

import (
	"database/sql"
	"dbland/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"strconv"
)

type MysqlConnector struct {
}

func (c MysqlConnector) Ping(config *model.ConnectionConfig) error {
	db, err := sqlx.Open("mysql", c.dsn(config))
	if err != nil {
		return err
	}
	err = db.Ping()
	return err
}

func (c MysqlConnector) Connect(config *model.ConnectionConfig) (*sqlx.DB, error) {
	return sqlx.Open("mysql", c.dsn(config))
}

func (c MysqlConnector) ShowDatabases(db *sqlx.DB) ([]string, error) {
	var result []string
	sqlStr := "show databases"
	rows, err := db.Queryx(sqlStr)
	defer func() {
		_ = rows.Close()
		_ = db.Close()
	}()
	for rows.Next() {
		var row string
		err = rows.Scan(&row)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, err
}

func (c MysqlConnector) Use(db *sqlx.DB, database string) error {
	sqlStr := "use " + database
	_, err := db.Exec(sqlStr)
	return err
}

func (c MysqlConnector) ShowTables(db *sqlx.DB, database string) (*[]Table, error) {
	var tables []Table
	sqlStr := `SELECT TABLE_NAME, TABLE_ROWS, TABLE_COLLATION, CREATE_TIME, UPDATE_TIME
			FROM INFORMATION_SCHEMA.TABLES
			WHERE TABLE_SCHEMA = ?`

	rows, err := db.Queryx(sqlStr, database)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
		_ = db.Close()
	}()

	for rows.Next() {
		var table Table
		var tableName sql.NullString
		var tableRows sql.NullString
		var tableCollation sql.NullString
		var createTime sql.NullString
		var updateTime sql.NullString
		var err = rows.Scan(&tableName, &tableRows, &tableCollation, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		if tableName.Valid {
			table.Name = tableName.String
		}
		if tableRows.Valid {
			table.Rows = tableRows.String
		}
		if tableCollation.Valid {
			table.Collation = tableCollation.String
		}
		if createTime.Valid {
			table.CreateTime = createTime.String
		}
		if updateTime.Valid {
			table.UpdateTime = updateTime.String
		}

		tables = append(tables, table)
	}
	return &tables, err
}

func (c MysqlConnector) Column(db *sqlx.DB, params ...string) (*[]Column, error) {
	var columns []Column
	schema := params[0]
	table := params[1]
	sqlStr := `SELECT COLUMN_NAME, COLUMN_TYPE, COLUMN_COMMENT
			FROM INFORMATION_SCHEMA.COLUMNS
			WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?`

	rows, err := db.Queryx(sqlStr, schema, table)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = db.Close()
	}()

	for rows.Next() {
		var column Column
		var columnName sql.NullString
		var columnType sql.NullString
		var columnComment sql.NullString
		err = rows.StructScan(&column)
		if err != nil {
			return nil, err
		}

		if columnName.Valid {
			column.ColumnName = columnName.String
		}
		if columnType.Valid {
			column.ColumnType = columnType.String
		}
		if columnComment.Valid {
			column.ColumnComment = columnComment.String
		}

		columns = append(columns, column)
	}
	return &columns, err
}

func (c MysqlConnector) Query(db *sqlx.DB, sqlStr string) (*Query, error) {
	rows, err := db.Queryx(sqlStr)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	columnTypes, err := rows.ColumnTypes()

	var result []map[string]interface{}
	values := make([][]byte, len(columns))
	scans := make([]interface{}, len(columns))

	for i, _ := range columns {
		scans[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scans...)
		if err != nil {
			return nil, err
		}
		row := make(map[string]interface{})
		for i, value := range values {
			columnType := columnTypes[i]
			typeName := columnType.DatabaseTypeName()

			switch typeName {
			case "TINYINT", "SMALLINT", "INT", "INTEGER", "BIGINT", "UNSIGNED BIGINT", "UNSIGNED TINYINT":
				if value == nil {
					row[columnType.Name()] = value
					break
				}
				val, err := strconv.Atoi(string(value))
				if err != nil {
					return nil, err
				}
				row[columnType.Name()] = val
				break
			case "TEXT", "VARCHAR", "DATETIME", "BIT":
				row[columnType.Name()] = string(value)
				break
			default:
				slog.Error("dont support type", "type name", typeName)
				break
			}
		}
		result = append(result, row)

	}

	query := &Query{
		Rows:    result,
		Columns: columns,
		Total:   len(result),
	}

	return query, err
}

func (c MysqlConnector) Count(db *sqlx.DB, sqlStr string) (int, error) {
	rows, err := db.Queryx(sqlStr)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = rows.Close()
		_ = db.Close()
	}()
	for rows.Next() {
		var total int
		err = rows.Scan(&total)
		if err != nil {
			return 0, err
		}
		return total, err
	}
	return 0, err
}

func (c MysqlConnector) Ddl(db *sqlx.DB, table string) (string, error) {
	sqlStr := "SHOW CREATE TABLE " + table
	row := db.QueryRowx(sqlStr)
	var tableName string
	var ddl string
	err := row.Scan(&tableName, &ddl)
	return ddl, err
}

func (c MysqlConnector) dsn(database *model.ConnectionConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", *database.Username, *database.Password, *database.Host, *database.Port)
}

func (c MysqlConnector) getColumnType(key string, columnTypes []*sql.ColumnType) *sql.ColumnType {
	for _, columnType := range columnTypes {
		if columnType.Name() == key {
			return columnType
		}
	}
	return nil
}
