package order

type SwapOpenOrdersResponse struct {
	Status string `json:"status"`
	Data   struct {
		Orders []struct {
			Symbol         string  `json:"symbol"`
			ContractCode   string  `json:"contract_code"`
			Volume         float64 `json:"volume"`
			Price          float64 `json:"price"`
			OrderPriceType string  `json:"order_price_type"`
			OrderType      int     `json:"order_type"`
			Direction      string  `json:"direction"`
			Offset         string  `json:"offset"`
			LeverRate      int     `json:"lever_rate"`
			OrderID        int64   `json:"order_id"`
			OrderIDStr     string  `json:"order_id_str"`
			ClientOrderID  int64   `json:"client_order_id"`
			CreatedAt      int64   `json:"created_at"`
			TradeVolume    float64 `json:"trade_volume"`
			TradeTurnover  float64 `json:"trade_turnover"`
			Fee            float64 `json:"fee"`
			FeeAsset       string  `json:"fee_asset"`
			TradeAvgPrice  float64 `json:"trade_avg_price"`
			MarginFrozen   float64 `json:"margin_frozen"`
			Profit         float64 `json:"profit"`
			IsTPSL         int     `json:"is_tpsl"`
			UpdateTime     int64   `json:"update_time"`
			RealProfit     float64 `json:"real_profit"`
			OrderSource    string  `json:"order_source"`
		} `json:"orders"`
		TotalPage   int `json:"total_page"`
		CurrentPage int `json:"current_page"`
		TotalSize   int `json:"total_size"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
