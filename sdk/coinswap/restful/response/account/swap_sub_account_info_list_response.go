package account

type SwapSubAccountInfoListResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		SubList []struct {
			SubUID          int64 `json:"sub_uid"`
			AccountInfoList []struct {
				Symbol           string  `json:"symbol"`
				ContractCode     string  `json:"contract_code"`
				MarginBalance    float64 `json:"margin_balance"`
				LiquidationPrice float64 `json:"liquidation_price"`
				RiskRate         float64 `json:"risk_rate"`
			} `json:"account_info_list,omitempty"`
		} `json:"sub_list,omitempty"`
		CurrentPage int `json:"current_page"`
		TotalPage   int `json:"total_page"`
		TotalSize   int `json:"total_size"`
	} `json:"data,omitempty"`
}
