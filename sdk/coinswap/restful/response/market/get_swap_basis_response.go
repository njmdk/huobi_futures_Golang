package market

type GetSwapBasisResponse struct {
	Ch   string `json:"ch"`
	Data []struct {
		ID            int64  `json:"id"`
		ContractPrice string `json:"contract_price"`
		IndexPrice    string `json:"index_price"`
		Basis         string `json:"basis"`
		BasisRate     string `json:"basis_rate"`
	} `json:"data"`
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
}
