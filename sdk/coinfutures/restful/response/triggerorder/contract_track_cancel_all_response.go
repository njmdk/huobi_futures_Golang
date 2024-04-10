package triggerorder

type ContractTrackCancelAllResponse struct {
	Status string `json:"status"`
	Data   struct {
		Errors []struct {
			OrderID string `json:"order_id"`
			ErrCode int64  `json:"err_code,omitempty"`
			ErrMsg  string `json:"err_msg,omitempty"`
		} `json:"errors,omitempty"`
		Successes string `json:"successes,omitempty"`
	} `json:"data,omitempty"`
	Timestamp int64 `json:"ts"`
}
