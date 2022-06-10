package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getMain)

	router.Run("localhost:8080")
}

func getMain(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, []string{"test", "test2"})
}
