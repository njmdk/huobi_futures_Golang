package account

type SwapMasterSubTransferRecordResponse struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"ts"`
	Data      struct {
		TransferRecords []struct {
			ID             int64   `json:"id"`
			Timestamp      int64   `json:"ts"`
			Symbol         string  `json:"symbol"`
			ContractCode   string  `json:"contract_code"`
			SubUID         string  `json:"sub_uid"`
			SubAccountName string  `json:"sub_account_name"`
			TransferType   int     `json:"transfer_type"`
			Amount         float64 `json:"amount"`
		} `json:"transfer_record"`
		TotalPage   int `json:"total_page"`
		CurrentPage int `json:"current_page"`
		TotalSize   int `json:"total_size"`
	} `json:"data"`
}
