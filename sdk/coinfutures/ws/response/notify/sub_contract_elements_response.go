package notify

type SubContractElementsResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

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
			InstrumentType int    `json:"instrument_type"`
			DeliveryTime   string `json:"delivery_time"`
			CreateDate     string `json:"create_date"`
			ContractStatus int    `json:"contract_status"`
			DeliveryDate   string `json:"delivery_date"`
		} `json:"contract_infos"`
		OpenOrderLimit          int   `json:"open_order_limit"`
		OffsetOrderLimit        int   `json:"offset_order_limit"`
		LongPositionLimit       int   `json:"long_position_limit"`
		ShortPositionLimit      int   `json:"short_position_limit"`
		WeekHigNormalLimit      int   `json:"week_hig_normal_limit"`
		WeekMinNormalLimit      int   `json:"week_min_normal_limit"`
		WeekHigOpenLimit        int   `json:"week_hig_open_limit"`
		WeekMinOpenLimit        int   `json:"week_min_open_limit"`
		WeekHigTradeLimit       int   `json:"week_hig_trade_limit"`
		WeekMinTradeLimit       int   `json:"week_min_trade_limit"`
		BiweekHigNormalLimit    int   `json:"biweek_hig_normal_limit"`
		BiweekMinNormalLimit    int   `json:"biweek_min_normal_limit"`
		BiweekHigOpenLimit      int   `json:"biweek_hig_open_limit"`
		BiweekMinOpenLimit      int   `json:"biweek_min_open_limit"`
		BiweekHigTradeLimit     int   `json:"biweek_hig_trade_limit"`
		BiweekMinTradeLimit     int   `json:"biweek_min_trade_limit"`
		QuarterHigNormalLimit   int   `json:"quarter_hig_normal_limit"`
		QuarterMinNormalLimit   int   `json:"quarter_min_normal_limit"`
		QuarterHigOpenLimit     int   `json:"quarter_hig_open_limit"`
		QuarterMinOpenLimit     int   `json:"quarter_min_open_limit"`
		QuarterHigTradeLimit    int   `json:"quarter_hig_trade_limit"`
		QuarterMinTradeLimit    int   `json:"quarter_min_trade_limit"`
		BiquarterHigNormalLimit int   `json:"biquarter_hig_normal_limit"`
		BiquarterMinNormalLimit int   `json:"biquarter_min_normal_limit"`
		BiquarterHigOpenLimit   int   `json:"biquarter_hig_open_limit"`
		BiquarterMinOpenLimit   int   `json:"biquarter_min_open_limit"`
		BiquarterHigTradeLimit  int   `json:"biquarter_hig_trade_limit"`
		BiquarterMinTradeLimit  int   `json:"biquarter_min_trade_limit"`
		InstrumentType          []int `json:"instrument_type"`
		OrderLimits             []struct {
			InstrumentType int    `json:"instrument_type"`
			Open           string `json:"open"`
			Close          string `json:"close"`
		} `json:"order_limits"`
	} `json:"data"`
}
