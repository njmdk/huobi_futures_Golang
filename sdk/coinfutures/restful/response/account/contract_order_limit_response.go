package account

type ContractOrderLimitResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		OrderPriceType string `json:"order_price_type"`
		List           []struct {
			Symbol string `json:"symbol"`
			Types  []struct {
				ContractType string `json:"contract_type"`
				OpenLimit    int64  `json:"open_limit"`
				CloseLimit   int64  `json:"close_limit"`
			} `json:"types"`
		} `json:"list"`
	} `json:"data"`
}
