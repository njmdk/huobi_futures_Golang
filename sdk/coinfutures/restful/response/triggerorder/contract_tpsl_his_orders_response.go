package triggerorder

type ContractTpslHisOrdersResponse struct {
	Status string `json:"status"`
	Data   struct {
		Orders []struct {
			Symbol              string  `json:"symbol"`
			ContractType        string  `json:"contract_type"`
			ContractCode        string  `json:"contract_code"`
			Volume              float64 `json:"volume"`
			OrderType           int     `json:"order_type"`
			TPSLOrderType       string  `json:"tpsl_order_type"`
			Direction           string  `json:"direction"`
			OrderID             int64   `json:"order_id"`
			OrderIDStr          string  `json:"order_id_str"`
			OrderSource         string  `json:"order_source"`
			OrderPrice          float64 `json:"order_price"`
			TriggerType         string  `json:"trigger_type"`
			TriggerPrice        float64 `json:"trigger_price"`
			CreatedAt           int64   `json:"created_at"`
			OrderPriceType      string  `json:"order_price_type"`
			Status              int     `json:"status"`
			SourceOrderID       string  `json:"source_order_id"`
			RelationTPSLOrderID string  `json:"relation_tpsl_order_id"`
			CanceledAt          int64   `json:"canceled_at"`
			FailCode            int     `json:"fail_code"`
			FailReason          string  `json:"fail_reason"`
			TriggeredPrice      float64 `json:"triggered_price"`
			RelationOrderID     string  `json:"relation_order_id"`
			UpdateTime          int64   `json:"update_time"`
		} `json:"orders"`
		TotalPage   int `json:"total_page"`
		TotalSize   int `json:"total_size"`
		CurrentPage int `json:"current_page"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
