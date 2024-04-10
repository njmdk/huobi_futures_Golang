package market

type GetMarkPriceKlineResponse struct {
	Ch   string `json:"ch"`
	Data []struct {
		ID     int64  `json:"id"`
		Vol    string `json:"vol"`
		Count  string `json:"count"`
		Open   string `json:"open"`
		Close  string `json:"close"`
		Low    string `json:"low"`
		High   string `json:"high"`
		Amount string `json:"amount"`
	} `json:"data,omitempty"`
	Status string  `json:"status"`
	Ts     float64 `json:"ts"`
}
