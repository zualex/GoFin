package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/zualex/gofin/internal/router"
	"github.com/zualex/gofin/test/factories/walletfactory"
)

func TestShowWallet(t *testing.T) {
	mockResponse := walletfactory.NewFactory(Ctrl.Config, Ctrl.WalletService).Create(3)

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

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, len(mockResponse), len(responseWallets))
	i := 0
	for i < len(mockResponse) {
		assert.Equal(t, mockResponse[i].Name, responseWallets[i].(map[string]interface{})["Name"])
		assert.Equal(t, mockResponse[i].Currency, responseWallets[i].(map[string]interface{})["Currency"])
		i += 1
	}
}
