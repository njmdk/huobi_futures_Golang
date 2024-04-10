package common

type GetSwapHisOpenInterestResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		Symbol       string `json:"symbol"`
		ContractCode string `json:"contract_code"`
		Tick         []struct {
			Volume     string `json:"volume"`
			AmountType int    `json:"amount_type"`
			Ts         int64  `json:"ts"`
		} `json:"tick,omitempty"`
	} `json:"data,omitempty"`
}
