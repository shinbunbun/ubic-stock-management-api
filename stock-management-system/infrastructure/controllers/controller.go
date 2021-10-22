package controllers

import (
	"bytes"
	"encoding/json"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/repositories"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/usecase"
)

type Controller struct {
	Interactor *usecase.Interactor
}

func NewController(db *database.DynamoDBHandler) *Controller {
	handler := db.NewUbicFoodHandler()
	return &Controller{
		Interactor: &usecase.Interactor{
			UserRepository: &repositories.UserRepository{
				UbicFoodHandler: handler,
			},
			StockRepository: &repositories.StockRepository{
				UbicFoodHandler: handler,
			},
			TransactionRepository: &repositories.TransactionRepository{
				UbicFoodHandler: handler,
			},
		},
	}
}

func NewControllerWithTableName(db *database.DynamoDBHandler, tableName string) *Controller {
	handler := db.NewUbicFoodHandlerWithTableName(tableName)
	return &Controller{
		Interactor: &usecase.Interactor{
			UserRepository: &repositories.UserRepository{
				UbicFoodHandler: handler,
			},
			StockRepository: &repositories.StockRepository{
				UbicFoodHandler: handler,
			},
			TransactionRepository: &repositories.TransactionRepository{
				UbicFoodHandler: handler,
			},
		},
	}
}

func jsonDump(body interface{}) (int, string, error) {
	byte, err := json.Marshal(body)

	if err != nil {
		return 404, "json dump error", nil
	}

	var buf bytes.Buffer
	json.HTMLEscape(&buf, byte)

	return 200, buf.String(), nil
}
