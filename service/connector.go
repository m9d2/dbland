package service

import (
	"dbland/connectors"
	"dbland/model"
	"dbland/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"strconv"
	"sync"
	"time"
)

type ConnectorService struct {
	connectionConfigRepository repository.ConnectionConfigRepository
	mysqlConnector             connectors.MysqlConnector
	oracleConnector            connectors.OracleConnector
	sqLiteConnector            connectors.SQLiteConnector
	postgreSQLConnector        connectors.PostgreSQLConnector
	executionLogRepository     repository.ExecutionLogRepository
}

func (s ConnectorService) Ping(c *gin.Context) error {
	var config model.ConnectionConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		return err
	}

	connector, err := s.getConnector(*config.Type)
	if err != nil {
		return err
	}

	err = connector.Ping(&config)
	return err
}

func (s ConnectorService) ShowDatabases(c *gin.Context) ([]string, error) {
	var req QueryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	db, connector, err := s.getDBAndConnector(&req)
	if err != nil {
		return nil, err
	}
	return connector.ShowDatabases(db)
}

func (s ConnectorService) ShowTables(c *gin.Context) (interface{}, error) {
	var req QueryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	db, connector, err := s.getDBAndConnector(&req)
	if err != nil {
		return nil, err
	}
	return connector.ShowTables(db, req.Database)
}

func (s ConnectorService) Column(c *gin.Context) (interface{}, error) {
	var req QueryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	slog.Info("columns", "database", req.Database, "table", req.Table)

	db, connector, err := s.getDBAndConnector(&req)
	if err != nil {
		return nil, err
	}
	return connector.Column(db, req.Database, req.Table)
}

func (s ConnectorService) Ddl(c *gin.Context) (interface{}, error) {
	var req QueryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	slog.Info("ddl", "table", req.Table)

	db, connector, err := s.getDBAndConnector(&req)
	if err != nil {
		return nil, err
	}
	return connector.Ddl(db, req.Table)
}

func (s ConnectorService) Query(c *gin.Context) (interface{}, error) {
	var req QueryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	slog.Info("query", "sql", req.Sql)

	db, connector, err := s.getDBAndConnector(&req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = db.Close()
	}()

	startTime := time.Now()
	// exec sql
	var result *connectors.Query

	var sqlStr = ""
	if req.Page == 0 && req.Size == 0 {
		result, err = connector.Query(db, req.Sql)
	} else {
		offset := (req.Page - 1) * req.Size
		sqlStr = req.Sql + fmt.Sprintf(" limit %v OFFSET %v", req.Size, offset)
		result, err = connector.Query(db, sqlStr)
	}

	elapsedTime := time.Since(startTime)
	cost := elapsedTime.Seconds()
	str := fmt.Sprintf("%.3f", cost)
	cost, _ = strconv.ParseFloat(str, 64)
	result.ElapsedTime = cost

	var countSql = fmt.Sprintf("select count(1) FROM (%v) as _A", req.Sql)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count, _ := connector.Count(db, countSql)
		result.Total = count
		wg.Done()
	}()
	wg.Wait()

	// save log
	req.Sql = sqlStr
	go s.saveExecutionLog(&req, cost, err)
	return result, err
}

func (s ConnectorService) saveExecutionLog(queryReq *QueryReq, cost float64, err error) {
	var status int
	if err != nil {
		status = model.ExecutionLogFail
	} else {
		status = model.ExecutionLogSuccess
	}

	var config *model.ConnectionConfig
	config, err = s.connectionConfigRepository.GetById(queryReq.Cid)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	tx := repository.DB.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	log := &model.ExecutionLog{
		Source:      *config.Name,
		Status:      status,
		Database:    queryReq.Database,
		Sql:         queryReq.Sql,
		Cost:        cost,
		CreatedTime: time.Now(),
	}
	s.executionLogRepository.Save(tx, log)
}

type QueryReq struct {
	Cid      uint   `json:"cid" binding:"required"`
	Sql      string `json:"sql"`
	Database string `json:"db"`
	Table    string `json:"table"`
	Page     int    `json:"page"`
	Size     int    `json:"size"`
}

func (s ConnectorService) getDBAndConnector(req *QueryReq) (*sqlx.DB, connectors.Connector, error) {
	config, err := s.connectionConfigRepository.GetById(req.Cid)
	if err != nil {
		return nil, nil, err
	}

	var connector connectors.Connector
	connector, err = s.getConnector(*config.Type)
	if err != nil {
		return nil, nil, err
	}

	var db *sqlx.DB
	db, err = connector.Connect(config)
	if err != nil {
		return nil, nil, err
	}

	if req.Database != "" {
		err = connector.Use(db, req.Database)
		if err != nil {
			return nil, nil, err
		}
	}

	return db, connector, nil
}

func (s ConnectorService) getConnector(dbType string) (connectors.Connector, error) {
	var connector connectors.Connector
	switch dbType {
	case connectors.Mysql:
		connector = s.mysqlConnector
	case connectors.Oracle:
		connector = s.oracleConnector
	case connectors.SQLite:
		connector = s.sqLiteConnector
	case connectors.PostgreSQL:
		connector = s.postgreSQLConnector

	default:
		err := fmt.Errorf("database not support")
		return nil, err
	}
	return connector, nil
}
