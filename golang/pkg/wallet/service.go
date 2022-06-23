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

func (service *Service) FindById(id int) (Wallet, error) {
	return service.repository.FindById(id)
}

func (service *Service) Create(wallet Wallet) {
	service.repository.Create(wallet)
}

func (service *Service) Update(wallet Wallet) {
	service.repository.Update(wallet)
}
