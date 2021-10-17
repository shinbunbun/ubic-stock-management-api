package repositories

import "github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"

type RegisterRepository struct {
	UbicFoodHandler database.UbicFoodHandler
}

func (r *RegisterRepository) AddCodeToDB(email string) error {
	widget := database.UbicFoodWidget{
		DataType: "code-email",
		Data:     email,
		DataKind: "code",
	}
	r.UbicFoodHandler.AddItem(widget)
	return nil
}
