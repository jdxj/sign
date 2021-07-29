package common

type AddTaskReq struct {
	ID     string `json:"id"`
	Domain int    `json:"domain"`
	Type   []int  `json:"type"`
	Key    string `json:"key"`
}

type AddTaskResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
