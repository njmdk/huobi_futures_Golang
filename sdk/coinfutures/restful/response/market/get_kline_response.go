package market

type GetKlineResponse struct {
	Ch   string `json:"ch"`
	Data []struct {
		ID     int     `json:"id"`
		Vol    float64 `json:"vol"`
		Count  float64 `json:"count"`
		Open   float64 `json:"open"`
		Close  float64 `json:"close"`
		Low    float64 `json:"low"`
		High   float64 `json:"high"`
		Amount float64 `json:"amount"`
	} `json:"data,omitempty"`
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
}
