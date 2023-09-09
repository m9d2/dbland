package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(c *gin.Context, data interface{}) {
	res := &Result{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    data,
	}
	c.JSON(http.StatusOK, res)
}

func Fail(c *gin.Context, err error) {
	res := &Result{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
	c.JSON(http.StatusOK, res)
}

func Unauthorized(c *gin.Context) {
	res := &Result{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}
	c.JSON(http.StatusOK, res)
}

func JSON(c *gin.Context, data interface{}) {
	if err, ok := data.(error); ok {
		Fail(c, err)
	} else {
		Ok(c, data)
	}
}
