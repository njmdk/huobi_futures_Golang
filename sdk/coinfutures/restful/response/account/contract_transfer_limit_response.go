package account

type ContractTransferLimitResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Symbol                 string  `json:"symbol"`
		TransferInMaxEach      float64 `json:"transfer_in_max_each"`
		TransferInMinEach      float64 `json:"transfer_in_min_each"`
		TransferOutMaxEach     float64 `json:"transfer_out_max_each"`
		TransferOutMinEach     float64 `json:"transfer_out_min_each"`
		TransferInMaxDaily     float64 `json:"transfer_in_max_daily"`
		TransferOutMaxDaily    float64 `json:"transfer_out_max_daily"`
		NetTransferInMaxDaily  float64 `json:"net_transfer_in_max_daily"`
		NetTransferOutMaxDaily float64 `json:"net_transfer_out_max_daily"`
	} `json:"data"`
}
