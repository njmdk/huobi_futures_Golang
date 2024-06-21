package unifiedaccount

type FixPositionMarginChangeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data struct {
		OrderID       string  `json:"order_id"`
		ClientOrderID int64   `json:"client_order_id"`
		Amount        float64 `json:"amount"`
		Asset         string  `json:"asset"`
		ContractCode  string  `json:"contract_code"`
		Type          int     `json:"type"`
		Direction     int     `json:"direction"`
	} `json:"data"`
}
