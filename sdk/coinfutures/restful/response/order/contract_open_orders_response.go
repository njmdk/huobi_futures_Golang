package order

type ContractOpenOrdersResponse struct {
	Status string `json:"status"`
	Data   struct {
		Orders []struct {
			Symbol         string  `json:"symbol"`
			ContractType   string  `json:"contract_type"`
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
			TradeAvgPrice  float64 `json:"trade_avg_price"`
			MarginFrozen   float64 `json:"margin_frozen"`
			Profit         float64 `json:"profit"`
			IsTPSL         int     `json:"is_tpsl"`
			Status         int     `json:"status"`
			OrderSource    string  `json:"order_source"`
			FeeAsset       string  `json:"fee_asset"`
			UpdateTime     int64   `json:"update_time"`
			RealProfit     float64 `json:"real_profit"`
		} `json:"orders"`
		TotalPage   int `json:"total_page"`
		CurrentPage int `json:"current_page"`
		TotalSize   int `json:"total_size"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
