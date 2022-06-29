package controller

import (
	"github.com/zualex/gofin/pkg/config"
	"github.com/zualex/gofin/pkg/wallet"
	"github.com/zualex/gofin/web"
)

type Controller struct {
	Config        *config.Config
	Web           *web.Web
	WalletService *wallet.Service
}
