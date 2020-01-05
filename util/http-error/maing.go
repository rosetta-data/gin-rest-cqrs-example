package error

// HTTPError error implementation
type HTTPError struct {
	Code    int
	Message string
}

// New create HTTPError instance
func New() *HTTPError {
	return &HTTPError{}
}

// Error return http 500 internal server error data
func (err *HTTPError) Error() *HTTPError {
	return &HTTPError{Code: 500, Message: "Internal server error"}
}

// BadRequest return http 400 bad request error data
func (err *HTTPError) BadRequest() *HTTPError {
	return &HTTPError{Code: 400, Message: "Bad request"}
}

// Unauthorized return http 401 unauthorized error data
func (err *HTTPError) Unauthorized() *HTTPError {
	return &HTTPError{Code: 401, Message: "Unauthorized"}
}

// Forbidden return http 403 forbidden error data
func (err *HTTPError) Forbidden() *HTTPError {
	return &HTTPError{Code: 401, Message: "Forbidden"}
}

// NotFound return http 404 not found error data
func (err *HTTPError) NotFound() *HTTPError {
	return &HTTPError{Code: 404, Message: "Not found"}
}

// Conflict return http 209 conflict error data
func (err *HTTPError) Conflict() *HTTPError {
	return &HTTPError{Code: 409, Message: "Conflict"}
}
