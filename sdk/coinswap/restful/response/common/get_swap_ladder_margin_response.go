package common

type GetSwapLadderMarginResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Symbol       string `json:"symbol"`
		ContractCode string `json:"contract_code"`
		List         []struct {
			LeverRate int `json:"lever_rate"`
			Ladders   []struct {
				MinMarginBalance   float64 `json:"min_margin_balance"`
				MaxMarginBalance   float64 `json:"max_margin_balance"`
				MinMarginAvailable float64 `json:"min_margin_available"`
				MaxMarginAvailable float64 `json:"max_margin_available"`
			} `json:"ladders,omitempty"`
		} `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
