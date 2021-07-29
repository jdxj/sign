package common

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AddTaskReq struct {
	ID     string `json:"id"`
	Domain int    `json:"domain"`
	Type   []int  `json:"type"`
	Key    string `json:"key"`
}
