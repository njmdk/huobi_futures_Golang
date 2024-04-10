package account

type AccountTransferResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		OrderId string `json:"order_id"`

		ClientOrderId string `json:"client_order_id,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
