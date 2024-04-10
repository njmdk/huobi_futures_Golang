package market

type GetPriceLimitResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		HighLimit float32 `json:"high_limit"`

		LowLimit float32 `json:"low_limit"`

		ContractType string `json:"contract_type"`

		Pair string `json:"pair"`

		BusinessType string `json:"business_type"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
