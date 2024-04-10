package common

type GetContractIndexResponse struct {
	Status string `json:"status"`
	Data   []struct {
		Symbol     string  `json:"symbol"`
		IndexPrice float64 `json:"index_price"`
		IndexTs    int64   `json:"index_ts"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
