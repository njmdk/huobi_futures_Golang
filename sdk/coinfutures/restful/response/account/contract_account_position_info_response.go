package account

type ContractAccountPositionInfoResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts,omitempty"`
	Data   []struct {
		Symbol            string  `json:"symbol,omitempty"`
		MarginBalance     float64 `json:"margin_balance,omitempty"`
		MarginPosition    float64 `json:"margin_position,omitempty"`
		MarginFrozen      float64 `json:"margin_frozen,omitempty"`
		MarginAvailable   float64 `json:"margin_available,omitempty"`
		ProfitReal        float64 `json:"profit_real,omitempty"`
		ProfitUnreal      float64 `json:"profit_unreal,omitempty"`
		RiskRate          float64 `json:"risk_rate,omitempty"`
		WithdrawAvailable float64 `json:"withdraw_available,omitempty"`
		LiquidationPrice  float64 `json:"liquidation_price,omitempty"`
		LeverRate         int     `json:"lever_rate,omitempty"`
		AdjustFactor      float64 `json:"adjust_factor,omitempty"`
		MarginStatic      float64 `json:"margin_static,omitempty"`
		NewRiskRate       string  `json:"new_risk_rate,omitempty"`
		TradePartition    string  `json:"trade_partition,omitempty"`
		Positions         []struct {
			Symbol         string  `json:"symbol,omitempty"`
			ContractCode   string  `json:"contract_code,omitempty"`
			ContractType   string  `json:"contract_type,omitempty"`
			Volume         float64 `json:"volume,omitempty"`
			Available      float64 `json:"available,omitempty"`
			Frozen         float64 `json:"frozen,omitempty"`
			CostOpen       float64 `json:"cost_open,omitempty"`
			CostHold       float64 `json:"cost_hold,omitempty"`
			ProfitUnreal   float64 `json:"profit_unreal,omitempty"`
			ProfitRate     float64 `json:"profit_rate,omitempty"`
			Profit         float64 `json:"profit,omitempty"`
			PositionMargin float64 `json:"position_margin,omitempty"`
			LeverRate      int     `json:"lever_rate,omitempty"`
			Direction      string  `json:"direction,omitempty"`
			LastPrice      float64 `json:"last_price,omitempty"`
			AdlRiskPercent string  `json:"adl_risk_percent,omitempty"`
		} `json:"positions,omitempty"`
	} `json:"data,omitempty"`
}
