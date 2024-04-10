package market

type GetSwapQueryElementsResponse struct {
	Status string                      `json:"status"`
	Data   []SwapQueryElementsResponse `json:"data"`
}

type SwapQueryElementsResponse struct {
	ContractCode             string         `json:"contract_code"`
	ModeType                 int            `json:"mode_type"`
	SwapDeliveryType         int            `json:"swap_delivery_type"`
	InstrumentIndexCode      string         `json:"instrument_index_code"`
	RealTimeSettlement       int            `json:"real_time_settlement"`
	TransferProfitRatio      float64        `json:"transfer_profit_ratio"`
	CrossTransferProfitRatio float64        `json:"cross_transfer_profit_ratio"`
	InstrumentType           []int          `json:"instrument_type"`
	TradePartition           string         `json:"trade_partition"`
	MinLevel                 string         `json:"min_level"`
	MaxLevel                 string         `json:"max_level"`
	SettlePeriod             int            `json:"settle_period"`
	FundingRateCap           string         `json:"funding_rate_cap"`
	FundingRateFloor         string         `json:"funding_rate_floor"`
	LongPositionLimit        string         `json:"long_position_limit,omitempty"`
	OffsetOrderLimit         string         `json:"offset_order_limit,omitempty"`
	OpenOrderLimit           string         `json:"open_order_limit,omitempty"`
	ShortPositionLimit       string         `json:"short_position_limit,omitempty"`
	ContractInfos            []ContractInfo `json:"contract_infos"`
	PriceTicks               []PriceTick    `json:"price_ticks,omitempty"`
	InstrumentValues         []Instrument   `json:"instrument_values,omitempty"`
	OrderLimits              []OrderLimit   `json:"order_limits,omitempty"`
	NormalLimits             []NormalLimit  `json:"normal_limits,omitempty"`
	TradeLimits              []TradeLimit   `json:"trade_limits"`
}

type ContractInfo struct {
	ContractCode   string `json:"contract_code"`
	InstrumentType int    `json:"instrument_type"`
	SettlementDate string `json:"settlement_date"`
	DeliveryTime   string `json:"delivery_time"`
	CreateDate     string `json:"create_date"`
	ContractStatus int    `json:"contract_status"`
	DeliveryDate   string `json:"delivery_date"`
}

type PriceTick struct {
	BusinessType int    `json:"business_type"`
	Price        string `json:"price"`
}

type Instrument struct {
	BusinessType int    `json:"business_type"`
	Price        string `json:"price"`
}

type OrderLimit struct {
	InstrumentType   int    `json:"instrument_type"`
	Open             string `json:"open"`
	Close            string `json:"close"`
	OpenAfterClosing string `json:"open_after_closing"`
}

type TradeLimit struct {
	InstrumentType int    `json:"instrument_type"`
	Open           string `json:"open"`
	Close          string `json:"close"`
}

type NormalLimit struct {
	InstrumentType int    `json:"instrument_type"`
	Open           string `json:"open"`
	Close          string `json:"close"`
}
