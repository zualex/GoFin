package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zualex/gofin/pkg/wallet"
)

func (ctrl *Controller) ShowWallet(c *gin.Context) {
	fmt.Println("!!!")
	fmt.Println(c.ContentType())

	vars := ctrl.Web.GetCommonVars("Кошельки", c.Request.URL.Path)
	vars["wallets"] = ctrl.WalletService.List()

	render(c, http.StatusOK, "wallet.tmpl", vars)
}

func (ctrl *Controller) ShowCreateWallet(c *gin.Context) {
	vars := ctrl.Web.GetCommonVars("Создание кошелька", c.Request.URL.Path)
	vars["currencies"] = ctrl.Config.Currencies
	vars["model"] = wallet.Wallet{}

	c.HTML(http.StatusOK, "wallet_create.tmpl", vars)
}

func (ctrl *Controller) CreateWallet(c *gin.Context) {
	var wallet wallet.Wallet
	if err := c.ShouldBind(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctrl.WalletService.Create(wallet)

	c.Redirect(http.StatusFound, ctrl.Config.Routes["wallets"].GetUrl())
}

func (ctrl *Controller) ShowUpdateWallet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	wallet, err := ctrl.WalletService.FindById(id)
	if err != nil {
		ctrl.NotFoundPage(c)
	}

	vars := ctrl.Web.GetCommonVars("Изменение кошелька "+wallet.Name, c.Request.URL.Path)
	vars["currencies"] = ctrl.Config.Currencies
	vars["model"] = wallet

	c.HTML(http.StatusOK, "wallet_create.tmpl", vars)
}

func (ctrl *Controller) UpdateWallet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := ctrl.WalletService.FindById(id)
	if err != nil {
		ctrl.NotFoundPage(c)
	}

	var wallet wallet.Wallet
	wallet.Id = id

	if err := c.ShouldBind(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctrl.WalletService.Update(wallet)

	c.Redirect(http.StatusFound, ctrl.Config.Routes["wallets"].GetUrl())
}
