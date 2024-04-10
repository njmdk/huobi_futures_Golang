package account

type SwapApiTradingStatusResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		IsDisable       int    `json:"is_disable"`
		OrderPriceTypes string `json:"order_price_types"`
		DisableReason   string `json:"disable_reason"`
		DisableInterval int64  `json:"disable_interval"`
		RecoveryTime    int64  `json:"recovery_time"`
		COR             struct {
			OrdersThreshold      int64   `json:"orders_threshold"`
			Orders               int64   `json:"orders"`
			InvalidCancelOrders  int64   `json:"invalid_cancel_orders"`
			CancelRatioThreshold float32 `json:"cancel_ratio_threshold"`
			CancelRatio          float32 `json:"cancel_ratio"`
			IsTrigger            int     `json:"is_trigger"`
			IsActive             int     `json:"is_active"`
		} `json:"COR"`
		TDN struct {
			DisablesThreshold int64 `json:"disables_threshold"`
			Disables          int64 `json:"disables"`
			IsTrigger         int   `json:"is_trigger"`
			IsActive          int   `json:"is_active"`
		} `json:"TDN"`
	} `json:"data"`
}
