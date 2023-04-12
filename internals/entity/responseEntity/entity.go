package responseEntity

type LoginRes struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Token   *string `json:"token"`
	Data    any     `json:"data"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *any   `json:"data"`
}

func LoginResponse(code int, message string, token *string, data any) *LoginRes {
	return &LoginRes{Code: code, Message: message, Token: token, Data: data}
}

func SuccessResponse(code int, message string, data *any) *Response {
	return &Response{Code: code, Message: message, Data: data}
}


func ErrorResponse(code int, message string, data *any) *Response {
	return &Response{Code: code, Message: message, Data: data}
}
