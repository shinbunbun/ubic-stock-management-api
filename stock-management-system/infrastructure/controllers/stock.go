package controllers

import "github.com/Yuto/ubic-stock-management-api/stock-management-system/domain"

func (c *Controller) GetStockALL() (int, string, error) {
	stocks, err := c.Interactor.FindStockAll()
	if err != nil {
		return internalErrorMessage("Failed to get stocks")
	}
	return jsonDump(stocks)
}

func (c *Controller) GetStockLikeName(likeName string) (int, string, error) {
	stocks, err := c.Interactor.FindStockLikeName(likeName)
	if err != nil {
		return internalErrorMessage("Failed to get stocks")
	}
	return jsonDump(stocks)
}

func (c *Controller) ChangeStockAmount(id string, add int) (int, string, error) {
	err := c.Interactor.ChangeStockAmount(id, add)
	if err != nil {
		return internalErrorMessage("Failed to change amount!")
	}
	return message("successful change amount!")
}

func (c *Controller) CreateStock(image string, makerName string, productName string, amount int) (int, string, error) {
	id, err := c.Interactor.CreateStock(image, makerName, productName, amount)
	if err != nil {
		return internalErrorMessage("Failed to create stock")
	}
	res := domain.Stock{
		ID:          id,
		Image:       image,
		MakerName:   makerName,
		Amount:      amount,
		ProductName: productName,
	}
	return jsonDump(res)
}

func (c *Controller) DeleteStock(id string) (int, string, error) {
	err := c.Interactor.DeleteStock(id)
	if err != nil {
		return internalErrorMessage("Failed to delete stock")
	}
	return message("successful delete stock!")
}
