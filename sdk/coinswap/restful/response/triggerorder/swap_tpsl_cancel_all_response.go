package triggerorder

type SwapTpslCancelAllResponse struct {
	Status string `json:"status"`
	Data   struct {
		Errors []struct {
			OrderID      string `json:"order_id"`
			ErrorCode    int    `json:"err_code"`
			ErrorMessage string `json:"err_msg"`
		} `json:"errors"`
		Successes string `json:"successes"`
	} `json:"data"`
	Timestamp int64 `json:"ts"`
}
