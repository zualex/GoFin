package config

import "strings"

type Route struct {
	Pattern string
}

var Routes = map[string]Route{
	"main":          {"/"},
	"wallets":       {"/wallets/"},
	"wallet.create": {"/wallets/create/"},
	"wallet.save":   {"/wallets/save/"},
	"wallet.update": {"/wallets/update/:id"},
	"categories":    {"/categories/"},
}

func (route Route) GetPattern() string {
	return route.Pattern
}

func (route Route) GetUrl(params ...string) string {
	if len(params) > 0 {
		return strings.Replace(route.Pattern, ":id", params[0], -1) //TODO сделать более умную подстановку
	}

	return route.Pattern
}
