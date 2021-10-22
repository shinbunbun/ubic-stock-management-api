package controllers

func (c *Controller) GetStockALL() (int, string, error) {
	stocks, err := c.Interactor.FindStockAll()
	if err != nil {
		return 404, "getStockALL(): error", nil
	}
	return jsonDump(stocks)
}
