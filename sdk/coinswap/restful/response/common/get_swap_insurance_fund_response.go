package common

type GetSwapInsuranceFundResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		Symbol       string `json:"symbol"`
		ContractCode string `json:"contract_code"`
		Tick         []struct {
			InsuranceFund float64 `json:"insurance_fund"`
			Ts            int64   `json:"ts"`
		} `json:"tick,omitempty"`
		TotalPage   int `json:"total_page"`
		CurrentPage int `json:"current_page"`
		TotalSize   int `json:"total_size"`
	} `json:"data,omitempty"`
}
