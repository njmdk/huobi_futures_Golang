package market

type GetDetailMergedResponse struct {
	Ch     string `json:"ch"`
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Tick   struct {
		ID     int64     `json:"id"`
		Vol    string    `json:"vol"`
		Count  int       `json:"count"`
		Open   string    `json:"open"`
		Close  string    `json:"close"`
		Low    string    `json:"low"`
		High   string    `json:"high"`
		Amount string    `json:"amount"`
		Ask    []float32 `json:"ask"`
		Bid    []float32 `json:"bid"`
		Ts     int64     `json:"ts"`
	} `json:"tick"`
}
