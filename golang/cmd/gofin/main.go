package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zualex/gofin/pkg/config"
	"github.com/zualex/gofin/pkg/db"
	"github.com/zualex/gofin/web"
	"github.com/zualex/gofin/web/controller"
)

func mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.tmpl", web.GetCommonVars("Главная", c.Request.URL.Path))
}

func main() {
	dbConn, err := db.NewDbConn()
	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()

	router := gin.Default()
	router.Static("/dist", "../frontend/dist")
	router.Static("/plugins", "../frontend/plugins")
	router.LoadHTMLGlob("../frontend/pages/**/*.tmpl")

	controller := controller.Controller{
		Database: dbConn,
	}

	router.GET(config.Routes["main"], mainPage)
	router.GET(config.Routes["wallets"], controller.ShowWallet)
	router.GET(config.Routes["wallet.create"], controller.CreateWallet)
	router.POST(config.Routes["wallet.save"], controller.SaveWallet)
	router.GET(config.Routes["categories"], mainPage)

	router.Run(":8080")
}
