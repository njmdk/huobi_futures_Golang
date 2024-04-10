package order

type ContractCancelResponse struct {
	Status string `json:"status"`
	Data   struct {
		Errors []struct {
			OrderID   string `json:"order_id"`
			ErrorCode int    `json:"err_code"`
			ErrorMsg  string `json:"err_msg"`
		} `json:"errors,omitempty"`
		Success string `json:"success,omitempty"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
