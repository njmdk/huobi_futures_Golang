package market

type GetBatchMergedV2Response struct {
	Status string `json:"status"`
	Ticks  Tick   `json:"ticks"`
	Ts     int64  `json:"ts"`
}

type Tick struct {
	ContractCode string    `json:"contract_code"`
	BusinessType string    `json:"business_type"`
	ID           int64     `json:"id"`
	Amount       string    `json:"amount"`
	Ask          []float64 `json:"ask"`
	Bid          []float64 `json:"bid"`
	Open         string    `json:"open"`
	Close        string    `json:"close"`
	Count        float64   `json:"count"`
	High         string    `json:"high"`
	Low          string    `json:"low"`
	Volume       string    `json:"vol"`
	NumberOf     string    `json:"number_of"`
	Timestamp    int64     `json:"ts"`
}
