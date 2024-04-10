package account

type SwapSubAuthResponse struct {
	Status string `json:"status"`
	Data   struct {
		Errors []struct {
			SubUID  string `json:"sub_uid"`
			ErrCode int    `json:"err_code"`
			ErrMsg  string `json:"err_msg"`
		} `json:"errors,omitempty"`
		Successes string `json:"successes,omitempty"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
