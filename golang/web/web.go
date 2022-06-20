package web

var Routes = map[string]string{
	"main":          "/",
	"wallets":       "/wallets/",
	"wallet.create": "/wallets/create/",
	"wallet.save":   "/wallets/save/",
	"categories":    "/categories/",
}

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
		"routes":  Routes,
	}
}

func GetSidebar(url string) []SidebarItem {
	sidebar := []SidebarItem{
		{"Главная", Routes["main"], "fa-tachometer-alt", false},
		{"Кошельки", Routes["wallets"], "fa-wallet", false},
		{"Категории", Routes["categories"], "fa-folder", false},
	}

	for i, v := range sidebar {
		if url == v.Url {
			sidebar[i].IsActive = true
		}
	}

	return sidebar
}
