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
	}
}

func GetSidebar(url string) []SidebarItem {
	sidebar := []SidebarItem{
		{"Главная", config.Routes["main"].GetUrl(), "fa-tachometer-alt", false},
		{"Кошельки", config.Routes["wallets"].GetUrl(), "fa-wallet", false},
		{"Категории", config.Routes["categories"].GetUrl(), "fa-folder", false},
	}

	for i, v := range sidebar {
		if url == v.Url {
			sidebar[i].IsActive = true
		}
	}

	return sidebar
}
