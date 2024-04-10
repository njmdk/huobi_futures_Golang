package triggerorder

type ContractTpslCancelAllResponse struct {
	Status string `json:"status"`
	Data   struct {
		Errors []struct {
			OrderID string `json:"order_id"`
			ErrCode int64  `json:"err_code,omitempty"`
			ErrMsg  string `json:"err_msg,omitempty"`
		} `json:"errors,omitempty"`
		Successes string `json:"successes"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
