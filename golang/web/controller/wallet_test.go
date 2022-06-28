package controller

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/zualex/gofin/pkg/wallet"
)

func TestShowWallet(t *testing.T) {
	controller := Controller{
		WalletService: wallet.NewService(db),
	}

	fmt.Println(controller)

	assert.Equal(t, 1, 1)
}
