package response

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func ClientResponse(code int, message string, data interface{}, err interface{}) Response {

	return Response{
		Code:    code,
		Message: message,
		Data:    data,
		Error:   err,
	}

}