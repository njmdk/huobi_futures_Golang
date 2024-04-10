package triggerorder

type ContractTrackOrderResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		OrderId    int64  `json:"order_id"`
		OrderIdStr string `json:"order_id_str"`
	} `json:"data,omitempty"`
}
