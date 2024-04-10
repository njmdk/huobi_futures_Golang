package common

type GetContractApiStateResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Symbol            string `json:"symbol"`
		Open              int    `json:"open"`
		Close             int    `json:"close"`
		Cancel            int    `json:"cancel"`
		TransferIn        int    `json:"transfer_in"`
		TransferOut       int    `json:"transfer_out"`
		MasterTransferSub int    `json:"master_transfer_sub"`
		SubTransferMaster int    `json:"sub_transfer_master"`
	} `json:"data"`
}
