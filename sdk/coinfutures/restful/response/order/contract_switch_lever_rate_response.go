package order

type ContractSwitchLeverRateResponse struct {
	Status string `json:"status"`
	Data   struct {
		Symbol    string `json:"symbol"`
		LeverRate int    `json:"lever_rate"`
	} `json:"data,omitempty"`
	Ts int64 `json:"ts"`
}
