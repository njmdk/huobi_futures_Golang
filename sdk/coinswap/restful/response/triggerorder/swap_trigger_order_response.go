package triggerorder

type SwapTriggerOrderResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		OrderID    int    `json:"order_id"`
		OrderIDStr string `json:"order_id_str"`
	} `json:"data"`
}
