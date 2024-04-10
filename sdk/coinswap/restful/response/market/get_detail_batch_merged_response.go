package market

type GetDetailBatchMergedResponse struct {
	Status string `json:"status"`
	Ticks  []struct {
		ContractCode string      `json:"contract_code"`
		ID           int64       `json:"id"`
		Amount       string      `json:"amount"`
		Ask          [][]float32 `json:"ask"`
		Bid          [][]float32 `json:"bid"`
		Open         string      `json:"open"`
		Close        string      `json:"close"`
		Count        float64     `json:"count"`
		High         string      `json:"high"`
		Low          string      `json:"low"`
		Vol          string      `json:"vol"`
		NumberOf     string      `json:"number_of"`
		Ts           int64       `json:"ts"`
	} `json:"ticks"`
	Ts int64 `json:"ts"`
}
