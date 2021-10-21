package stock

type Stock struct {
	ID          string `json:"id"`
	Amount      int    `json:"amount"`
	MakerName   string `json:"makername"`
	ProductName string `json:"productname"`
}
