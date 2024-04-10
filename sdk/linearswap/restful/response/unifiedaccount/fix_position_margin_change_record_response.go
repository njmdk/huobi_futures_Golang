package unifiedaccount

type FixPositionMarginChangeRecordResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data []struct {
		QueryID   int64   `json:"query_id"`
		OrderID   string  `json:"order_id"`
		Amount    float64 `json:"amount"`
		Asset     string  `json:"asset"`
		Symbol    string  `json:"symbol"`
		Type      int     `json:"type"`
		Direction int     `json:"direction"`
	} `json:"data"`
}
