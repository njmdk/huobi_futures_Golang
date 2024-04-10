package account

type GetPositionLimitResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		BuyLimit float32 `json:"buy_limit"`

		SellLimit float32 `json:"sell_limit"`

		MarginMode string `json:"margin_mode"`

		LeverRate int `json:"lever_rate"`

		BuyLimitValue float32 `json:"buy_limit_value"`

		SellLimitValue float32 `json:"sell_limit_value"`

		MarkPrice float32 `json:"mark_price"`

		Pair string `json:"pair"`

		BusinessType string `json:"business_type"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
