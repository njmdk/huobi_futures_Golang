package market

type GetHistoryTradeResponse struct {
	Ch     string `json:"ch"`
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Data []struct {
			ID        int64   `json:"id"`
			Price     float64 `json:"price"`
			Amount    int     `json:"amount"`
			Direction string  `json:"direction"`
			Ts        int64   `json:"ts"`
			Quantity  float64 `json:"quantity"`
		} `json:"data"`
		Id int64 `json:"id"`
		Ts int64 `json:"ts"`
	} `json:"data"`
}
