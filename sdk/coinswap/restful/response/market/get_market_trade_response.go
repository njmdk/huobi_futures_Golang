package market

type GetMarketTradeResponse struct {
	Ch     string `json:"ch"`
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Tick   struct {
		Data []struct {
			Price        string `json:"price"`
			Amount       string `json:"amount"`
			Direction    string `json:"direction"`
			Quantity     string `json:"quantity"`
			ContractCode string `json:"contract_code"`
			Id           int64  `json:"id"`
			Ts           int64  `json:"ts"`
		} `json:"data"`
		ID int64 `json:"id"`
		Ts int64 `json:"ts"`
	} `json:"tick"`
}
