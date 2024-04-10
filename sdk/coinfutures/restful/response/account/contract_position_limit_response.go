package account

type ContractPositionLimitResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Symbol string `json:"symbol"`
		List   []struct {
			ContractType string  `json:"contract_type"`
			BuyLimit     float64 `json:"buy_limit"`
			SellLimit    float64 `json:"sell_limit"`
		} `json:"list"`
	} `json:"data"`
}
