package repositories

import "github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"

type StockRepositor struct {
	UbicFoodHandler *database.UbicFoodHandler
}
