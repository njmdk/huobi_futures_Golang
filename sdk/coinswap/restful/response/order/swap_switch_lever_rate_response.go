package order

type SwapSwitchLeverRateResponse struct {
	Status string `json:"status"`
	Data   struct {
		ContractCode string `json:"contract_code"`
		LeverRate    int    `json:"lever_rate"`
	} `json:"data,omitempty"`
	ErrCode int    `json:"err_code,omitempty"`
	ErrMsg  string `json:"err_msg,omitempty"`
	Ts      int64  `json:"ts"`
}
