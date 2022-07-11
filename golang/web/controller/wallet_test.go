package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/zualex/gofin/internal/router"
)

// TODO добавить фикстуры

func TestShowWallet(t *testing.T) {
	// mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`

	r := router.GetRouter(Ctrl.Config)
	r.GET("/wallets/", Ctrl.ShowMain)
	req, _ := http.NewRequest("GET", "/wallets/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// responseData, _ := ioutil.ReadAll(w.Body)
	// assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
