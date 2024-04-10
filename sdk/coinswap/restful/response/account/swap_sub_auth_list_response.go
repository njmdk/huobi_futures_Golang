package account

type SwapSubAuthListResponse struct {
	Status string `json:"status"`
	Data   struct {
		Errors []struct {
			SubUID  string `json:"sub_uid"`
			ErrCode int    `json:"err_code"`
			ErrMsg  string `json:"err_msg"`
		} `json:"errors,omitempty"`
		Successes []struct {
			QueryID int64  `json:"query_id,omitempty"`
			SubUID  string `json:"sub_uid"`
			SubAuth int    `json:"sub_auth"`
		} `json:"successes,omitempty"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
