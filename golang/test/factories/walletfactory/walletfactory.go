package walletfactory

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"github.com/zualex/gofin/pkg/config"
	"github.com/zualex/gofin/pkg/wallet"
)

type FactoryConfig struct {
	Config  *config.Config
	Service *wallet.Service
}

func NewFactory(cfg *config.Config, service *wallet.Service) *FactoryConfig {
	return &FactoryConfig{cfg, service}
}

func (factoryCfg *FactoryConfig) GetFactory() *factory.Factory {
	return factory.NewFactory(
		&wallet.Wallet{},
	).SeqInt("Id", func(n int) (interface{}, error) {
		return n, nil
	}).Attr("Name", func(args factory.Args) (interface{}, error) {
		return randomdata.FullName(randomdata.RandomGender), nil
	}).Attr("Currency", func(args factory.Args) (interface{}, error) {
		countCurrency := len(factoryCfg.Config.Currencies)
		indexCurrency := randomdata.Number(countCurrency)

		return factoryCfg.Config.Currencies[indexCurrency], nil
	})
}

func (factoryCfg *FactoryConfig) Create(count int) []wallet.Wallet {
	f := factoryCfg.GetFactory()
	i := 0
	for i < count {
		w := f.MustCreate().(*wallet.Wallet)
		factoryCfg.Service.Create(*w)

		i += 1
	}

	return factoryCfg.Service.List()
}
