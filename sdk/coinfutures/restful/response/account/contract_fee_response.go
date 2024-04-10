package account

type ContractFeeResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Symbol        string `json:"symbol"`
		FeeAsset      string `json:"fee_asset"`
		OpenMakerFee  string `json:"open_maker_fee"`
		OpenTakerFee  string `json:"open_taker_fee"`
		CloseMakerFee string `json:"close_maker_fee"`
		CloseTakerFee string `json:"close_taker_fee"`
		DeliveryFee   string `json:"delivery_fee"`
	} `json:"data"`
}
