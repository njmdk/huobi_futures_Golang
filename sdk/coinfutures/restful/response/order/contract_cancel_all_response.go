package order

type ContractCancelAllResponse struct {
	Status string `json:"status"`
	Data   struct {
		Errors []struct {
			OrderID   string `json:"order_id"`
			ErrorCode int    `json:"err_code"`
			ErrorMsg  string `json:"err_msg"`
		} `json:"errors,omitempty"`
		Successes string `json:"successes,omitempty"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
