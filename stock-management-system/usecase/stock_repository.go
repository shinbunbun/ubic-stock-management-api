package usecase

import "github.com/Yuto/ubic-stock-management-api/stock-management-system/domain"

type StockRepository interface {
	FindAll() ([]domain.Stock, error)
	FindByID(string) (domain.Stock, error)
	ChangeAmount(string, int) error
	Delete(string) error
	Create(string, string, string, int) (string, error)
	FindLike(string) ([]domain.Stock, error)
}

func (it *Interactor) FindStockAll() ([]domain.Stock, error) {
	return it.StockRepository.FindAll()
}

func (it *Interactor) ChangeStockAmount(id string, add int) error {
	return it.StockRepository.ChangeAmount(id, add)
}

func (it *Interactor) CreateStock(image, makerName, productName string, amount int) (string, error) {
	return it.StockRepository.Create(image, makerName, productName, amount)
}

func (it *Interactor) DeleteStock(id string) error {
	return it.StockRepository.Delete(id)
}
