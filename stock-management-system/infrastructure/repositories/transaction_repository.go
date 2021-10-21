package repositories

import (
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/domain"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"
)

type TransactionRepository struct {
	UbicFoodHandler *database.UbicFoodHandler
}

func (tr *TransactionRepository) FindByUserID(id string) ([]domain.Transaction, error) {
	// UserIDからTransactionを検索します
	datas, err := tr.UbicFoodHandler.GetByID(id)
	if err != nil {
		return []domain.Transaction{}, err
	}
	return newTransactions(datas)
}

func (tr *TransactionRepository) FindByID(id string) (domain.Transaction, error) {
	// UUIDからTransactionを検索します
	datas, err := tr.UbicFoodHandler.GetByID(id)
	if err != nil {
		return domain.Transaction{}, err
	}
	return newTransaction(datas)
}

func (tr *TransactionRepository) Create(stockID, userID, date string) (string, error) {
	// 与えられたデータから新しくTransactionを生成して返します
	widgets := []database.UbicFoodWidget{
		{
			Data:     stockID,
			DataType: "transaction-food",
			DataKind: "transaction",
		},
		{
			Data:     userID,
			DataType: "transaction-user",
			DataKind: "transaction",
		},
		{
			Data:     date,
			DataType: "transaction-date",
			DataKind: "transaction",
		},
	}
	return tr.UbicFoodHandler.AddMultipleItems(widgets)
}

func (tr *TransactionRepository) Delete(id string) error {
	return tr.UbicFoodHandler.DeleteByID(id)
}

func newTransactions(widgets []database.UbicFoodWidget) ([]domain.Transaction, error) {
	// widgetsから，0個以上のTransactionを生成して返します
	table := make(map[string][]database.UbicFoodWidget)
	for _, data := range widgets {
		id := data.ID
		arr, ok := table[id]
		if !ok {
			arr = make([]database.UbicFoodWidget, 0, 4)
		}
		arr = append(arr, data)
		table[id] = arr
	}
	res := make([]domain.Transaction, 0)
	for _, datas := range table {
		data, err := newTransaction(datas)
		if err == nil {
			res = append(res, data)
		}
	}
	return res, nil
}

func newTransaction(widgets []database.UbicFoodWidget) (domain.Transaction, error) {
	// widgetからTransactionを生成して返します
	res := domain.Transaction{}
	id := ""
	for _, data := range widgets {
		if data.DataKind != "transaction" {
			continue
		}
		if id == "" {
			id = data.ID
		} else if id != data.ID {
			return domain.Transaction{}, TransactionNotFoundErr
		}
		switch data.DataType {
		case "transaction-date":
			res.Date = data.Data
		case "transaction-food":
			res.StockID = data.Data
		case "transaction-user":
			res.UserID = data.Data
		}
	}
	if id == "" {
		return domain.Transaction{}, TransactionNotFoundErr
	}
	res.ID = id
	for _, data := range widgets {
		if id != data.ID {
			return domain.Transaction{}, TransactionNotFoundErr
		}
	}
	return res, nil
}
