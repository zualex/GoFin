package main

import (
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zualex/gofin/pkg/config"
	"github.com/zualex/gofin/pkg/db"
	"github.com/zualex/gofin/pkg/wallet"
	"github.com/zualex/gofin/web"
	"github.com/zualex/gofin/web/controller"
)

var cfg *config.Config

func route(pattern string, params ...int) string {
	return cfg.Routes[pattern].GetUrl(params...)
}

func main() {
	dbConn, err := db.NewDbConn()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	cfg = config.New(os.Getenv("APP_ENV"), dbConn)
	web := web.New(cfg)

	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"route": route,
	})
	router.Static("/dist", "../frontend/dist")
	router.Static("/plugins", "../frontend/plugins")
	router.LoadHTMLGlob("../frontend/pages/**/*.tmpl")

	controller := controller.Controller{
		Config:        cfg,
		Web:           web,
		WalletService: wallet.NewService(cfg.Db),
	}

	router.NoRoute(controller.NotFoundPage)
	router.GET(cfg.Routes["main"].GetPattern(), controller.ShowMain)
	router.GET(cfg.Routes["wallets"].GetPattern(), controller.ShowWallet)
	router.GET(cfg.Routes["wallet.show_create"].GetPattern(), controller.ShowCreateWallet)
	router.POST(cfg.Routes["wallet.create"].GetPattern(), controller.CreateWallet)
	router.GET(cfg.Routes["wallet.show_update"].GetPattern(), controller.ShowUpdateWallet)
	router.POST(cfg.Routes["wallet.update"].GetPattern(), controller.UpdateWallet)
	router.GET(cfg.Routes["categories"].GetPattern(), controller.ShowMain)

	router.Run(":8080")
}
