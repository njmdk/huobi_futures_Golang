package account

type ContractMasterSubTransferResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		OrderID       string `json:"order_id"`
		ClientOrderID int64  `json:"client_order_id,omitempty"`
	} `json:"data"`
}
