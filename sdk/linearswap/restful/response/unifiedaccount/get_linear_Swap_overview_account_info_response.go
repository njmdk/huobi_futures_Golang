package unifiedaccount

type GetLinearSwapOverviewAccountInfoResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data []struct {
		MarginAsset     string  `json:"margin_asset"`
		MarginBalance   float64 `json:"margin_balance"`
		MarginAvailable float64 `json:"margin_available"`
	} `json:"data"`
}
