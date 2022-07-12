package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, code int, name string, obj interface{}) {
	if strings.ToLower(c.ContentType()) == "application/json" {
		c.JSON(http.StatusOK, obj)
	} else {
		c.HTML(http.StatusOK, "wallet.tmpl", obj)
	}
}

func (ctrl *Controller) ShowMain(c *gin.Context) {
	c.HTML(http.StatusOK, "main.tmpl", ctrl.Web.GetCommonVars("Главная", c.Request.URL.Path))
}
