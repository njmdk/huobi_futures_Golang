package account

type SwapLeverPositionLimitResponse struct {
	Status string                   `json:"status"`
	Data   []SwapLeverPositionLimit `json:"data"`
	Ts     int64                    `json:"ts"`
}

type SwapLeverPositionLimit struct {
	Symbol       string `json:"symbol"`
	ContractCode string `json:"contract_code"`
	MarginMode   string `json:"margin_mode"`
	List         []List `json:"list"`
}

type List struct {
	LeverRate      int     `json:"lever_rate"`
	BuyLimitValue  float64 `json:"buy_limit_value"`
	SellLimitValue float64 `json:"sell_limit_value"`
}
