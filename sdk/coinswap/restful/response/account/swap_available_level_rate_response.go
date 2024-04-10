package account

type SwapAvailableLevelRateResponse struct {
	Status string `json:"status"`
	Data   []struct {
		ContractCode       string `json:"contract_code"`
		AvailableLevelRate string `json:"available_level_rate"`
	} `json:"data"`
	Timestamp int64 `json:"ts"`
}
