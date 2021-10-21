package repositories

import (
	"errors"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"
)

type CompleteRepository struct {
	UbicFoodHandler database.UbicFoodHandler
}

func (r *CompleteRepository) CheckCode(code string) (string, error) {
	widgetArr, err := r.UbicFoodHandler.GetByID(code)
	if err != nil {
		return "", err
	}
	email := ""
	for i := 0; i < len(widgetArr); i++ {
		if widgetArr[i].DataType == "code-email" {
			email = widgetArr[i].Data
		}
	}
	if email == "" {
		return "", errors.New("cannot find email address")
	}
	return email, nil
}

func (r *CompleteRepository) DeleteCode(code string) error {
	err := r.UbicFoodHandler.DeleteByID(code)
	if err != nil {
		return err
	}
	return nil
}

func (r *CompleteRepository) RegisterUser(email string) (string, error) {
	widget := database.UbicFoodWidget{
		DataType: "user-email",
		Data:     email,
		DataKind: "user",
	}
	return r.UbicFoodHandler.AddItem(widget)
}
