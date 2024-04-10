package market

type GetHistoryBasisResponse struct {
	Ch     string `json:"ch"`
	Status string `json:"status"`
	Data   []struct {
		ID            int64  `json:"id"`
		ContractPrice string `json:"contract_price"`
		IndexPrice    string `json:"index_price"`
		Basis         string `json:"basis"`
		BasisRate     string `json:"basis_rate"`
	} `json:"data"`
	TS int64 `json:"ts"`
}
