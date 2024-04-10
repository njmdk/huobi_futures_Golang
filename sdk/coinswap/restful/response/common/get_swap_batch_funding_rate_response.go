package common

type GetSwapBatchFundingRateResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Symbol          string `json:"symbol"`
		ContractCode    string `json:"contract_code"`
		FeeAsset        string `json:"fee_asset"`
		FundingTime     string `json:"funding_time"`
		FundingRate     string `json:"funding_rate"`
		EstimatedRate   string `json:"estimated_rate,omitempty"`
		NextFundingTime string `json:"next_funding_time,omitempty"`
	} `json:"data,omitempty"`
}
