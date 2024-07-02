package order

type PlaceBatchOrderErrors struct {
	Index int `json:"index"`

	ErrorCode int `json:"err_code"`

	ErrorMessage string `json:"err_msg"`
}

type PlaceBatchOrderSuccess struct {
	Index int `json:"index"`

	OrderId int64 `json:"order_id"`

	ClientOrderId int64 `json:"client_order_id,omitempty"`

	OrderIdStr string `json:"order_id_str"`
}

type PlaceBatchOrderData struct {
	Errors []PlaceBatchOrderErrors `json:"errors,omitempty"`

	Success []PlaceBatchOrderSuccess `json:"success,omitempty"`
}
type PlaceBatchOrderResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data PlaceBatchOrderData `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
