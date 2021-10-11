package usecase

import (
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/domain"
)

const (
	UserNotFoundError   = UserRepositoryErr("User Not Found Error")
	CantDeleteUserError = UserRepositoryErr("Can't delete user, because there are no correspond user")
	FailedCreateUser    = UserRepositoryErr("Can't create user, because your email was already used")
)

type UserRepositoryErr string

func (e UserRepositoryErr) Error() string {
	return string(e)
}

type Interactor struct {
	UserRepository UserRepository
}

func (it *Interactor) FindUserByID(id string) (domain.User, error) {
	return it.UserRepository.FindByID(id)
}

func (it *Interactor) FindUserByEmail(email string) (domain.User, error) {
	return it.UserRepository.FindByEmail(email)
}

func (it *Interactor) DeleteUserByID(id string) error {
	_, err := it.UserRepository.FindByID(id)

	switch err {
	case UserNotFoundError:
		return CantDeleteUserError
	case nil:
		return it.UserRepository.Delete(id)
	default:
		return err
	}
}

func (it *Interactor) CreateUser(email string, name string, password string) (string, error) {
	_, err := it.UserRepository.FindByEmail(email)
	switch err {
	case nil:
		return "", FailedCreateUser
	case UserNotFoundError:

	default:
		return "", err
	}
	return it.UserRepository.Create(email, name, password)
}
