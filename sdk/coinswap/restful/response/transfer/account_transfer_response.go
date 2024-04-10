package transfer

type AccountTransferResponse struct {
	Code    int64  `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    int64  `json:"data"`
}
