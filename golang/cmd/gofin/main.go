package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

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

func route(pattern, id string) string {
	return strings.Replace(config.Routes[pattern].GetPattern(), ":id", id, -1) //TODO сделать более умную подстановку
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

	router.GET(config.Routes["main"].GetPattern(), mainPage)
	router.GET(config.Routes["wallets"].GetPattern(), controller.ShowWallet)
	router.GET(config.Routes["wallet.create"].GetPattern(), controller.CreateWallet)
	router.GET(config.Routes["wallet.update"].GetPattern(), controller.UpdateWallet)
	router.POST(config.Routes["wallet.save"].GetPattern(), controller.SaveWallet)
	router.GET(config.Routes["categories"].GetPattern(), mainPage)

	router.Run(":8080")
}
