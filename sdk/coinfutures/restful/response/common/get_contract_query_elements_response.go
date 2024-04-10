package common

type GetContractQueryElementsResponse struct {
	Status string `json:"status"`
	Data   []struct {
		ContractCode             string  `json:"contract_code"`
		TransferProfitRatio      float64 `json:"transfer_profit_ratio"`
		MinLevel                 string  `json:"min_level"`
		MaxLevel                 string  `json:"max_level"`
		OpenOrderLimit           string  `json:"open_order_limit"`
		OffsetOrderLimit         string  `json:"offset_order_limit"`
		LongPositionLimit        string  `json:"long_position_limit"`
		ShortPositionLimit       string  `json:"short_position_limit"`
		WeekHighNormalLimit      float64 `json:"week_hig_normal_limit"`
		WeekMinNormalLimit       float64 `json:"week_min_normal_limit"`
		WeekHighOpenLimit        float64 `json:"week_hig_open_limit"`
		WeekMinOpenLimit         float64 `json:"week_min_open_limit"`
		WeekHighTradeLimit       float64 `json:"week_hig_trade_limit"`
		WeekMinTradeLimit        float64 `json:"week_min_trade_limit"`
		BiweekHighNormalLimit    int     `json:"biweek_hig_normal_limit"`
		BiweekMinNormalLimit     int     `json:"biweek_min_normal_limit"`
		BiweekHighOpenLimit      int     `json:"biweek_hig_open_limit"`
		BiweekMinOpenLimit       int     `json:"biweek_min_open_limit"`
		BiweekHighTradeLimit     int     `json:"biweek_hig_trade_limit"`
		BiweekMinTradeLimit      int     `json:"biweek_min_trade_limit"`
		QuarterHighNormalLimit   float64 `json:"quarter_hig_normal_limit"`
		QuarterMinNormalLimit    float64 `json:"quarter_min_normal_limit"`
		QuarterHighOpenLimit     float64 `json:"quarter_hig_open_limit"`
		QuarterMinOpenLimit      float64 `json:"quarter_min_open_limit"`
		QuarterHighTradeLimit    float64 `json:"quarter_hig_trade_limit"`
		QuarterMinTradeLimit     float64 `json:"quarter_min_trade_limit"`
		BiquarterHighNormalLimit float64 `json:"biquarter_hig_normal_limit"`
		BiquarterMinNormalLimit  float64 `json:"biquarter_min_normal_limit"`
		InstrumentIndexCode      string  `json:"instrument_index_code"`
		RealTimeSettlement       int     `json:"real_time_settlement"`
		Ts                       int64   `json:"ts"`
		BiquarterHigOpenLimit    float64 `json:"biquarter_hig_open_limit"`
		BiquarterMinOpenLimit    float64 `json:"biquarter_min_open_limit"`
		BiquarterHigTradeLimit   float64 `json:"biquarter_hig_trade_limit"`
		BiquarterMinTradeLimit   float64 `json:"biquarter_min_trade_limit"`
		ContractInfos            []struct {
			ContractCode   string `json:"contract_code"`
			InstrumentType int    `json:"instrument_type"`
			CreateDate     string `json:"create_date"`
			ContractStatus int    `json:"contract_status"`
			DeliveryTime   string `json:"delivery_time"`
			DeliveryDate   string `json:"delivery_date"`
		} `json:"contract_infos"`
		OrderLimits []struct {
			Open           string `json:"open"`
			Close          string `json:"close"`
			InstrumentType int    `json:"instrument_type"`
		} `json:"order_limits"`
	} `json:"data"`
}
