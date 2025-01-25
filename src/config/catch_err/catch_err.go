package catch_err

import "net/http"

type MyError struct {
	Message string  `json:"message"`
	Err     string  `json:"err"`
	Code    int     `json:"code"`
	Cause   []Cause `json:"cause"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *MyError) Error() string {
	return r.Message
}

func NewError(message, err string, code int, cause []Cause) *MyError {
	return &MyError{Message: message, Err: err, Code: code, Cause: cause}
}

func NewForbiddenError(message string) *MyError {
	return &MyError{Message: message, Err: "forbidden_error", Code: http.StatusBadRequest}
}

func NewUnauthorizedError(message string) *MyError {
	return &MyError{Message: message, Err: "unauthorized_error", Code: http.StatusBadRequest}
}

func NewNotFoundError(message string) *MyError {
	return &MyError{Message: message, Err: "content_not_found_error", Code: http.StatusBadRequest}
}
func NewBadRequestError(message string) *MyError {
	return &MyError{Message: message, Err: "bad_request_error", Code: http.StatusBadRequest}
}
func NewInternalServerError(message string) *MyError {
	return &MyError{Message: message, Err: "internal_server_error", Code: http.StatusBadRequest}
}
