package account

type ContractAvailableLevelRateResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Symbol             string `json:"symbol"`
		AvailableLevelRate string `json:"available_level_rate"`
	} `json:"data"`
}
