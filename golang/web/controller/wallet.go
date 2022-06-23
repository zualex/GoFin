package controller

import (
	"net/http"
	"strconv"

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

func (controller *Controller) ShowCreateWallet(c *gin.Context) {
	vars := web.GetCommonVars("Создание кошелька", c.Request.URL.Path)
	vars["currencies"] = config.Currencies
	vars["model"] = wallet.Wallet{}

	c.HTML(http.StatusOK, "wallet_create.tmpl", vars)
}

func (controller *Controller) CreateWallet(c *gin.Context) {
	var wallet wallet.Wallet
	if err := c.ShouldBind(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.WalletService.Create(wallet)

	c.Redirect(http.StatusFound, config.Routes["wallets"].GetUrl())
}

func (controller *Controller) ShowUpdateWallet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	wallet, err := controller.WalletService.FindById(id)
	if err != nil {
		controller.NotFoundPage(c)
	}

	vars := web.GetCommonVars("Изменение кошелька "+wallet.Name, c.Request.URL.Path)
	vars["currencies"] = config.Currencies
	vars["model"] = wallet

	c.HTML(http.StatusOK, "wallet_create.tmpl", vars)
}

func (controller *Controller) UpdateWallet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := controller.WalletService.FindById(id)
	if err != nil {
		controller.NotFoundPage(c)
	}

	var wallet wallet.Wallet
	wallet.Id = id

	if err := c.ShouldBind(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.WalletService.Update(wallet)

	c.Redirect(http.StatusFound, config.Routes["wallets"].GetUrl())
}
