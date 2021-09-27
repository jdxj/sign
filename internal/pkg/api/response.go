package api

type RawResponse struct {
	ErrorCode        int         `json:"error_code"`
	ErrorDescription string      `json:"error_description"`
	Response         interface{} `json:"response"`
}
