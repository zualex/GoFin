package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/zualex/gofin/internal/router"
)

// TODO добавить фикстуры

func TestShowWallet(t *testing.T) {
	mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`

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

	assert.Equal(t, mockResponse, data["wallets"])
	assert.Equal(t, http.StatusOK, w.Code)
}
