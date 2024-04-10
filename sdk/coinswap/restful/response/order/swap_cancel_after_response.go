package order

type SwapCancelAfterResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data struct {
		CurrentTime int64 `json:"current_time"`
		TriggerTime int64 `json:"trigger_time"`
	} `json:"data"`
}
