package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zualex/gofin/pkg/config"
	"github.com/zualex/gofin/pkg/wallet"
	"github.com/zualex/gofin/web"
)

func (controller *Controller) ShowWallet(c *gin.Context) {
	vars := web.GetCommonVars("Кошельки", c.Request.URL.Path)
	vars["wallets"] = controller.WalletService.List()

	c.HTML(http.StatusOK, "wallet.tmpl", vars)
}

func (controller *Controller) CreateWallet(c *gin.Context) {
	vars := web.GetCommonVars("Создание кошелька", c.Request.URL.Path)
	vars["currencies"] = config.Currencies

	c.HTML(http.StatusOK, "wallet_create.tmpl", vars)
}

func (controller *Controller) UpdateWallet(c *gin.Context) {
	id := c.Param("id")
	vars := web.GetCommonVars("Изменение кошелька "+id, c.Request.URL.Path)
	vars["currencies"] = config.Currencies

	c.HTML(http.StatusOK, "wallet_create.tmpl", vars)
}

func (controller *Controller) SaveWallet(c *gin.Context) {
	var wallet wallet.Wallet
	if err := c.ShouldBind(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.WalletService.Save(wallet)

	c.Redirect(http.StatusFound, config.Routes["wallets"].GetUrl())
}
