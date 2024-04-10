package triggerorder

type SwapTrackOpenOrdersResponse struct {
	Status string `json:"status"`
	Data   struct {
		TotalPage   int `json:"total_page"`
		TotalSize   int `json:"total_size"`
		CurrentPage int `json:"current_page"`
		Orders      []struct {
			Symbol         string  `json:"symbol"`
			ContractCode   string  `json:"contract_code"`
			Volume         float64 `json:"volume"`
			OrderType      int     `json:"order_type"`
			Direction      string  `json:"direction"`
			Offset         string  `json:"offset"`
			LeverRate      int     `json:"lever_rate"`
			OrderID        int64   `json:"order_id"`
			OrderIDStr     string  `json:"order_id_str"`
			OrderSource    string  `json:"order_source"`
			CreatedAt      int64   `json:"created_at"`
			OrderPriceType string  `json:"order_price_type"`
			Status         int     `json:"status"`
			CallbackRate   float64 `json:"callback_rate"`
			ActivePrice    float64 `json:"active_price"`
			IsActive       int     `json:"is_active"`
		} `json:"orders"`
	} `json:"data"`

	Ts int64 `json:"ts"`
}
