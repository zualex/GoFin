package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/zualex/gofin/internal/router"
	"github.com/zualex/gofin/pkg/wallet"
)

// TODO добавить фикстуры

func TestShowWallet(t *testing.T) {
	testWallet := wallet.Wallet{Name: "test1", Currency: Ctrl.Config.Currencies[0]}
	Ctrl.WalletService.Create(testWallet)
	mockResponse := []wallet.Wallet{testWallet}

	r := router.GetRouter(Ctrl.Config)
	r.GET("/wallets/", Ctrl.ShowWallet)
	req, _ := http.NewRequest("GET", "/wallets/", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	var data map[string]interface{}
	err := json.Unmarshal(responseData, &data)

	if err != nil {
		panic(err)
	}

	responseWallets := data["wallets"].([]interface{})

	assert.Equal(t, mockResponse[0].Name, responseWallets[0].(map[string]interface{})["Name"])
	assert.Equal(t, mockResponse[0].Currency, responseWallets[0].(map[string]interface{})["Currency"])
	assert.Equal(t, http.StatusOK, w.Code)
}
