package usecase

import "github.com/Yuto/ubic-stock-management-api/stock-management-system/domain"

type UserRepository interface {
	FindByEmail(string) (domain.User, error)
	FindByID(string) (domain.User, error)
	Delete(string) error
	Create(string, string, string) (string, error)
}
