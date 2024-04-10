package common

type SettlementList struct {
	ContractCode    string  `json:"contract_code"`
	SettlementPrice float64 `json:"settlement_price"`
	SettlementType  string  `json:"settlement_type"`
}

type SettlementRecord struct {
	Symbol         string           `json:"symbol"`
	SettlementTime int64            `json:"settlement_time"`
	ClawbackRatio  float64          `json:"clawback_ratio"`
	List           []SettlementList `json:"list"`
}

type GetContractSettlementRecordsResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		SettlementRecord []SettlementRecord `json:"settlement_record"`
		TotalPage        int                `json:"total_page"`
		CurrentPage      int                `json:"current_page"`
		TotalSize        int                `json:"total_size"`
	} `json:"data"`
}
