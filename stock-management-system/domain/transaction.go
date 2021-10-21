package domain

type Transaction struct {
	ID      string `json:"id"`
	StockID string `json:"stockid"`
	UserID  string `json:"userid"`
	Date    string `json:"date"`
}
