package connectors

import (
	"database/sql"
	"dbland/model"
	"github.com/jmoiron/sqlx"
	"github.com/sijms/go-ora/v2"
	"log/slog"
	"strconv"
)

type OracleConnector struct {
}

func (c OracleConnector) Ping(config *model.ConnectionConfig) error {
	dsn := c.dsn(config)
	db, err := sqlx.Open("oracle", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	return err
}

func (c OracleConnector) Connect(config *model.ConnectionConfig) (*sqlx.DB, error) {
	dsn := c.dsn(config)
	return sqlx.Open("oracle", dsn)
}

func (c OracleConnector) ShowDatabases(db *sqlx.DB) ([]string, error) {
	return nil, nil
}

func (c OracleConnector) Use(db *sqlx.DB, database string) error {
	return nil
}

func (c OracleConnector) ShowTables(db *sqlx.DB, database string) (*[]Table, error) {
	var tables []Table
	sqlStr := "SELECT table_name FROM user_tables"

	rows, err := db.Queryx(sqlStr)
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

func (c OracleConnector) Column(db *sqlx.DB, params ...string) (*[]Column, error) {
	return nil, nil
}

func (c OracleConnector) Query(db *sqlx.DB, sqlStr string) (*Query, error) {
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
			case "TINYINT", "SMALLINT", "INT", "INTEGER", "BIGINT", "UNSIGNED BIGINT", "UNSIGNED TINYINT":
				val, err := strconv.Atoi(string(value))
				if err != nil {
					return nil, err
				}
				row[columnType.Name()] = val
				break
			case "NCHAR", "TimeStampDTY", "CHAR":
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

func (c OracleConnector) Count(db *sqlx.DB, sqlStr string) (int, error) {
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

func (c OracleConnector) Ddl(db *sqlx.DB, table string) (string, error) {
	sqlStr := "SELECT DBMS_METADATA.GET_DDL('TABLE', 'SYS_USER') AS TABLE_DDL FROM DUAL"
	row := db.QueryRowx(sqlStr, table)
	var result string
	err := row.Scan(&result)
	return result, err
}

func (c OracleConnector) dsn(database *model.ConnectionConfig) string {
	port, _ := strconv.Atoi(*database.Port)
	connStr := go_ora.BuildUrl(*database.Host, port, "TRADEVL", *database.Username, *database.Password, nil)
	return connStr
}
