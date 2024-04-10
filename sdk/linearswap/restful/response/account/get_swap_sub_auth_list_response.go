package account

type GetSwapSubAuthListResponse struct {
	Status string             `json:"status"`
	Data   GetSwapSubAuthList `json:"data"`
	Ts     int64              `json:"ts"`
}

type GetSwapSubAuthList struct {
	QueryID   int64     `json:"query_id"`
	Errors    []Error   `json:"errors"`
	Successes []Success `json:"successes"`
}

type Error struct {
	SubUID       string `json:"sub_uid"`
	ErrorCode    int32  `json:"err_code"`
	ErrorMessage string `json:"err_msg"`
}

type Success struct {
	SubUID  string `json:"sub_uid"`
	SubAuth int32  `json:"sub_auth"`
}
