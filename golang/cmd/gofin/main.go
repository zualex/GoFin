package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zualex/gofin/web"
)

func mainPage(c *gin.Context) {

	c.HTML(http.StatusOK, "main.tmpl", gin.H{
		"title":   "Main website",
		"body":    "Текст",
		"sidebar": web.GetSidebar(c.Request.URL.Path),
		"url":     c.Request.URL.Path,
	})
}

func main() {
	router := gin.Default()
	router.Static("/dist", "../frontend/dist")
	router.Static("/plugins", "../frontend/plugins")
	router.LoadHTMLGlob("../frontend/pages/**/*.tmpl")

	router.GET("/", mainPage)
	router.GET("/wallets/", mainPage)
	router.GET("/categories/", mainPage)

	router.Run(":8080")
}
