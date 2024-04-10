package account

type SwapUserSettlementRecordsResponse struct {
	Status string `json:"status"`
	Data   struct {
		TotalPage         int `json:"total_page"`
		CurrentPage       int `json:"current_page"`
		TotalSize         int `json:"total_size"`
		SettlementRecords []struct {
			Symbol               string  `json:"symbol"`
			ContractCode         string  `json:"contract_code"`
			MarginBalanceInit    float64 `json:"margin_balance_init"`
			MarginBalance        float64 `json:"margin_balance"`
			SettlementProfitReal float64 `json:"settlement_profit_real"`
			SettlementTime       int64   `json:"settlement_time"`
			Clawback             float64 `json:"clawback"`
			FundingFee           float64 `json:"funding_fee"`
			OffsetProfitLoss     float64 `json:"offset_profitloss"`
			Fee                  float64 `json:"fee"`
			FeeAsset             string  `json:"fee_asset"`
			Positions            []struct {
				Symbol                 string  `json:"symbol"`
				ContractCode           string  `json:"contract_code"`
				Direction              string  `json:"direction"`
				Volume                 float64 `json:"volume"`
				CostOpen               float64 `json:"cost_open"`
				CostHoldPre            float64 `json:"cost_hold_pre"`
				CostHold               float64 `json:"cost_hold"`
				SettlementProfitUnreal float64 `json:"settlement_profit_unreal"`
				SettlementPrice        float64 `json:"settlement_price"`
				SettlementType         string  `json:"settlement_type"`
			} `json:"positions"`
		} `json:"settlement_records"`
	} `json:"data"`
	Timestamp int64 `json:"ts"`
}
