package web

import (
	"github.com/zualex/gofin/pkg/config"
)

type Web struct {
	Config *config.Config
}

func New(config *config.Config) *Web {
	return &Web{config}
}

type SidebarItem struct {
	Name     string
	Url      string
	Icon     string
	IsActive bool
}

func (web *Web) GetCommonVars(title, url string) map[string]interface{} {
	return map[string]interface{}{
		"title":   title,
		"sidebar": web.GetSidebar(url),
		"url":     url,
	}
}

func (web *Web) GetSidebar(url string) []SidebarItem {
	sidebar := []SidebarItem{
		{"Главная", web.Config.Routes["main"].GetUrl(), "fa-tachometer-alt", false},
		{"Кошельки", web.Config.Routes["wallets"].GetUrl(), "fa-wallet", false},
		{"Категории", web.Config.Routes["categories"].GetUrl(), "fa-folder", false},
	}

	for i, v := range sidebar {
		if url == v.Url {
			sidebar[i].IsActive = true
		}
	}

	return sidebar
}
