package common

type GetContractContractInfoResponse struct {
	Status string `json:"status"`
	Data   []struct {
		Symbol         string  `json:"symbol"`
		ContractCode   string  `json:"contract_code"`
		ContractType   string  `json:"contract_type"`
		ContractSize   float64 `json:"contract_size"`
		PriceTick      float64 `json:"price_tick"`
		DeliveryDate   string  `json:"delivery_date"`
		CreateDate     string  `json:"create_date"`
		SettlementTime string  `json:"settlement_time"`
		DeliveryTime   string  `json:"delivery_time"`
		ContractStatus int     `json:"contract_status"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
