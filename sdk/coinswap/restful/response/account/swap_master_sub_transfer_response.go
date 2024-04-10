package account

type SwapMasterSubTransferResponse struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"ts"`
	Data      struct {
		OrderID       string `json:"order_id"`
		ClientOrderID *int64 `json:"client_order_id,omitempty"`
	} `json:"data"`
}
