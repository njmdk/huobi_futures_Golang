package order

type LinearCancelAfterResponse struct {
	Code int32             `json:"code"`
	Msg  string            `json:"msg"`
	Ts   int64             `json:"ts"`
	Data LinearCancelAfter `json:"data"`
}

type LinearCancelAfter struct {
	CurrentTime int64 `json:"current_time"`
	TriggerTime int64 `json:"trigger_time"`
}
