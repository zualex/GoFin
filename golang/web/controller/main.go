package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) ShowMain(c *gin.Context) {
	c.HTML(http.StatusOK, "main.tmpl", ctrl.Web.GetCommonVars("Главная", c.Request.URL.Path))
}
