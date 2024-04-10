package common

type GetContractInsuranceFundResponse struct {
	Status string                `json:"status"`
	Ts     int64                 `json:"ts"`
	Data   ContractInsuranceFund `json:"data,omitempty"`
}

type ContractInsuranceFund struct {
	Symbol string `json:"symbol"`
	Tick   []Tick `json:"tick,omitempty"`
}

type Tick struct {
	InsuranceFund float64 `json:"insurance_fund"`
	Ts            int64   `json:"ts"`
}
