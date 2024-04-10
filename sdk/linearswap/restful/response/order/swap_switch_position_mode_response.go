package order

type SwapSwitchPositionModeResponse struct {
	Status string                   `json:"status"`
	Data   []SwapSwitchPositionMode `json:"data"`
	Ts     int64                    `json:"ts"`
}

type SwapSwitchPositionMode struct {
	MarginAccount string `json:"margin_account"`
	PositionMode  string `json:"position_mode"`
}
