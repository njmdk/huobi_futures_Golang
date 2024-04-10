package order

type SwapOrderResponse struct {
	Status string `json:"status"`
	Data   struct {
		OrderID       int64  `json:"order_id"`
		OrderIDStr    string `json:"order_id_str"`
		ClientOrderID int64  `json:"client_order_id,omitempty"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
