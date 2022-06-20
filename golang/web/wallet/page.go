package wallet

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zualex/gofin/config"
	"github.com/zualex/gofin/web"
)

func Show(c *gin.Context) {
	c.HTML(http.StatusOK, "wallet.tmpl", web.GetCommonVars("Кошельки", c.Request.URL.Path))
}

func Create(c *gin.Context) {
	vars := web.GetCommonVars("Создание кошелька", c.Request.URL.Path)
	vars["currencies"] = config.Currencies

	c.HTML(http.StatusOK, "wallet_create.tmpl", vars)
}

func Save(c *gin.Context) {
	// TODO Save and redirect
	c.HTML(http.StatusOK, "wallet.tmpl", web.GetCommonVars("Кошельки", c.Request.URL.Path))
}
