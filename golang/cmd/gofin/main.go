package main

import (
	"os"

	"github.com/zualex/gofin/internal/app"
)

func main() {
	app.Run(os.Getenv("APP_ENV"))
	// app.Run("test")
}
