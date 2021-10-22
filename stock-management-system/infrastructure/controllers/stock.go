package controllers

func (c *Controller) GetStockALL() (int, string, error) {
	stocks, err := c.Interactor.FindStockAll()
	if err != nil {
		return 404, "getStockALL(): error", nil
	}
	return jsonDump(stocks)
}

func (c *Controller) ChangeStockAmount(id string, add int) (int, string, error) {
	err := c.Interactor.ChangeStockAmount(id, add)
	if err != nil {
		return 404, "Failed to change amount", nil
	}
	return jsonDump(map[string]string{"message": "successful change!"})
}
