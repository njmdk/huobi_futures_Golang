package account

type SwapOrderLimitResponse struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"ts"`
	Data      struct {
		OrderPriceType string `json:"order_price_type"`
		List           []struct {
			Symbol       string  `json:"symbol"`
			ContractCode string  `json:"contract_code"`
			OpenLimit    float64 `json:"open_limit"`
			CloseLimit   float64 `json:"close_limit"`
		} `json:"list"`
	} `json:"data"`
}
