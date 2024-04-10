package account

type ContractMasterSubTransferRecordResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		TransferRecords []struct {
			ID             int64   `json:"id"`
			Ts             int64   `json:"ts"`
			Symbol         string  `json:"symbol"`
			SubUID         string  `json:"sub_uid"`
			SubAccountName string  `json:"sub_account_name"`
			TransferType   int     `json:"transfer_type"`
			Amount         float32 `json:"amount"`
		} `json:"transfer_record"`
		TotalPage   int `json:"total_page"`
		CurrentPage int `json:"current_page"`
		TotalSize   int `json:"total_size"`
	} `json:"data"`
}
