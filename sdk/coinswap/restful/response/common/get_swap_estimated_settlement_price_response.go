package common

type GetSwapEstimatedSettlementPriceResponse struct {
	Status string `json:"status"`
	Data   []struct {
		ContractCode             string  `json:"contract_code"`
		EstimatedSettlementPrice float64 `json:"estimated_settlement_price"`
		SettlementType           string  `json:"settlement_type"`
	} `json:"data,omitempty"`
	Ts int64 `json:"ts"`
}
