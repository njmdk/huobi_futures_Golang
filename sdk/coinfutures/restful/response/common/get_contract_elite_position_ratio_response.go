package common

type GetContractElitePositionRatioResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		Symbol string `json:"symbol"`
		List   []struct {
			BuyRatio  float32 `json:"buy_ratio"`
			SellRatio float32 `json:"sell_ratio"`
			Ts        int64   `json:"ts"`
		} `json:"list"`
	} `json:"data"`
}
