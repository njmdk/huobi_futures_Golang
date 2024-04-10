package market

type GetHeartbeatResponse struct {
	Status string            `json:"status"`
	Data   HeartbeatResponse `json:"data"`
}

type HeartbeatResponse struct {
	Heartbeat                       int   `json:"heartbeat"`
	SwapHeartbeat                   int   `json:"swap_heartbeat"`
	EstimatedRecoveryTime           int64 `json:"estimated_recovery_time"`
	SwapEstimatedRecoveryTime       int64 `json:"swap_estimated_recovery_time"`
	LinearSwapHeartbeat             int64 `json:"linear_swap_heartbeat"`
	LinearSwapEstimatedRecoveryTime int64 `json:"linear_swap_estimated_recovery_time"`
}
