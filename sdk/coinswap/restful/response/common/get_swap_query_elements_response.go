package common

type GetSwapQueryElementsResponse struct {
	Status string `json:"status"`
	Data   []struct {
		ContractCode        string  `json:"contract_code"`
		InstrumentIndexCode string  `json:"instrument_index_code"`
		RealTimeSettlement  int     `json:"real_time_settlement"`
		TransferProfitRatio float64 `json:"transfer_profit_ratio"`
		MinLevel            string  `json:"min_level"`
		MaxLevel            string  `json:"max_level"`
		OpenOrderLimit      string  `json:"open_order_limit"`
		OffsetOrderLimit    string  `json:"offset_order_limit"`
		LongPositionLimit   string  `json:"long_position_limit"`
		ShortPositionLimit  string  `json:"short_position_limit"`
		PriceTick           string  `json:"price_tick"`
		InstrumentValue     string  `json:"instrument_value"`
		SettlePeriod        int     `json:"settle_period"`
		FundingRateCap      string  `json:"funding_rate_cap"`
		FundingRateFloor    string  `json:"funding_rate_floor"`
		HighNormalLimit     string  `json:"high_normal_limit"`
		MinNormalLimit      string  `json:"min_normal_limit"`
		HighOpenLimit       string  `json:"high_open_limit"`
		MinOpenLimit        string  `json:"min_open_limit"`
		HighTradeLimit      string  `json:"high_trade_limit"`
		MinTradeLimit       string  `json:"min_trade_limit"`
		ContractInfos       []struct {
			ContractCode   string `json:"contract_code"`
			InstrumentType []int  `json:"instrument_type"`
			SettlementDate string `json:"settlement_date"`
			CreateDate     string `json:"create_date"`
			ContractStatus int    `json:"contract_status"`
		} `json:"contract_infos,omitempty"`
	} `json:"data,omitempty"`
	Ts int64 `json:"ts"`
}
