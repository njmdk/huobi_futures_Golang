package market

type GetHistoryIndexResponse struct {
	Ch   string `json:"ch"`
	Data []struct {
		ID     float64 `json:"id"`
		Vol    float64 `json:"vol"`
		Count  float64 `json:"count"`
		Open   float64 `json:"open"`
		Close  float64 `json:"close"`
		Low    float64 `json:"low"`
		High   float64 `json:"high"`
		Amount float64 `json:"amount"`
	} `json:"data"`
	Status string `json:"status"`
	TS     int64  `json:"ts"`
}
