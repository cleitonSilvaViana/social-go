package fail

import "net/http"

// ResponseError is the strucutre that is used to return an error to client.
type ResponseError struct {
	StatusCode int    `json:"statusCode,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (r *ResponseError) Error() string {
	return r.Message
}

var INTERNAL_SERVER_ERROR = &ResponseError{
	StatusCode: http.StatusInternalServerError,
	Message:    "internal server error",
}
