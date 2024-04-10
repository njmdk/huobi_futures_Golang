package unifiedaccount

type GetUnifiedAccountInfoResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data []struct {
		MarginBalance     float64 `json:"margin_balance"`
		MarginStatic      float64 `json:"margin_static"`
		CrossProfitUnreal float64 `json:"cross_profit_unreal"`
		CrossMarginStatic float64 `json:"cross_margin_static"`
		MarginAsset       string  `json:"margin_asset"`
		MarginFrozen      float64 `json:"margin_frozen"`
		WithdrawAvailable float64 `json:"withdraw_available"`
		CrossRiskRate     float64 `json:"cross_risk_rate"`
		CrossSwap         []struct {
			Symbol            string  `json:"symbol"`
			ContractCode      string  `json:"contract_code"`
			MarginMode        string  `json:"margin_mode"`
			MarginAvailable   float64 `json:"margin_available"`
			CrossMaxAvailable int     `json:"cross_max_available"`
			LeverRate         float64 `json:"lever_rate"`
			ContractType      string  `json:"contract_type"`
			BusinessType      string  `json:"business_type"`
		} `json:"cross_swap"`
		CrossFutures []struct {
			Symbol          string  `json:"symbol"`
			ContractCode    string  `json:"contract_code"`
			MarginMode      string  `json:"margin_mode"`
			MarginAvailable float64 `json:"margin_available"`
			LeverRate       float64 `json:"lever_rate"`
			ContractType    string  `json:"contract_type"`
			BusinessType    string  `json:"business_type"`
		} `json:"cross_futures"`
		IsolatedSwap []struct {
			Symbol            string  `json:"symbol"`
			ContractCode      string  `json:"contract_code"`
			MarginMode        string  `json:"margin_mode"`
			MarginAvailable   float64 `json:"margin_available"`
			WithdrawAvailable float64 `json:"withdraw_available"`
			LeverRate         int     `json:"lever_rate"`
			PositionMode      string  `json:"position_mode"`
		} `json:"isolated_swap"`
	} `json:"data"`
}
