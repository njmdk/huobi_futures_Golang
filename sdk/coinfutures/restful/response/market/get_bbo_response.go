package market

type GetMarketBboResponse struct {
	Status string `json:"status"`
	Ticks  []struct {
		Symbol string    `json:"symbol"`
		MRID   int64     `json:"mrid"`
		Ask    []float32 `json:"ask,omitempty"`
		Bid    []float32 `json:"bid,omitempty"`
		Ts     int64     `json:"ts"`
	} `json:"ticks,omitempty"`
	Ts int64 `json:"ts"`
}
