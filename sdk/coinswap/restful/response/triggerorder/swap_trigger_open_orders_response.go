package triggerorder

type SwapTriggerOpenOrdersResponse struct {
	Status string `json:"status"`
	Data   struct {
		TotalPage   int `json:"total_page"`
		CurrentPage int `json:"current_page"`
		TotalSize   int `json:"total_size"`
		Orders      []struct {
			Symbol         string  `json:"symbol"`
			ContractCode   string  `json:"contract_code"`
			TriggerType    string  `json:"trigger_type"`
			Volume         float64 `json:"volume"`
			OrderType      int     `json:"order_type"`
			Direction      string  `json:"direction"`
			Offset         string  `json:"offset"`
			LeverRate      int     `json:"lever_rate"`
			OrderID        int64   `json:"order_id"`
			OrderIDStr     string  `json:"order_id_str"`
			OrderSource    string  `json:"order_source"`
			TriggerPrice   float64 `json:"trigger_price"`
			OrderPrice     float64 `json:"order_price"`
			CreatedAt      int64   `json:"created_at"`
			OrderPriceType string  `json:"order_price_type"`
			Status         int     `json:"status"`
		} `json:"orders"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
