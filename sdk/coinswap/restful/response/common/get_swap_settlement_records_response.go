package common

type GetSwapSettlementRecordsResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		SettlementRecord []struct {
			Symbol          string  `json:"symbol"`
			ContractCode    string  `json:"contract_code"`
			SettlementTime  int64   `json:"settlement_time"`
			ClawbackRatio   float64 `json:"clawback_ratio"`
			SettlementPrice float64 `json:"settlement_price"`
			SettlementType  string  `json:"settlement_type"`
		} `json:"settlement_record,omitempty"`
		TotalPage   int `json:"total_page"`
		CurrentPage int `json:"current_page"`
		TotalSize   int `json:"total_size"`
	} `json:"data,omitempty"`
}
