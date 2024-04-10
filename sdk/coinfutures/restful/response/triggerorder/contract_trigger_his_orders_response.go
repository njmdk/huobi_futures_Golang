package triggerorder

type ContractTriggerHisOrdersResponse struct {
	Status string `json:"status"`
	Data   struct {
		Orders []struct {
			Symbol          string  `json:"symbol"`
			ContractCode    string  `json:"contract_code"`
			ContractType    string  `json:"contract_type"`
			TriggerType     string  `json:"trigger_type"`
			Volume          float64 `json:"volume"`
			OrderType       int     `json:"order_type"`
			Direction       string  `json:"direction"`
			Offset          string  `json:"offset"`
			LeverRate       int     `json:"lever_rate"`
			OrderID         int     `json:"order_id"`
			OrderIDStr      string  `json:"order_id_str"`
			RelationOrderID string  `json:"relation_order_id"`
			OrderPriceType  string  `json:"order_price_type"`
			Status          int     `json:"status"`
			OrderSource     string  `json:"order_source"`
			TriggerPrice    float64 `json:"trigger_price"`
			TriggeredPrice  float64 `json:"triggered_price"`
			OrderPrice      float64 `json:"order_price"`
			CreatedAt       int64   `json:"created_at"`
			TriggeredAt     int64   `json:"triggered_at"`
			OrderInsertAt   int64   `json:"order_insert_at"`
			CanceledAt      int64   `json:"canceled_at"`
			UpdateTime      int64   `json:"update_time"`
			FailCode        int     `json:"fail_code"`
			FailReason      string  `json:"fail_reason"`
		} `json:"orders"`
		TotalPage   int `json:"total_page"`
		CurrentPage int `json:"current_page"`
		TotalSize   int `json:"total_size"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
