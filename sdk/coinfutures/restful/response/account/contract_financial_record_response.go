package account

type ContractFinancialRecordResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data []struct {
		QueryID      int64   `json:"query_id"`
		ID           int64   `json:"id"`
		Ts           int64   `json:"ts"`
		Symbol       string  `json:"symbol"`
		ContractCode string  `json:"contract_code"`
		Type         int     `json:"type"`
		Amount       float64 `json:"amount"`
	} `json:"data"`
}
