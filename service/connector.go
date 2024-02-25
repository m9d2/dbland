package service

import (
	"dbland/connectors"
	"dbland/model"
	"dbland/repository"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type ConnectorService struct {
	connectorFactory           connectors.ConnectorFactory
	connectionConfigRepository repository.ConnectionConfigRepository
	executionLogRepository     repository.ExecutionLogRepository
}

type QueryReq struct {
	Cid      uint   `json:"cid" binding:"required"`
	Sql      string `json:"sql"`
	Database string `json:"db"`
	Table    string `json:"table"`
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	Sort     Sort   `json:"sort"`
}

type Sort struct {
	Prop  string `json:"prop"`
	Order string `json:"order"`
}

func (s ConnectorService) Ping(c *gin.Context) error {
	var config model.Config
	if err := c.ShouldBindJSON(&config); err != nil {
		return err
	}

	connector, err := s.connectorFactory.GetInstance(*config.Type)
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

	db, connector, err := s.getConnector(&req)
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

	db, connector, err := s.getConnector(&req)
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

	db, connector, err := s.getConnector(&req)
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

	db, connector, err := s.getConnector(&req)
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

	db, connector, err := s.getConnector(&req)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = db.Close()
	}()

	err = connector.Use(db, req.Database)

	startTime := time.Now()
	// exec sql
	var result *connectors.Query

	var sqlStr = req.Sql
	if req.Sort.Prop != "" && req.Sort.Order != "" && !containsKeywords(strings.ToLower(sqlStr), "order by", "limit", "offset") {
		sqlStr = sqlStr + fmt.Sprintf(" ORDER BY %v %v", req.Sort.Prop, req.Sort.Order)
	}
	if req.Page != 0 && req.Size != 0 && !containsKeywords(strings.ToLower(sqlStr), "limit", "offset") {
		offset := (req.Page - 1) * req.Size
		sqlStr = sqlStr + fmt.Sprintf(" limit %v OFFSET %v", req.Size, offset)
	}
	slog.Info("Query", "sql", sqlStr)
	result, err = connector.Query(db, sqlStr)
	if err != nil {
		return nil, err
	}

	elapsedTime := time.Since(startTime)
	cost := elapsedTime.Seconds()
	str := fmt.Sprintf("%.3f", cost)
	cost, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return nil, err
	}
	result.ElapsedTime = cost

	var countSql = fmt.Sprintf("select count(1) FROM (%v) as _A", req.Sql)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count, _ := connector.Count(db, countSql)
		result.Total = count
		result.TotalPage = int(math.Ceil(float64(count) / float64(req.Size)))
		wg.Done()
	}()
	wg.Wait()

	// save log
	req.Sql = sqlStr
	println(c.Request.UserAgent())
	go s.saveExecutionLog(&req, cost, c.ClientIP(), c.Request.UserAgent(), err)
	return result, err
}

func (s ConnectorService) Execute(c *gin.Context) (int, error) {
	var req QueryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return 0, err
	}

	slog.Info("query", "sql", req.Sql)
	db, connector, err := s.getConnector(&req)
	if err != nil {
		return 0, err
	}
	return connector.Execute(db, req.Sql)
}

func (s ConnectorService) saveExecutionLog(queryReq *QueryReq, cost float64, ip string, userAgent string, err error) {
	var status int
	if err != nil {
		status = model.ExecutionLogFail
	} else {
		status = model.ExecutionLogSuccess
	}

	var config *model.Config
	config = s.connectionConfigRepository.GetById(queryReq.Cid)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	log := &model.Log{
		Source:      *config.Name,
		Status:      status,
		Database:    queryReq.Database,
		Sql:         queryReq.Sql,
		Cost:        cost,
		Ip:          ip,
		UserAgent:   userAgent,
		CreatedTime: time.Now(),
	}
	s.executionLogRepository.Save(log)
}

func (s ConnectorService) getConnector(req *QueryReq) (*sqlx.DB, connectors.Connector, error) {
	config := s.connectionConfigRepository.GetById(req.Cid)

	connector, err := s.connectorFactory.GetInstance(*config.Type)
	if err != nil {
		return nil, nil, err
	}

	var db *sqlx.DB
	db, err = connector.Connect(config)
	if err != nil {
		return nil, nil, err
	}

	return db, connector, nil
}

func containsKeywords(str string, keywords ...string) bool {
	for _, keyword := range keywords {
		keywordLower := strings.ToLower(keyword)
		index := strings.Index(str, keywordLower)
		if index != -1 {
			return true
		}
	}
	return false
}
