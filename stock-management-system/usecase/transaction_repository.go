package usecase

import "github.com/Yuto/ubic-stock-management-api/stock-management-system/domain"

type TransactionRepository interface {
	FindByUserID(string) ([]domain.Transaction, error)
	Create(string, string, string) (string, error)
	Delete(string) error
}

const ()

type TransactionRepositoryErr string

func (e TransactionRepositoryErr) Error() string {
	return string(e)
}

func (it *Interactor) FindTransactionsByUserID(id string) ([]domain.Transaction, error) {
	// UserIDがidであるユーザが借りたものの一覧を返します
	return it.TransactionRepository.FindByUserID(id)
}

func (it *Interactor) CreateTransaction(stockID, userID, date string) (string, error) {
	// 新しくTransactionを作ります
	return it.TransactionRepository.Create(stockID, userID, date)
}

func (it *Interactor) DeleteTransaction(id string) error {
	return it.TransactionRepository.Delete(id)
}
