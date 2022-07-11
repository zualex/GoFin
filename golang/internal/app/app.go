package app

import (
	"fmt"
	"log"

	"github.com/zualex/gofin/internal/router"
	"github.com/zualex/gofin/pkg/config"
	"github.com/zualex/gofin/pkg/db"
	"github.com/zualex/gofin/pkg/wallet"
	"github.com/zualex/gofin/web"
	"github.com/zualex/gofin/web/controller"
)

type App struct {
	Config *config.Config
}

func Run(env string) *App {
	dbConn, err := db.NewDbConn(env)
	if err != nil {
		log.Fatal(err)
	}

	if env != config.EnvTest {
		defer dbConn.Close()
	} else {
		defer func() {
			for _, t := range []string{"wallets"} {
				_, _ = dbConn.Exec(fmt.Sprintf("DELETE FROM %s", t))
			}

			dbConn.Close()
		}()
	}

	cfg := config.New(env, dbConn)
	web := web.New(cfg)

	controller := controller.Controller{
		Config:        cfg,
		Web:           web,
		WalletService: wallet.NewService(cfg.Db),
	}

	r := router.GetRouter(cfg)

	r.NoRoute(controller.NotFoundPage)
	r.GET(cfg.Routes["main"].GetPattern(), controller.ShowMain)
	r.GET(cfg.Routes["wallets"].GetPattern(), controller.ShowWallet)
	r.GET(cfg.Routes["wallet.show_create"].GetPattern(), controller.ShowCreateWallet)
	r.POST(cfg.Routes["wallet.create"].GetPattern(), controller.CreateWallet)
	r.GET(cfg.Routes["wallet.show_update"].GetPattern(), controller.ShowUpdateWallet)
	r.POST(cfg.Routes["wallet.update"].GetPattern(), controller.UpdateWallet)
	r.GET(cfg.Routes["categories"].GetPattern(), controller.ShowMain)

	r.Run(":8080")

	return &App{cfg}
}
