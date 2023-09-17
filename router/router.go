package router

import (
	"dbland/api"
	"dbland/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net"
	"strings"
	"time"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.Conf.Server.Environment)
	r := gin.New()
	r.Use(Logger())
	r.Static("/ui", "./static")
	cookieStore := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("sid", cookieStore))
	r.Use(corsMiddleware())
	g := r.Group(config.Conf.Server.Path)
	{
		query := api.ConnectorHandler{}
		query.InitRouter(g)

		user := api.UserHandler{}
		user.InitRouter(g)

		connectionConfig := api.ConnectionConfigHandler{}
		connectionConfig.InitRouter(g)
	}
	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		elapsedMilliseconds := latency.Milliseconds()
		status := c.Writer.Status()
		url := c.Request.RequestURI
		method := c.Request.Method
		ip := getClientIP(c)
		slog.Info("Request", "code", status, "times", elapsedMilliseconds, "method", method, "url", url, "ip", ip)
	}
}

func getClientIP(c *gin.Context) string {
	if xForwardedFor := c.Request.Header.Get("X-Forwarded-For"); xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		for _, ip := range ips {
			if parsedIP := net.ParseIP(strings.TrimSpace(ip)); parsedIP != nil {
				return parsedIP.String()
			}
		}
	}

	remoteIP, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	return remoteIP
}
