package common

type GetContractPriceLimitResponse struct {
	Status string `json:"status"`
	Data   []struct {
		Symbol       string  `json:"symbol"`
		HighLimit    float64 `json:"high_limit"`
		LowLimit     float64 `json:"low_limit"`
		ContractCode string  `json:"contract_code"`
		ContractType string  `json:"contract_type"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
