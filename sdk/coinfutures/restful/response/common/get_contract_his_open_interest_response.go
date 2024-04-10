package common

type GetContractHisOpenInterestResponse struct {
	Status string                  `json:"status"`
	Ts     int64                   `json:"ts"`
	Data   ContractHisOpenInterest `json:"data"`
}

type ContractHisOpenInterest struct {
	Symbol       string `json:"symbol"`
	ContractType string `json:"contract_type"`
	Tick         []struct {
		Volume     string `json:"volume"`
		AmountType int    `json:"amount_type"`
		Ts         int64  `json:"ts"`
	} `json:"tick"`
}
