package order

type ContractOrderDetailResponse struct {
	Status string `json:"status"`
	Data   struct {
		Symbol          string  `json:"symbol"`
		ContractType    string  `json:"contract_type"`
		ContractCode    string  `json:"contract_code"`
		LeverRate       int     `json:"lever_rate"`
		Direction       string  `json:"direction"`
		Offset          string  `json:"offset"`
		Volume          float64 `json:"volume"`
		Price           float64 `json:"price"`
		CreatedAt       int64   `json:"created_at"`
		CanceledAt      int64   `json:"canceled_at"`
		OrderSource     string  `json:"order_source"`
		OrderPriceType  string  `json:"order_price_type"`
		MarginFrozen    float64 `json:"margin_frozen"`
		Profit          float64 `json:"profit"`
		CanceledSource  string  `json:"canceled_source"`
		TotalPage       int     `json:"total_page"`
		CurrentPage     int     `json:"current_page"`
		TotalSize       int     `json:"total_size"`
		InstrumentPrice float64 `json:"instrument_price"`
		FinalInterest   float64 `json:"final_interest"`
		AdjustValue     float64 `json:"adjust_value"`
		FeeAsset        string  `json:"fee_asset"`
		LiquidationType string  `json:"liquidation_type"`
		OrderID         int64   `json:"order_id"`
		OrderIDStr      string  `json:"order_id_str"`
		ClientOrderID   int64   `json:"client_order_id"`
		TradeVolume     float64 `json:"trade_volume"`
		TradeTurnover   float64 `json:"trade_turnover"`
		Fee             float64 `json:"fee"`
		TradeAvgPrice   float64 `json:"trade_avg_price"`
		OrderStatus     int     `json:"order_status"`
		OrderType       string  `json:"order_type"`
		IsTPSL          int     `json:"is_tpsl"`
		RealProfit      float64 `json:"real_profit"`
		Trades          []struct {
			ID            string  `json:"id"`
			TradeID       int64   `json:"trade_id"`
			TradePrice    float64 `json:"trade_price"`
			TradeVolume   float64 `json:"trade_volume"`
			TradeTurnover float64 `json:"trade_turnover"`
			TradeFee      float64 `json:"trade_fee"`
			Role          string  `json:"role"`
			CreatedAt     int64   `json:"created_at"`
			RealProfit    float64 `json:"real_profit"`
			Profit        float64 `json:"profit"`
		} `json:"trades"`
	} `json:"data"`
}
