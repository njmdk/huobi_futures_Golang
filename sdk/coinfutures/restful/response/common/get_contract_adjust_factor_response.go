package common

type GetContractAdjustFactorResponse struct {
	Status string                 `json:"status"`
	Ts     int64                  `json:"ts"`
	Data   []ContractAdjustFactor `json:"data,omitempty"`
}

type ContractAdjustFactor struct {
	Symbol string `json:"symbol"`
	List   []List `json:"list,omitempty"`
}

type List struct {
	LeverRate float32  `json:"lever_rate"`
	Ladders   []Ladder `json:"ladders,omitempty"`
}

type Ladder struct {
	MinSize      float32 `json:"min_size"`
	MaxSize      float32 `json:"max_size"`
	Ladder       int     `json:"ladder"`
	AdjustFactor float32 `json:"adjust_factor"`
}
