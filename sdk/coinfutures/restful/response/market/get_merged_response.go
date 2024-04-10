package market

type GetDetailMergedResponse struct {
	Ch     string `json:"ch"`
	Status string `json:"status"`
	Tick   struct {
		ID     int       `json:"id"`
		Vol    string    `json:"vol"`
		Count  float64   `json:"count"`
		Open   string    `json:"open"`
		Close  string    `json:"close"`
		Low    string    `json:"low"`
		High   string    `json:"high"`
		Amount string    `json:"amount"`
		Ts     int64     `json:"ts"`
		Ask    []float32 `json:"ask"`
		Bid    []float32 `json:"bid"`
	} `json:"tick"`
	Ts int64 `json:"ts"`
}
