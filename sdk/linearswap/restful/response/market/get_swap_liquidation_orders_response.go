package market

type GetSwapLiquidationOrdersResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data []struct {
		QueryID       int64   `json:"query_id"`
		Symbol        string  `json:"symbol"`
		ContractCode  string  `json:"contract_code"`
		CreatedAt     int64   `json:"created_at"`
		Direction     string  `json:"direction"`
		Offset        string  `json:"offset"`
		Volume        float64 `json:"volume"`
		Amount        float64 `json:"amount"`
		Price         float64 `json:"price"`
		TradeTurnover float64 `json:"trade_turnover"`
		Pair          string  `json:"pair"`
		BusinessType  string  `json:"business_type"`
	} `json:"data"`
}
