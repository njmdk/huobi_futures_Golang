package unifiedaccount

type SwapUnifiedAccountTypeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data struct {
		AccountType int `json:"account_type"`
	} `json:"data"`
}
