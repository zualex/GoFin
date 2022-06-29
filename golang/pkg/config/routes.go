package config

import (
	"strconv"
	"strings"
)

type Route struct {
	Pattern string
}

var routes = map[string]Route{
	"main":               {"/"},
	"wallets":            {"/wallets/"},
	"wallet.show_create": {"/wallets/create/"},
	"wallet.create":      {"/wallets/create/"},
	"wallet.show_update": {"/wallets/update/:id"},
	"wallet.update":      {"/wallets/update/:id"},
	"categories":         {"/categories/"},
}

func (route Route) GetPattern() string {
	return route.Pattern
}

func (route Route) GetUrl(params ...int) string {
	if len(params) > 0 {
		//TODO сделать более умную подстановку
		id := strconv.Itoa(params[0])
		result := strings.Replace(route.Pattern, ":id", id, -1)
		result = strings.Replace(result, "*id", id, -1)

		return result
	}

	return route.Pattern
}
