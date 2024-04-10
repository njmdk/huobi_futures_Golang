package triggerorder

type ContractTriggerOrderResponse struct {
	Status  string `json:"status"`
	ErrCode int    `json:"err_code,omitempty"`
	ErrMsg  string `json:"err_msg,omitempty"`
	Data    struct {
		OrderID    int64  `json:"order_id"`
		OrderIDStr string `json:"order_id_str"`
	} `json:"data,omitempty"`
	Ts int64 `json:"ts"`
}
