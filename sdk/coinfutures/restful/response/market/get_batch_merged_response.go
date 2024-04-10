package market

type GetdetailBatchMergedResponse struct {
	Status string `json:"status"`
	Ticks  []struct {
		Symbol   string    `json:"symbol"`
		ID       int64     `json:"id"`
		Amount   string    `json:"amount"`
		Ask      []float32 `json:"ask"`
		Bid      []float32 `json:"bid"`
		Open     string    `json:"open"`
		Close    string    `json:"close"`
		Count    float64   `json:"count"`
		High     string    `json:"high"`
		Low      string    `json:"low"`
		Vol      string    `json:"vol"`
		NumberOf string    `json:"number_of"`
		TS       int64     `json:"ts"`
	} `json:"ticks"`
	TS int64 `json:"ts"`
}
