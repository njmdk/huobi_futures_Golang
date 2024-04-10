package order

type ContractMatchResultsExactResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data []struct {
		ID               string  `json:"id"`
		MatchID          int64   `json:"match_id"`
		OrderID          int64   `json:"order_id"`
		OrderIDStr       string  `json:"order_id_str"`
		Symbol           string  `json:"symbol"`
		ContractType     string  `json:"contract_type"`
		ContractCode     string  `json:"contract_code"`
		Direction        string  `json:"direction"`
		Offset           string  `json:"offset"`
		TradeVolume      float32 `json:"trade_volume"`
		TradePrice       float32 `json:"trade_price"`
		TradeTurnover    float32 `json:"trade_turnover"`
		CreateDate       int64   `json:"create_date"`
		OffsetProfitLoss float32 `json:"offset_profitloss"`
		TradeFee         float32 `json:"trade_fee"`
		Role             string  `json:"role"`
		RealProfit       float32 `json:"real_profit"`
		FeeAsset         string  `json:"fee_asset"`
		OrderSource      string  `json:"order_source"`
		QueryID          int64   `json:"query_id"`
	} `json:"data"`
}
