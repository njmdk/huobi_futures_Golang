package account

type SwapPositionLimitResponse struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"ts"`
	Data      []struct {
		Symbol       string  `json:"symbol"`
		ContractCode string  `json:"contract_code"`
		BuyLimit     float64 `json:"buy_limit"`
		SellLimit    float64 `json:"sell_limit"`
	} `json:"data"`
}
