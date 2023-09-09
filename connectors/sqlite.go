package connectors

import (
	"database/sql"
	"dbland/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"strconv"
)

type SQLiteConnector struct {
}

func (c SQLiteConnector) Ping(config *model.ConnectionConfig) error {
	db, err := sqlx.Open("sqlite3", c.dsn(config))
	if err != nil {
		return err
	}
	return db.Ping()
}

func (c SQLiteConnector) Connect(config *model.ConnectionConfig) (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", c.dsn(config))
}

func (c SQLiteConnector) ShowDatabases(db *sqlx.DB) ([]string, error) {
	result := []string{"main"}
	return result, nil
}

func (c SQLiteConnector) Use(db *sqlx.DB, database string) error {
	return nil
}

func (c SQLiteConnector) ShowTables(db *sqlx.DB, database string) (*[]Table, error) {
	var tables []Table
	sqlStr := "SELECT name FROM sqlite_master WHERE type = 'table' AND name != 'sqlite_sequence'"
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
		var err = rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}

		if tableName.Valid {
			table.Name = tableName.String
		}

		tables = append(tables, table)
	}
	return &tables, err
}

func (c SQLiteConnector) Column(db *sqlx.DB, params ...string) (*[]Column, error) {
	var columns []Column
	table := params[0]
	sqlStr := fmt.Sprintf("PRAGMA table_info(%v)", table)

	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = db.Close()
	}()

	for rows.Next() {
		var column Column
		var id sql.NullString
		var columnName sql.NullString
		var columnType sql.NullString
		var notnull sql.NullString
		var dfltValue sql.NullString
		var pk sql.NullString
		err = rows.Scan(&id, &columnName, &columnType, &notnull, &dfltValue, &pk)
		if err != nil {
			return nil, err
		}

		if columnName.Valid {
			column.ColumnName = columnName.String
		}
		if columnType.Valid {
			column.ColumnType = columnType.String
		}
		columns = append(columns, column)
	}
	return &columns, err
}

func (c SQLiteConnector) Query(db *sqlx.DB, sqlStr string) (*Query, error) {
	rows, err := db.Queryx(sqlStr)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
		_ = db.Close()
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
			case "INTEGER":
				val, err := strconv.ParseFloat(string(value), 64)
				if err != nil {
					return nil, err
				}
				row[columnType.Name()] = val
				break
			case "TEXT", "DATE":
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

func (c SQLiteConnector) Count(db *sqlx.DB, sqlStr string) (int, error) {
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
		err = rows.Scan(total)
		if err != nil {
			return total, nil
		}
	}
	return 0, err
}

func (c SQLiteConnector) Ddl(db *sqlx.DB, table string) (string, error) {
	return "", nil
}

func (c SQLiteConnector) dsn(config *model.ConnectionConfig) string {
	return fmt.Sprintf(*config.DbFile)
}
