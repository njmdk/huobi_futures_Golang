package order

type SwapOrderDetailResponse struct {
	Status string `json:"status"`
	Data   struct {
		Symbol          string  `json:"symbol"`
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
		TotalPage       int     `json:"total_page"`
		CurrentPage     int     `json:"current_page"`
		TotalSize       int     `json:"total_size"`
		InstrumentPrice float64 `json:"instrument_price"`
		FinalInterest   float64 `json:"final_interest"`
		AdjustValue     float64 `json:"adjust_value"`
		Fee             float64 `json:"fee"`
		FeeAsset        string  `json:"fee_asset"`
		LiquidationType string  `json:"liquidation_type"`
		OrderID         int64   `json:"order_id"`
		OrderIDStr      string  `json:"order_id_str"`
		ClientOrderID   int64   `json:"client_order_id"`
		OrderType       string  `json:"order_type"`
		TradeAvgPrice   float64 `json:"trade_avg_price"`
		TradeTurnover   float64 `json:"trade_turnover"`
		TradeVolume     float64 `json:"trade_volume"`
		IsTPSL          int     `json:"is_tpsl"`
		RealProfit      float64 `json:"real_profit"`
		CanceledSource  string  `json:"canceled_source"`
		Trades          []struct {
			TradeID       int64   `json:"trade_id"`
			ID            string  `json:"id"`
			TradePrice    float64 `json:"trade_price"`
			TradeVolume   float64 `json:"trade_volume"`
			TradeTurnover float64 `json:"trade_turnover"`
			TradeFee      float64 `json:"trade_fee"`
			FeeAsset      string  `json:"fee_asset"`
			Role          string  `json:"role"`
			CreatedAt     int64   `json:"created_at"`
			Profit        float64 `json:"profit"`
			RealProfit    float64 `json:"real_profit"`
		} `json:"trades"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
