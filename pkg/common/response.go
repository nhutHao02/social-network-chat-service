package common

type Response struct {
	Success bool          `json:"success"`
	Data    interface{}   `json:"data"`
	Error   *ErrorMessage `json:"error"`
}

type PagingSuccessResponse struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	TotalPage uint64      `json:"totalPage"`
}

type ErrorMessage struct {
	Errors  string `json:"errors"`
	Message string `json:"message"`
}

func NewSuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Data:    &data,
		Error:   nil,
	}
}

func NewErrorResponse(err string, message string) Response {
	return Response{
		Success: false,
		Data:    nil,
		Error: &ErrorMessage{
			Errors:  err,
			Message: message,
		},
	}
}

func NewPagingSuccessResponse(data interface{}, totalPage uint64) PagingSuccessResponse {
	return PagingSuccessResponse{
		Success:   true,
		Data:      &data,
		TotalPage: totalPage,
	}
}
