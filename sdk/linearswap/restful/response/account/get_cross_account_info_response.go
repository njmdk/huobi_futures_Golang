package account

type GetCrossAccountInfoResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		MarginMode        string            `json:"margin_mode"`
		MarginAccount     string            `json:"margin_account"`
		MarginAsset       string            `json:"margin_asset"`
		MarginBalance     float32           `json:"margin_balance"`
		MarginStatic      float32           `json:"margin_static"`
		MarginPosition    float32           `json:"margin_position"`
		MarginFrozen      float32           `json:"margin_frozen"`
		ProfitUnreal      float32           `json:"profit_unreal"`
		WithdrawAvailable float32           `json:"withdraw_available"`
		RiskRate          float32           `json:"risk_rate"`
		MoneyIn           float32           `json:"money_in"`
		MoneyOut          float32           `json:"money_out"`
		NewRiskRate       float32           `json:"new_risk_rate"`
		PositionMode      string            `json:"position_mode"`
		ContractDetail    []ContractDetail  `json:"contract_detail"`
		FuturesDetail     []FuturesContract `json:"futures_contract_detail"`
	} `json:"data"`
}

type ContractDetail struct {
	Symbol            string  `json:"symbol"`
	ContractCode      string  `json:"contract_code"`
	MarginPosition    float32 `json:"margin_position"`
	MarginFrozen      float32 `json:"margin_frozen"`
	MarginAvailable   float32 `json:"margin_available"`
	ProfitUnreal      float32 `json:"profit_unreal"`
	LiquidationPrice  float32 `json:"liquidation_price"`
	LeverRate         float32 `json:"lever_rate"`
	AdjustFactor      float32 `json:"adjust_factor"`
	ContractType      string  `json:"contract_type"`
	CrossMaxAvailable float32 `json:"cross_max_available"`
	TradePartition    string  `json:"trade_partition"`
	Pair              string  `json:"pair"`
	BusinessType      string  `json:"business_type"`
}

type FuturesContract struct {
	Symbol            string  `json:"symbol"`
	ContractCode      string  `json:"contract_code"`
	MarginPosition    float32 `json:"margin_position"`
	MarginFrozen      float32 `json:"margin_frozen"`
	MarginAvailable   float32 `json:"margin_available"`
	ProfitUnreal      float32 `json:"profit_unreal"`
	LiquidationPrice  float32 `json:"liquidation_price"`
	LeverRate         float32 `json:"lever_rate"`
	AdjustFactor      float32 `json:"adjust_factor"`
	ContractType      string  `json:"contract_type"`
	CrossMaxAvailable float32 `json:"cross_max_available"`
	TradePartition    string  `json:"trade_partition"`
	Pair              string  `json:"pair"`
	BusinessType      string  `json:"business_type"`
}
