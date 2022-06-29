package controller

import (
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) NotFoundPage(c *gin.Context) {
	c.HTML(404, "404.tmpl", ctrl.Web.GetCommonVars("404 страница не найдена", c.Request.URL.Path))
}
