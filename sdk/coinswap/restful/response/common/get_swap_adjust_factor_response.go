package common

type GetSwapAdjustFactorResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Symbol       string `json:"symbol"`
		ContractCode string `json:"contract_code"`
		List         []struct {
			LeverRate int `json:"lever_rate"`
			Ladders   []struct {
				MinSize      float64 `json:"min_size"`
				MaxSize      float64 `json:"max_size"`
				Ladder       int     `json:"ladder"`
				AdjustFactor float64 `json:"adjust_factor"`
			} `json:"ladders,omitempty"`
		} `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
