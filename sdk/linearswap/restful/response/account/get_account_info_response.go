package account

type GetAccountInfoResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		MarginAsset string `json:"margin_asset"`

		MarginBalance float32 `json:"margin_balance"`

		MarginStatic float32 `json:"margin_static"`

		MarginPosition float32 `json:"margin_position"`

		MarginFrozen float32 `json:"margin_frozen"`

		MarginAvailable float32 `json:"margin_available"`

		ProfitReal float32 `json:"profit_real"`

		ProfitUnreal float32 `json:"profit_unreal"`

		WithdrawAvailable float32 `json:"withdraw_available,omitempty"`

		RiskRate float32 `json:"risk_rate,omitempty"`

		NewRiskRate string `json:"new_risk_rate,omitempty"`

		MoneyIn float32 `json:"money_in,omitempty"`

		MoneyOut float32 `json:"money_out,omitempty"`

		TradePartition string `json:"trade_partition,omitempty"`

		LiquidationPrice float32 `json:"liquidation_price,omitempty"`

		LeverRate float32 `json:"lever_rate"`

		AdjustFactor float32 `json:"adjust_factor"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`

		PositionMode string `json:"position_mode"`

		ContractDetail []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			MarginPosition float32 `json:"margin_position"`

			MarginFrozen float32 `json:"margin_frozen"`

			MarginAvailable float32 `json:"margin_available"`

			ProfitUnreal float32 `json:"profit_unreal"`

			LiquidationPrice float32 `json:"liquidation_price"`

			LeverRate float32 `json:"lever_rate"`

			AdjustFactor float32 `json:"adjust_factor"`

			Pair string `json:"pair"`

			BusinessType string `json:"business_type"`

			CrossMaxAvailable string `json:"cross_max_available"`

			ContractType string `json:"contract_type"`

			TradePartition float32 `json:"trade_partition,omitempty"`
		} `json:"contract_detail,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
