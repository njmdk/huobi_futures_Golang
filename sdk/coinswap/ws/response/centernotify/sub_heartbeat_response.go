package centernotify

type SubHeartbeatResponse struct {
	Op    string `json:"op"`
	Topic string `json:"topic"`
	Event string `json:"event"`
	Ts    int64  `json:"ts"`
	Data  struct {
		Heartbeat             int   `json:"heartbeat"`
		EstimatedRecoveryTime int64 `json:"estimated_recovery_time"`
	} `json:"data"`
}
