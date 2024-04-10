package market

type GetHistoryTradeResponse struct {
	Ch   string `json:"ch"`
	Data []struct {
		Amount    float64 `json:"amount"`
		Direction string  `json:"direction"`
		ID        int64   `json:"id"`
		Price     float64 `json:"price"`
		TS        int64   `json:"ts"`
		Quantity  float64 `json:"quantity"`
	} `json:"data"`
	Status string `json:"status"`
	TS     int64  `json:"ts"`
}
