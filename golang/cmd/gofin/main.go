package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zualex/gofin/web"
	"github.com/zualex/gofin/web/wallet"
)

func mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.tmpl", web.GetCommonVars("Главная", c.Request.URL.Path))
}

func main() {
	router := gin.Default()
	router.Static("/dist", "../frontend/dist")
	router.Static("/plugins", "../frontend/plugins")
	router.LoadHTMLGlob("../frontend/pages/**/*.tmpl")

	router.GET(web.Routes["main"], mainPage)
	router.GET(web.Routes["wallets"], wallet.Show)
	router.GET(web.Routes["wallet.create"], wallet.Create)
	router.GET(web.Routes["wallet.save"], wallet.Save)
	router.GET(web.Routes["categories"], mainPage)

	router.Run(":8080")
}
