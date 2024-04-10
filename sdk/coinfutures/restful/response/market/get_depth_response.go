package market

type Tick struct {
	Asks    [][]float32 `json:"asks,omitempty"`
	Bids    [][]float32 `json:"bids,omitempty"`
	ID      int64       `json:"id"`
	MRID    int64       `json:"mrid"`
	Ts      int64       `json:"ts"`
	Version int64       `json:"version"`
}

type GetDepthResponse struct {
	Ch     string `json:"ch"`
	Status string `json:"status"`
	Tick   Tick   `json:"tick,omitempty"`
	Ts     int64  `json:"ts"`
}
