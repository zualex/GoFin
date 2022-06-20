package web

type Sidebar struct {
	Name     string
	Url      string
	Icon     string
	IsActive bool
}

func GetSidebar(url string) []Sidebar {
	sidebar := []Sidebar{
		{"Главная", "/", "fa-tachometer-alt", false},
		{"Кошелеки", "/wallets/", "fa-wallet", false},
		{"Категории", "/categories/", "fa-folder", false},
	}

	for i, v := range sidebar {
		if url == v.Url {
			sidebar[i].IsActive = true
		}
	}

	return sidebar
}
