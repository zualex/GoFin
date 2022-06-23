package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zualex/gofin/web"
)

func (controller *Controller) NotFoundPage(c *gin.Context) {
	c.HTML(404, "404.tmpl", web.GetCommonVars("404 страница не найдена", c.Request.URL.Path))
}
