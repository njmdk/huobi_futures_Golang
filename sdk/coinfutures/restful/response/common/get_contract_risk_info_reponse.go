package common

type GetContractRiskInfoResponse struct {
	Status string                `json:"status"`
	Ts     int64                 `json:"ts"`
	Data   []GetContractRiskInfo `json:"data,omitempty"`
}

type GetContractRiskInfo struct {
	Symbol            string  `json:"symbol"`
	InsuranceFund     float64 `json:"insurance_fund"`
	EstimatedClawback float64 `json:"estimated_clawback"`
}
