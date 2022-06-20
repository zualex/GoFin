package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.tmpl", gin.H{
		"title": "Main website",
		"body":  "Текст",
	})
}

func main() {
	router := gin.Default()
	router.Static("/dist", "../frontend/dist")
	router.Static("/plugins", "../frontend/plugins")
	router.LoadHTMLGlob("../frontend/pages/**/*.tmpl")

	router.GET("/", mainPage)

	router.Run(":8080")
}
