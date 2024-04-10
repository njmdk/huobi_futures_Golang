package common

type GetSwapIndexResponse struct {
	Status string `json:"status"`
	Data   []struct {
		ContractCode string  `json:"contract_code"`
		IndexPrice   float64 `json:"index_price"`
		IndexTs      int64   `json:"index_ts"`
	} `json:"data,omitempty"`
	Ts int64 `json:"ts"`
}
