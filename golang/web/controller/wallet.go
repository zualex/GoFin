package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zualex/gofin/pkg/config"
	"github.com/zualex/gofin/pkg/wallet"
	"github.com/zualex/gofin/web"
)

func (controller *Controller) ShowWallet(c *gin.Context) {
	c.HTML(http.StatusOK, "wallet.tmpl", web.GetCommonVars("Кошельки", c.Request.URL.Path))
}

func (controller *Controller) CreateWallet(c *gin.Context) {
	vars := web.GetCommonVars("Создание кошелька", c.Request.URL.Path)
	vars["currencies"] = config.Currencies

	c.HTML(http.StatusOK, "wallet_create.tmpl", vars)
}

func (controller *Controller) SaveWallet(c *gin.Context) {
	var wallet wallet.Wallet
	if err := c.ShouldBind(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO в репозиторий
	sql := `INSERT INTO "wallets" ("name", "currency") VALUES ($1, $2);`
	_, err := controller.Database.Exec(sql, wallet.Name, wallet.Currency)
	if err != nil {
		panic(err)
	}

	c.Redirect(http.StatusFound, config.Routes["wallets"])
}
