package common

type GetSwapContractInfoResponse struct {
	Status string `json:"status"`
	Data   []struct {
		Symbol         string  `json:"symbol"`
		ContractCode   string  `json:"contract_code"`
		ContractSize   float64 `json:"contract_size"`
		PriceTick      float64 `json:"price_tick"`
		SettlementDate string  `json:"settlement_date"`
		CreateDate     string  `json:"create_date"`
		DeliveryTime   string  `json:"delivery_time"`
		ContractStatus int     `json:"contract_status"`
	} `json:"data,omitempty"`
	Ts int64 `json:"ts"`
}
