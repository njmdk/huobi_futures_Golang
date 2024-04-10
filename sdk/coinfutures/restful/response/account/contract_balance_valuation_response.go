package account

type ContractBalanceValuationResponse struct {
	Status string `json:"status"`
	Data   []struct {
		ValuationAsset string `json:"valuation_asset"`
		Balance        string `json:"balance"`
	} `json:"data,omitempty"`
	TS int64 `json:"ts"`
}
