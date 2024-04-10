package market

type GetMarketBboResponse struct {
	Status string `json:"status"`
	Ticks  []struct {
		ContractCode string    `json:"contract_code"`
		Mrid         int64     `json:"mrid"`
		Ask          []float32 `json:"ask"`
		Bid          []float32 `json:"bid"`
		Ts           int64     `json:"ts"`
	} `json:"ticks"`
	Ts int64 `json:"ts"`
}
