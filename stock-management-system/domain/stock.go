package domain

type Stock struct {
	ID          string `json:"id"`
	Image       string `json:"image"`
	Amount      int    `json:"amount"`
	MakerName   string `json:"makername"`
	ProductName string `json:"productname"`
}
