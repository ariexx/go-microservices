package helper

type ResponseSuccess struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseError struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func ResponseSuccessHandler(message string, data interface{}) *ResponseSuccess {
	return &ResponseSuccess{
		Error:   false,
		Message: message,
		Data:    data,
	}
}

func ResponseErrorHandler(message string) *ResponseError {
	return &ResponseError{
		Error:   true,
		Message: message,
	}
}
