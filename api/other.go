package api

import (
	"dbland/service"
	"github.com/gin-gonic/gin"
)

type OtherHandler struct {
	otherService service.OtherService
}

func (h OtherHandler) InitRouter(g *gin.RouterGroup) {
	g.GET("readme", h.loadReadme)
}

func (h OtherHandler) loadReadme(c *gin.Context) {
	content, err := h.otherService.LoadReadme()
	if err != nil {
		JSON(c, err)
	} else {
		JSON(c, content)
	}
}
