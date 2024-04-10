package common

type GetContractEstimatedSettlementPriceResponse struct {
	Status string `json:"status"`
	Data   []struct {
		Symbol string `json:"symbol"`
		List   []struct {
			ContractType             string  `json:"contract_type"`
			ContractCode             string  `json:"contract_code"`
			EstimatedSettlementPrice float64 `json:"estimated_settlement_price"`
			SettlementType           string  `json:"settlement_type"`
		} `json:"list"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
