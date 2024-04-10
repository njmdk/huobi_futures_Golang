package common

type GetContractOpenInterestResponse struct {
	Status string `json:"status"`
	Data   []struct {
		Symbol        string  `json:"symbol"`
		ContractType  string  `json:"contract_type"`
		Volume        float64 `json:"volume"`
		Amount        float64 `json:"amount"`
		ContractCode  string  `json:"contract_code"`
		TradeAmount   float64 `json:"trade_amount"`
		TradeVolume   float64 `json:"trade_volume"`
		TradeTurnover float64 `json:"trade_turnover"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
