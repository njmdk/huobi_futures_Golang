package market

type GetTimestampResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
}
