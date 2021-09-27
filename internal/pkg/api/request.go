package api

type RawRequest struct {
	Token   string      `json:"token"`
	Request interface{} `json:"request"`
}
