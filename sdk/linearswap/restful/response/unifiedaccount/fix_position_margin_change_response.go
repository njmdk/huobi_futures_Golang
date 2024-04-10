package unifiedaccount

type FixPositionMarginChangeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data struct {
		OrderID       string `json:"order_id"`
		ClientOrderID int64  `json:"client_order_id"`
	} `json:"data"`
}
