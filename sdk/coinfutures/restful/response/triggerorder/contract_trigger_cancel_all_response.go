package triggerorder

type ContractTriggerCancelAllResponse struct {
	Status string `json:"status"`
	Data   struct {
		Errors []struct {
			OrderID string `json:"order_id,omitempty"`
			ErrCode int    `json:"err_code,omitempty"`
			ErrMsg  string `json:"err_msg,omitempty"`
		} `json:"errors,omitempty"`
		Successes string `json:"successes,omitempty"`
	} `json:"data,omitempty"`
	Ts int64 `json:"ts"`
}
