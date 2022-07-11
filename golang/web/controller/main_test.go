package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/zualex/gofin/internal/router"
	"github.com/zualex/gofin/pkg/config"
	"github.com/zualex/gofin/pkg/db"
	"github.com/zualex/gofin/pkg/wallet"
	"github.com/zualex/gofin/web"
)

var Ctrl Controller

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
	dbConn, err := db.OpenTestConn()
	if err != nil {
		return -1, fmt.Errorf("could not connect to test database: %w", err)
	}

	defer func() {
		db.TruncateTables(dbConn, []string{"wallets"})
		dbConn.Close()
	}()

	cfg := config.New(config.EnvTest, dbConn)
	web := web.New(cfg)
	Ctrl = Controller{
		Config:        cfg,
		Web:           web,
		WalletService: wallet.NewService(cfg.Db),
	}

	return m.Run(), nil
}

func TestShowMain(t *testing.T) {
	mockResponse := `<h1 class="m-0">Главная</h1>`

	r := router.GetRouter(Ctrl.Config)
	r.GET("/", Ctrl.ShowMain)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.NotEqual(t, -1, strings.Index(string(responseData), mockResponse))
	assert.Equal(t, http.StatusOK, w.Code)
}
