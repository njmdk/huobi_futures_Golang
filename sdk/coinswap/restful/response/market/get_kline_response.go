package market

type GetKlineResponse struct {
	Ch     string `json:"ch"`
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		ID     int64   `json:"id"`
		Vol    float64 `json:"vol"`
		Count  float64 `json:"count"`
		Open   float64 `json:"open"`
		Close  float64 `json:"close"`
		Low    float64 `json:"low"`
		High   float64 `json:"high"`
		Amount float64 `json:"amount"`
	} `json:"data"`
}
