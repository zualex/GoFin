package wallet

import "database/sql"

type Service struct {
	repository Repository
}

func NewService(db *sql.DB) *Service {
	return &Service{
		Repository{db},
	}
}

func (service *Service) List() []Wallet {
	return service.repository.List()
}

func (service *Service) Save(wallet Wallet) {
	service.repository.Save(wallet)
}
