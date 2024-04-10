package account

type ContractSubPositionInfoResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Symbol         string  `json:"symbol"`
		ContractCode   string  `json:"contract_code"`
		ContractType   string  `json:"contract_type"`
		Volume         float64 `json:"volume"`
		Available      float64 `json:"available"`
		Frozen         float64 `json:"frozen"`
		CostOpen       float64 `json:"cost_open"`
		CostHold       float64 `json:"cost_hold"`
		ProfitUnreal   float64 `json:"profit_unreal"`
		ProfitRate     float64 `json:"profit_rate"`
		Profit         float64 `json:"profit"`
		PositionMargin float64 `json:"position_margin"`
		LeverRate      int     `json:"lever_rate"`
		Direction      string  `json:"direction"`
		LastPrice      float64 `json:"last_price"`
		AdlRiskPercent string  `json:"adl_risk_percent,omitempty"`
		NewRiskRate    string  `json:"new_risk_rate"`
		TradePartition string  `json:"trade_partition"`
	} `json:"data,omitempty"`
}
