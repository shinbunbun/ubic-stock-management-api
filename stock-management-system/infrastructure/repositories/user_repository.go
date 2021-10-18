package repositories

import (
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/domain"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"
)

type RepositoryError string

func (e RepositoryError) Error() string {
	return string(e)
}

const (
	AlreadyExistsErr = RepositoryError("Can't make user because the email was already used")
)

type UserRepository struct {
	UbicFoodHandler *database.UbicFoodHandler
}

func (ur *UserRepository) FindByID(id string) (domain.User, error) {
	userDatas, err := ur.UbicFoodHandler.GetByID(id)
	if err != nil {
		return domain.User{}, err
	}
	res := domain.User{ID: id}

	for _, data := range userDatas {
		if data.DataKind != "user" {
			continue
		}
		switch data.DataType {
		case "user-email":
			res.Email = data.Data
		case "user-name":
			res.Name = data.Data
		case "user-password":
			res.Password = data.Data
		}
	}

	return res, nil
}

func (ur *UserRepository) FindByEmail(email string) (domain.User, error) {
	w, err := ur.UbicFoodHandler.GetByDataAndDataType(email, "user-email")
	if err != nil {
		return domain.User{}, err
	}
	return ur.FindByID(w.ID)
}

func (ur *UserRepository) Create(email string, name string, password string) (string, error) {
	if _, err := ur.FindByEmail(email); err != nil {
		return "", AlreadyExistsErr
	}
	widgets := []database.UbicFoodWidget{
		{
			Data:     email,
			DataType: "user-email",
			DataKind: "user",
		},
		{
			Data:     name,
			DataType: "user-name",
			DataKind: "user",
		},
		{
			Data:     password,
			DataType: "user-password",
			DataKind: "user",
		},
	}
	return ur.UbicFoodHandler.AddMultipleItems(widgets)
}

func (ur *UserRepository) Delete(id string) error {
	return ur.UbicFoodHandler.DeleteByID(id)
}
