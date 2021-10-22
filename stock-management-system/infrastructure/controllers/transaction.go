package controllers

import (
	"strconv"
	"time"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/domain"
)

func (c *Controller) Rent(stockID string, userID string) (int, string, error) {
	err := c.Interactor.ChangeStockAmount(stockID, -1)
	if err != nil {
		return internalErrorMessage("Faield to rent")
	}
	now := strconv.FormatInt(time.Now().Unix(), 10)
	id, err := c.Interactor.CreateTransaction(stockID, userID, now)
	if err != nil {
		for i := 0; i < 4; i++ {
			err := c.Interactor.ChangeStockAmount(stockID, 1)
			if err == nil {
				return internalErrorMessage("Failed to rent")
			}
		}
		panic("NOOOOOOOOOOOOOOOOOO")
	}
	res := domain.Transaction{
		ID:      id,
		StockID: stockID,
		UserID:  userID,
		Date:    now,
	}
	return jsonDump(res)
}

func (c *Controller) GiveBack(stockID string) (int, string, error) {
	err := c.Interactor.ChangeStockAmount(stockID, 1)
	if err != nil {
		return internalErrorMessage("Faield to give back")
	}
	err = c.Interactor.DeleteTransaction(stockID)
	if err != nil {
		for i := 0; i < 4; i++ {
			err := c.Interactor.ChangeStockAmount(stockID, -1)
			if err == nil {
				return internalErrorMessage("Failed to give back")
			}
		}
		panic("NOOOOOOOOOOOOOOOOOO")
	}
	return message("Successful give back!")
}

func (c *Controller) GetAllTransactions(userID string) (int, string, error) {
	transactions, err := c.Interactor.FindTransactionsByUserID(userID)
	if err != nil {
		return internalErrorMessage("Failed to get transactions")
	}
	return jsonDump(transactions)
}
