package market

type GetMarketDepthResponse struct {
	Ch     string `json:"ch"`
	Status string `json:"status"`
	Tick   struct {
		Asks    [][]float32 `json:"asks"`
		Bids    [][]float32 `json:"bids"`
		Mrid    int64       `json:"mrid"`
		Ch      string      `json:"ch"`
		Id      int64       `json:"id"`
		Version int64       `json:"version"`
		Ts      int64       `json:"ts"`
	} `json:"tick"`
	Ts int64 `json:"ts"`
}
