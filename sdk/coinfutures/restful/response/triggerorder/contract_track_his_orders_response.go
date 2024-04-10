package triggerorder

type ContractTrackHisOrdersResponse struct {
	Status string `json:"status"`
	Data   struct {
		TotalPage   int `json:"total_page"`
		TotalSize   int `json:"total_size"`
		CurrentPage int `json:"current_page"`
		Orders      []struct {
			Symbol           string  `json:"symbol"`
			ContractType     string  `json:"contract_type"`
			ContractCode     string  `json:"contract_code"`
			Volume           float64 `json:"volume"`
			OrderType        int     `json:"order_type"`
			Direction        string  `json:"direction"`
			Offset           string  `json:"offset"`
			LeverRate        int     `json:"lever_rate"`
			OrderID          int64   `json:"order_id"`
			OrderIDStr       string  `json:"order_id_str"`
			OrderSource      string  `json:"order_source"`
			CreatedAt        int64   `json:"created_at"`
			UpdateTime       int64   `json:"update_time"`
			OrderPriceType   string  `json:"order_price_type"`
			Status           int     `json:"status"`
			CanceledAt       int64   `json:"canceled_at"`
			FailCode         int     `json:"fail_code"`
			FailReason       string  `json:"fail_reason"`
			CallbackRate     float64 `json:"callback_rate"`
			ActivePrice      float64 `json:"active_price"`
			IsActive         int     `json:"is_active"`
			MarketLimitPrice float64 `json:"market_limit_price"`
			FormulaPrice     float64 `json:"formula_price"`
			RealVolume       float64 `json:"real_volume"`
			TriggeredPrice   float64 `json:"triggered_price"`
			RelationOrderID  string  `json:"relation_order_id"`
		} `json:"orders"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
