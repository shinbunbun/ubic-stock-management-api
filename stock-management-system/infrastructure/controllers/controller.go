package controllers

import (
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
