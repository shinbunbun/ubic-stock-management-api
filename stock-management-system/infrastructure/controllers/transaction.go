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
