package api

import (
	"dbland/service"
	"github.com/gin-gonic/gin"
)

type ConnectorHandler struct {
	service service.ConnectorService
}

func (h ConnectorHandler) InitRouter(g *gin.RouterGroup) {
	g.POST("ping", h.ping)
	g.POST("databases", h.showDatabases)
	g.POST("tables", h.showTables)
	g.POST("columns", h.column)
	g.POST("ddl", h.ddl)
	g.POST("query", h.query)
	g.POST("execute", h.execute)
}

func (h ConnectorHandler) ping(c *gin.Context) {
	err := h.service.Ping(c)
	JSON(c, err)
}

func (h ConnectorHandler) showDatabases(c *gin.Context) {
	data, err := h.service.ShowDatabases(c)
	if err != nil {
		JSON(c, err)
	} else {
		JSON(c, data)
	}
}

func (h ConnectorHandler) showTables(c *gin.Context) {
	data, err := h.service.ShowTables(c)
	if err != nil {
		JSON(c, err)
	} else {
		JSON(c, data)
	}
}

func (h ConnectorHandler) column(c *gin.Context) {
	data, err := h.service.Column(c)
	if err != nil {
		JSON(c, err)
	} else {
		JSON(c, data)
	}
}

func (h ConnectorHandler) ddl(c *gin.Context) {
	data, err := h.service.Ddl(c)
	if err != nil {
		JSON(c, err)
	} else {
		JSON(c, data)
	}
}

func (h ConnectorHandler) query(c *gin.Context) {
	data, err := h.service.Query(c)
	if err != nil {
		JSON(c, err)
	} else {
		JSON(c, data)
	}
}

func (h ConnectorHandler) execute(c *gin.Context) {
	data, err := h.service.Execute(c)
	if err != nil {
		JSON(c, err)
	} else {
		JSON(c, data)
	}
}
