package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func SuccessResponse(status int, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   nil,
	}
}

func ErrorResponse(status int, message string, err interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    nil,
		Error:   err,
	}
}
