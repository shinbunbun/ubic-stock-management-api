package repositories

import "github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"

type RegisterRepository struct {
	UbicFoodHandler database.UbicFoodHandler
}

func (r *RegisterRepository) AddCodeToDB(email string) (string, error) {
	widget := database.UbicFoodWidget{
		DataType: "code-email",
		Data:     email,
		DataKind: "code",
	}
	code, err := r.UbicFoodHandler.AddItem(widget)
	if err != nil {
		return "", err
	}
	return code, nil
}
