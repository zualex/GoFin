package router

import (
	"html/template"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/zualex/gofin/pkg/config"
)

func GetRouter(cfg *config.Config) *gin.Engine {
	if cfg.Env == config.EnvTest {
		// Выключаем stdout
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"route": func(pattern string, params ...int) string {
			return cfg.Routes[pattern].GetUrl(params...)
		},
	})
	router.Static("/dist", "/frontend/dist")
	router.Static("/plugins", "/frontend/plugins")
	router.LoadHTMLGlob("/frontend/pages/**/*.tmpl")

	return router
}
