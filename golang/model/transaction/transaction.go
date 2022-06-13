package transaction

import (
	"time"

	"github.com/zualex/gofin/model/category"
	"github.com/zualex/gofin/model/wallet"
)

type Transaction struct {
	Amount   float64
	Category category.Category
	Wallet   wallet.Wallet
	Note     string
	Type     string
	Datetime time.Time
}
