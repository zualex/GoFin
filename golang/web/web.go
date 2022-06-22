package web

import "github.com/zualex/gofin/pkg/config"

type SidebarItem struct {
	Name     string
	Url      string
	Icon     string
	IsActive bool
}

func GetCommonVars(title, url string) map[string]interface{} {
	return map[string]interface{}{
		"title":   title,
		"sidebar": GetSidebar(url),
		"url":     url,
		"routes":  config.Routes,
	}
}

func GetSidebar(url string) []SidebarItem {
	sidebar := []SidebarItem{
		{"Главная", config.Routes["main"], "fa-tachometer-alt", false},
		{"Кошельки", config.Routes["wallets"], "fa-wallet", false},
		{"Категории", config.Routes["categories"], "fa-folder", false},
	}

	for i, v := range sidebar {
		if url == v.Url {
			sidebar[i].IsActive = true
		}
	}

	return sidebar
}
