package triggerorder

type SwapTrackCancelAllResponse struct {
	Status string `json:"status"`
	Data   struct {
		Errors []struct {
			OrderID string `json:"order_id"`
			ErrCode int64  `json:"err_code"`
			ErrMsg  string `json:"err_msg"`
		} `json:"errors"`
		Successes string `json:"successes"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
