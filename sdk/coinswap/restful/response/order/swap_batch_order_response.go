package order

type SwapBatchOrderResponse struct {
	Status string `json:"status"`
	Data   struct {
		Errors []struct {
			Index   int    `json:"index"`
			ErrCode int    `json:"err_code"`
			ErrMsg  string `json:"err_msg"`
		} `json:"errors,omitempty"`
		Success []struct {
			Index         int64  `json:"index"`
			OrderID       int64  `json:"order_id"`
			OrderIDStr    string `json:"order_id_str"`
			ClientOrderID int64  `json:"client_order_id,omitempty"`
		} `json:"success,omitempty"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
