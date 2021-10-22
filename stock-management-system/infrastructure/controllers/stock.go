package controllers

func (c *Controller) GetStockALL() (int, string, error) {
	stocks, err := c.Interactor.FindStockAll()
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

func (c *Controller) DeleteStock(id string) (int, string, error) {
	err := c.Interactor.DeleteStock(id)
	if err != nil {
		return internalErrorMessage("Failed to delete stock")
	}
	return message("successful delete stock!")
}
