package triggerorder

type SwapTrackOrderResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		OrderID    int64  `json:"order_id"`
		OrderIDStr string `json:"order_id_str"`
	} `json:"data"`
}
