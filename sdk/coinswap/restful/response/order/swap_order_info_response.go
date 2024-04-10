package order

type SwapOrderInfoResponse struct {
	Status string `json:"status"`
	Data   []struct {
		Symbol          string  `json:"symbol"`
		ContractCode    string  `json:"contract_code"`
		Volume          float64 `json:"volume"`
		Price           float64 `json:"price"`
		OrderPriceType  string  `json:"order_price_type"`
		Direction       string  `json:"direction"`
		Offset          string  `json:"offset"`
		LeverRate       int     `json:"lever_rate"`
		OrderID         int64   `json:"order_id"`
		OrderIDStr      string  `json:"order_id_str"`
		ClientOrderID   int64   `json:"client_order_id"`
		CreatedAt       int64   `json:"created_at"`
		TradeVolume     float64 `json:"trade_volume"`
		TradeTurnover   float64 `json:"trade_turnover"`
		Fee             float64 `json:"fee"`
		FeeAsset        string  `json:"fee_asset"`
		TradeAvgPrice   float64 `json:"trade_avg_price"`
		MarginFrozen    float64 `json:"margin_frozen"`
		Profit          float64 `json:"profit"`
		Status          int     `json:"status"`
		OrderType       int     `json:"order_type"`
		OrderSource     string  `json:"order_source"`
		LiquidationType string  `json:"liquidation_type"`
		CanceledAt      int64   `json:"canceled_at"`
		IsTpsl          int     `json:"is_tpsl"`
		RealProfit      float64 `json:"real_profit"`
		CanceledSource  string  `json:"canceled_source,omitempty"`
	} `json:"data,omitempty"`
	Ts int64 `json:"ts"`
}
