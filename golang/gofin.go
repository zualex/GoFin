package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var dirTemplates = "template"

func mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title": "Main website",
		"body":  "Текст",
	})
}

func main() {
	router := gin.Default()
	router.Static("/dist", "../frontend/dist")
	router.Static("/plugins", "../frontend/plugins")
	router.LoadHTMLGlob(dirTemplates + "/*.html")
	router.GET("/", mainPage)
	// router.Run("localhost:8080")
	router.Run("0.0.0.0:8080")
}
