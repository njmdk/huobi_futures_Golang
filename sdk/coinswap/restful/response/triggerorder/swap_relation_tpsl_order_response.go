package triggerorder

type SwapRelationTpslOrderResponse struct {
	Status string `json:"status"`
	Data   struct {
		Symbol         string  `json:"symbol"`
		ContractCode   string  `json:"contract_code"`
		Volume         float64 `json:"volume"`
		Price          float64 `json:"price"`
		OrderPriceType string  `json:"order_price_type"`
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
		OrderStatus    int     `json:"status"`
		OrderType      int     `json:"order_type"`
		OrderSource    string  `json:"order_source"`
		FeeAsset       string  `json:"fee_asset"`
		CanceledAt     int64   `json:"canceled_at"`
		TPSLOrderInfo  []struct {
			Volume              float64 `json:"volume"`
			TPSLOrderType       string  `json:"tpsl_order_type"`
			Direction           string  `json:"direction"`
			OrderID             int64   `json:"order_id"`
			OrderIDStr          string  `json:"order_id_str"`
			TriggerType         string  `json:"trigger_type"`
			TriggerPrice        float64 `json:"trigger_price"`
			CreatedAt           int64   `json:"created_at"`
			OrderPriceType      string  `json:"order_price_type"`
			OrderPrice          float64 `json:"order_price"`
			Status              int     `json:"status"`
			RelationTPSLOrderID string  `json:"relation_tpsl_order_id"`
			CanceledAt          int64   `json:"canceled_at"`
			FailCode            int     `json:"fail_code"`
			FailReason          string  `json:"fail_reason"`
			TriggeredPrice      float64 `json:"triggered_price"`
			RelationOrderID     string  `json:"relation_order_id"`
		} `json:"tpsl_order_info,omitempty"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
