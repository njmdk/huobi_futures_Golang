package account

type SwapCrossLeverPositionLimitResponse struct {
	Status string                        `json:"status"`
	Data   []SwapCrossLeverPositionLimit `json:"data"`
	Ts     int64                         `json:"ts"`
}

type SwapCrossLeverPositionLimit struct {
	Symbol       string  `json:"symbol"`
	ContractCode string  `json:"contract_code"`
	MarginMode   string  `json:"margin_mode"`
	BusinessType string  `json:"business_type"`
	ContractType string  `json:"contract_type"`
	Pair         string  `json:"pair"`
	List         []List1 `json:"list"`
}

type List1 struct {
	LeverRate      int     `json:"lever_rate"`
	BuyLimitValue  float64 `json:"buy_limit_value"`
	SellLimitValue float64 `json:"sell_limit_value"`
}
