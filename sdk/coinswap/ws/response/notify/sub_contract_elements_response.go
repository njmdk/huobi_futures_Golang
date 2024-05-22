package notify

type SubContractElementsResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

	Uid string `json:"uid"`

	EventSender string `json:"event"`

	Data []struct {
		ContractCode        string  `json:"contract_code"`
		InstrumentIndexCode string  `json:"instrument_index_code"`
		RealTimeSettlement  int     `json:"real_time_settlement"`
		TransferProfitRatio float64 `json:"transfer_profit_ratio"`
		MinLevel            int     `json:"min_level"`
		MaxLevel            int     `json:"max_level"`
		ContractInfos       []struct {
			ContractCode   string `json:"contract_code"`
			DeliveryTime   string `json:"delivery_time"`
			CreateDate     string `json:"create_date"`
			ContractStatus int    `json:"contract_status"`
			SettlementDate string `json:"settlement_date"`
		} `json:"contract_infos"`
		OpenOrderLimit     int    `json:"open_order_limit"`
		OffsetOrderLimit   int    `json:"offset_order_limit"`
		LongPositionLimit  int    `json:"long_position_limit"`
		ShortPositionLimit int    `json:"short_position_limit"`
		PriceTick          string `json:"price_tick"`
		InstrumentValue    string `json:"instrument_value"`
		SettlePeriod       int    `json:"settle_period"`
		FundingRateCap     int    `json:"funding_rate_cap"`
		FundingRateFloor   int    `json:"funding_rate_floor"`
		HigNormalLimit     int    `json:"hig_normal_limit"`
		MinNormalLimit     int    `json:"min_normal_limit"`
		HigOpenLimit       int    `json:"hig_open_limit"`
		MinOpenLimit       int    `json:"min_open_limit"`
		HigTradeLimit      int    `json:"hig_trade_limit"`
		MinTradeLimit      int    `json:"min_trade_limit"`
	} `json:"data"`
}
