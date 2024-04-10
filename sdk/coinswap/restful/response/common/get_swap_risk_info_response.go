package common

type GetSwapRiskInfoResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		ContractCode      string  `json:"contract_code"`
		InsuranceFund     float64 `json:"insurance_fund"`
		EstimatedClawback float64 `json:"estimated_clawback"`
	} `json:"data,omitempty"`
}
