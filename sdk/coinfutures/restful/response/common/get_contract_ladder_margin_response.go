package common

type GetContractLadderMarginResponse struct {
	Status string `json:"status"`
	Data   []struct {
		Symbol string `json:"symbol"`
		List   []struct {
			LeverRate int `json:"lever_rate"`
			Ladders   []struct {
				MinMarginBalance   float64 `json:"min_margin_balance"`
				MaxMarginBalance   float64 `json:"max_margin_balance"`
				MinMarginAvailable float64 `json:"min_margin_available"`
				MaxMarginAvailable float64 `json:"max_margin_available"`
			} `json:"ladders"`
		} `json:"list"`
	} `json:"data"`
}
