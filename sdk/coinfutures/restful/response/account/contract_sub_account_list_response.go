package account

type ContractSubAccountListResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		SubUID     int64 `json:"sub_uid,omitempty"`
		SymbolList []struct {
			Symbol           string  `json:"symbol"`
			MarginBalance    float64 `json:"margin_balance"`
			LiquidationPrice float64 `json:"liquidation_price"`
			RiskRate         float64 `json:"risk_rate"`
			QueryID          int64   `json:"query_id"`
		} `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
