package notify

type SubContractElementsResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

	EventSender string `json:"event"`

	Data *struct {
		ContractCode             string  `json:"contract_code"`
		ModeType                 int     `json:"mode_type"`
		SwapDeliveryType         int     `json:"swap_delivery_type"`
		InstrumentIndexCode      string  `json:"instrument_index_code"`
		RealTimeSettlement       int     `json:"real_time_settlement"`
		TransferProfitRatio      float64 `json:"transfer_profit_ratio"`
		CrossTransferProfitRatio string  `json:"cross_transfer_profit_ratio"`
		InstrumentType           []int   `json:"instrument_type"`
		TradePartition           string  `json:"trade_partition"`
		MinLevel                 int     `json:"min_level"`
		MaxLevel                 int     `json:"max_level"`
		SettlePeriod             int     `json:"settle_period"`
		FundingRateCap           int     `json:"funding_rate_cap"`
		FundingRateFloor         int     `json:"funding_rate_floor"`
		LongPositionLimit        string  `json:"long_position_limit"`
		OffsetOrderLimit         string  `json:"offset_order_limit"`
		OpenOrderLimit           string  `json:"open_order_limit"`
		ShortPositionLimit       string  `json:"short_position_limit"`
		ContractInfos            []struct {
			ContractCode   string `json:"contract_code"`
			InstrumentType int    `json:"instrument_type"`
			SettlementDate string `json:"settlement_date"`
			DeliveryTime   string `json:"delivery_time"`
			CreateDate     string `json:"create_date"`
			ContractStatus int    `json:"contract_status"`
			DeliveryDate   string `json:"delivery_date"`
		} `json:"contract_infos"`
		PriceTicks []struct {
			BusinessType int    `json:"business_type"`
			Price        string `json:"price"`
		} `json:"price_ticks"`
		InstrumentValues []struct {
			BusinessType int    `json:"business_type"`
			Price        string `json:"price"`
		} `json:"instrument_values"`
		OrderLimits []struct {
			InstrumentType   int    `json:"instrument_type"`
			Open             string `json:"open"`
			Close            string `json:"close"`
			OpenAfterClosing string `json:"open_after_closing"`
		} `json:"order_limits"`
		NormalLimits []struct {
			InstrumentType int    `json:"instrument_type"`
			Open           string `json:"open"`
			Close          string `json:"close"`
		} `json:"normal_limits"`
		OpenLimits []struct {
			InstrumentType int    `json:"instrument_type"`
			Open           string `json:"open"`
			Close          string `json:"close"`
		} `json:"open_limits"`
		TradeLimits []struct {
			InstrumentType int    `json:"instrument_type"`
			Open           string `json:"open"`
			Close          string `json:"close"`
		} `json:"trade_limits"`
	} `json:"data"`
}
