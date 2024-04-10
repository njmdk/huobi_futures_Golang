package triggerorder

type SwapTpslOpenOrdersResponse struct {
	Status string `json:"status"`
	Data   struct {
		TotalPage   int `json:"total_page"`
		TotalSize   int `json:"total_size"`
		CurrentPage int `json:"current_page"`
		Orders      []struct {
			Symbol              string  `json:"symbol"`
			ContractCode        string  `json:"contract_code"`
			Volume              float32 `json:"volume"`
			OrderType           int     `json:"order_type"`
			TPSLOrderType       string  `json:"tpsl_order_type"`
			Direction           string  `json:"direction"`
			OrderID             int64   `json:"order_id"`
			OrderIDString       string  `json:"order_id_str"`
			OrderSource         string  `json:"order_source"`
			TriggerType         string  `json:"trigger_type"`
			TriggerPrice        float32 `json:"trigger_price"`
			CreatedAt           int64   `json:"created_at"`
			OrderPriceType      string  `json:"order_price_type"`
			OrderPrice          float32 `json:"order_price"`
			Status              int     `json:"status"`
			SourceOrderID       string  `json:"source_order_id"`
			RelationTPSLOrderID string  `json:"relation_tpsl_order_id"`
		} `json:"orders"`
	} `json:"data"`
	Timestamp int64 `json:"ts"`
}
