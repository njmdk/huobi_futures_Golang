package market

type GetMarketTradeResponse struct {
	Ch     string `json:"ch"`
	Status string `json:"status"`
	Tick   struct {
		ID   int64 `json:"id"`
		TS   int64 `json:"ts"`
		Data []struct {
			Amount    string `json:"amount"`
			Direction string `json:"direction"`
			ID        int64  `json:"id"`
			Price     string `json:"price"`
			TS        int64  `json:"ts"`
			Quantity  string `json:"quantity"`
			Symbol    string `json:"symbol"`
		} `json:"data"`
	} `json:"tick"`
	TS int64 `json:"ts"`
}
