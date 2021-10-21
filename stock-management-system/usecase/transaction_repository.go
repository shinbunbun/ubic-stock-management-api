package usecase

import "github.com/Yuto/ubic-stock-management-api/stock-management-system/domain"

type TransactionRepository interface {
	FindByUserID(string) ([]domain.Transaction, error)
	Create(string, string, string) (string, error)
	Delete(string) error
}
