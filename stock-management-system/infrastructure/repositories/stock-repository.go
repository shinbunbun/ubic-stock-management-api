package repositories

import (
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/domain"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"
)

type StockRepositor struct {
	UbicFoodHandler *database.UbicFoodHandler
}

func (sr *StockRepositor) FindAll() ([]domain.Stock, error) {
	// 全てのStockを返します
	widgets, err := sr.UbicFoodHandler.GetByDataKind("food")
	if err != nil {
		return []domain.Stock{}, err
	}
	return newStocks(widgets)
}

func (sr *StockRepositor) FindByID(id string) (domain.Stock, error) {
	// IDのStockを返します
	widgets, err := sr.UbicFoodHandler.GetByID(id)
	if err != nil {
		return domain.Stock{}, err
	}
	return newStock(widgets)
}

func newStocks(widgets []database.UbicFoodWidget) ([]domain.Stock, error) {
	table := make(map[string][]database.UbicFoodWidget)
	for _, widget := range widgets {
		id := widget.ID
		arr, ok := table[id]
		if !ok {
			arr = make([]database.UbicFoodWidget, 0, 4)
		}
		arr = append(arr, widget)
		table[id] = arr
	}
	res := make([]domain.Stock, 0)
	for _, datas := range table {
		data, err := newStock(datas)
		if err == nil {
			res = append(res, data)
		}
	}
	return res, nil
}

func newStock(widgets []database.UbicFoodWidget) (domain.Stock, error) {
	res := domain.Stock{}
	id := ""
	for _, widget := range widgets {
		if widget.DataKind != "food" {
			continue
		}
		if id == "" {
			id = widget.ID
		} else if id != widget.ID {
			return domain.Stock{}, StockNotFoundErr
		}
		switch widget.DataType {
		case "food-image":
			res.Image = widget.Data
		case "food-maker":
			res.MakerName = widget.Data
		case "food-name":
			res.ProductName = widget.Data
		case "food-stock":
			res.Amount = widget.IntData
		}
	}
	if id == "" {
		return domain.Stock{}, StockNotFoundErr
	}
	res.ID = id
	return res, nil
}
