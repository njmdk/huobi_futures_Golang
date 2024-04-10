package transfer

type FuturesTransferResponse struct {
	Status  string `json:"status"`
	Data    int64  `json:"data"`
	ErrCode string `json:"err-code"`
	ErrMsg  string `json:"err-msg"`
}
