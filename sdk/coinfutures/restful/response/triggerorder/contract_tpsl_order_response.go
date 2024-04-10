package triggerorder

type ContractTpslOrderResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		TpOrder struct {
			OrderID    int64  `json:"order_id"`
			OrderIDStr string `json:"order_id_str"`
		} `json:"tp_order"`
		SlOrder struct {
			OrderID    int64  `json:"order_id"`
			OrderIDStr string `json:"order_id_str"`
		} `json:"sl_order"`
	} `json:"data"`
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}
