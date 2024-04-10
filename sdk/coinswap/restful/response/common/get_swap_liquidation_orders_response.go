package common

type GetSwapLiquidationOrdersResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data []struct {
		Symbol       string  `json:"symbol"`
		ContractCode string  `json:"contract_code"`
		Direction    string  `json:"direction"`
		Offset       string  `json:"offset"`
		Volume       float64 `json:"volume"`
		Price        float64 `json:"price"`
		Amount       float64 `json:"amount"`
		CreatedAt    int64   `json:"created_at"`
		QueryID      int64   `json:"query_id"`
	} `json:"data,omitempty"`
}
