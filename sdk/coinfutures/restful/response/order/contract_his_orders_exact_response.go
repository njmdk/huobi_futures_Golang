package order

type ContractHisOrdersExactResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data []struct {
		QueryID         int64   `json:"query_id"`
		OrderID         int64   `json:"order_id"`
		OrderIDStr      string  `json:"order_id_str"`
		Symbol          string  `json:"symbol"`
		ContractCode    string  `json:"contract_code"`
		ContractType    string  `json:"contract_type"`
		LeverRate       int     `json:"lever_rate"`
		Direction       string  `json:"direction"`
		Offset          string  `json:"offset"`
		Volume          float64 `json:"volume"`
		Price           float64 `json:"price"`
		CreateDate      int64   `json:"create_date"`
		UpdateTime      int64   `json:"update_time"`
		OrderSource     string  `json:"order_source"`
		OrderPriceType  string  `json:"order_price_type"`
		MarginFrozen    float64 `json:"margin_frozen"`
		Profit          float64 `json:"profit"`
		TradeVolume     float64 `json:"trade_volume"`
		TradeTurnover   float64 `json:"trade_turnover"`
		Fee             float64 `json:"fee"`
		TradeAvgPrice   float64 `json:"trade_avg_price"`
		Status          int     `json:"status"`
		OrderType       int     `json:"order_type"`
		FeeAsset        string  `json:"fee_asset"`
		LiquidationType string  `json:"liquidation_type"`
		IsTpsl          int     `json:"is_tpsl"`
		RealProfit      float64 `json:"real_profit"`
		CanceledSource  string  `json:"canceled_source,omitempty"`
	} `json:"data"`
}
