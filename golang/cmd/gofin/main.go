package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zualex/gofin/pkg/config"
	"github.com/zualex/gofin/pkg/db"
	"github.com/zualex/gofin/pkg/wallet"
	"github.com/zualex/gofin/web"
	"github.com/zualex/gofin/web/controller"
)

func mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.tmpl", web.GetCommonVars("Главная", c.Request.URL.Path))
}

func route(pattern string, params ...int) string {
	return config.Routes[pattern].GetUrl(params...)
}

func main() {
	dbConn, err := db.NewDbConn()
	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()

	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"route": route,
	})
	router.Static("/dist", "../frontend/dist")
	router.Static("/plugins", "../frontend/plugins")
	router.LoadHTMLGlob("../frontend/pages/**/*.tmpl")

	controller := controller.Controller{
		WalletService: wallet.NewService(dbConn),
	}

	router.NoRoute(controller.NotFoundPage)
	router.GET(config.Routes["main"].GetPattern(), mainPage)
	router.GET(config.Routes["wallets"].GetPattern(), controller.ShowWallet)
	router.GET(config.Routes["wallet.show_create"].GetPattern(), controller.ShowCreateWallet)
	router.POST(config.Routes["wallet.create"].GetPattern(), controller.CreateWallet)
	router.GET(config.Routes["wallet.show_update"].GetPattern(), controller.ShowUpdateWallet)
	router.POST(config.Routes["wallet.update"].GetPattern(), controller.UpdateWallet)
	router.GET(config.Routes["categories"].GetPattern(), mainPage)

	router.Run(":8080")
}
