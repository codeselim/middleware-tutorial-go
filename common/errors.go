package common

import (
	"encoding/json"
	"net/http"
)

// Errors provide a way to return detailed information
// for an http request error. The error is normally
// JSON encoded.
type HttpError struct {
	StatusCode int    `json:"-"`
	StatusText string `json:"statusText"`
	ErrorCode  string `json:"errorCode,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (err *HttpError) Error() string {
	b, _ := json.Marshal(err)
	return string(b)
}

func ToHTTPError(err error) *HttpError {
	switch e := err.(type) {
	case *HttpError:
		return e
	default:
		// Any error types we don't specifically look out for default
		// to serving a HTTP 500 - Internal Server Error
		return &HttpError{
			StatusCode: http.StatusInternalServerError,
			StatusText: http.StatusText(http.StatusInternalServerError),
		}
	}
}

func NotFound(errorCode string, message string) error {
	return &HttpError{
		StatusCode: http.StatusNotFound,
		StatusText: http.StatusText(http.StatusNotFound),
		ErrorCode:  errorCode,
		Message:    message,
	}
}

func BadRequest(errorCode string, message string) error {
	return &HttpError{
		StatusCode: http.StatusBadRequest,
		StatusText: http.StatusText(http.StatusBadRequest),
		ErrorCode:  errorCode,
		Message:    message,
	}
}

func InternalServerError(errorCode string, message string) error {
	return &HttpError{
		StatusCode: http.StatusInternalServerError,
		StatusText: http.StatusText(http.StatusInternalServerError),
		ErrorCode:  errorCode,
		Message:    message,
	}
}

func New(statusCode int, errorCode string, message string) error {
	return &HttpError{
		StatusCode: statusCode,
		StatusText: http.StatusText(statusCode),
		ErrorCode:  errorCode,
		Message:    message,
	}
}
