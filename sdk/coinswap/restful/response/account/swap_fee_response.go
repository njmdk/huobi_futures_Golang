package account

type SwapFeeResponse struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"ts"`
	Data      []struct {
		Symbol        string `json:"symbol"`
		ContractCode  string `json:"contract_code"`
		FeeAsset      string `json:"fee_asset"`
		OpenMakerFee  string `json:"open_maker_fee"`
		OpenTakerFee  string `json:"open_taker_fee"`
		CloseMakerFee string `json:"close_maker_fee"`
		CloseTakerFee string `json:"close_taker_fee"`
	} `json:"data"`
}
