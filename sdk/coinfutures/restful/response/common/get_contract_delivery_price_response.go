package common

type GetContractDeliveryPriceResponse struct {
	Status string `json:"status"`
	Data   struct {
		DeliveryPrice float64 `json:"delivery_price"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
