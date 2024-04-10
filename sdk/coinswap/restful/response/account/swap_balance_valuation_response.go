package account

type SwapBalanceValuationResponse struct {
	Status string `json:"status"`
	Data   []struct {
		ValuationAsset string `json:"valuation_asset"`
		Balance        string `json:"balance"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
