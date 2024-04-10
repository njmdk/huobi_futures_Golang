package account

type SwapFinancialRecordResponse struct {
	Code int                   `json:"code"`
	Msg  string                `json:"msg"`
	Ts   int64                 `json:"ts"`
	Data []SwapFinancialRecord `json:"data"`
}

type SwapFinancialRecord struct {
	QueryID           int64   `json:"query_id"`
	ID                int64   `json:"id"`
	Asset             string  `json:"asset"`
	ContractCode      string  `json:"contract_code"`
	MarginAccount     string  `json:"margin_account"`
	FaceMarginAccount string  `json:"face_margin_account"`
	Type              int32   `json:"type"`
	Amount            float32 `json:"amount"`
}
