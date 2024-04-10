package order

type ContractOrderRequest struct {
	Symbol string `json:"symbol"`

	ContractCode string `json:"contract_code"`

	ReduceOnly string `json:"reduce_only"`

	Pair string `json:"pair"`

	ContractType string `json:"contract_type"`

	ClientOrderId int64 `json:"client_order_id,omitempty"`

	Price float32 `json:"price,omitempty"`

	Volume int64 `json:"volume"`

	Direction string `json:"direction"`

	Offset string `json:"offset"`

	LeverRate int `json:"lever_rate"`

	OrderPriceType string `json:"order_price_type"`

	TpTriggerPrice float32 `json:"tp_trigger_price,omitempty"`

	TpOrderPrice float32 `json:"tp_order_price,omitempty"`

	TpOrderPriceType string `json:"tp_order_price_type,omitempty"`

	SlTriggerPrice float32 `json:"sl_trigger_price,omitempty"`

	SlOrderPrice float32 `json:"sl_order_price,omitempty"`

	SlOrderPriceType string `json:"sl_order_price_type,omitempty"`
}

type BatchContractOrderRequest []ContractOrderRequest
