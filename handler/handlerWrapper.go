package handler

import (
	"github.com/codeselim/middleware-tutorial-go/common"
	"net/http"
)

type HandlerWrapperFunc func(w http.ResponseWriter, r *http.Request) error
type HandlerWrapper struct {
	// Customize your HandlerWrapper, add any need property...
	//Error error
	//Env string
	Handler HandlerWrapperFunc
}

func NewHandlerWrapper(h HandlerWrapperFunc) HandlerWrapper {
	return HandlerWrapper{
		Handler: h,
	}
}

// ServeHTTP allows our HandlerWrapper type to satisfy http.Handler.
func (hw HandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//execute the wrapped handler
	err := hw.Handler(w, r)

	if err != nil {
		var statusCode int
		var errorResponse string

		switch e := err.(type) {
		case *common.HttpError:
			statusCode = e.StatusCode
			errorResponse = e.Error()
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500 - Internal Server Error
			fallbackError := &common.HttpError{
				StatusCode: http.StatusInternalServerError,
				StatusText: http.StatusText(http.StatusInternalServerError),
			}
			statusCode = fallbackError.StatusCode
			errorResponse = fallbackError.Error()
		}

		w.WriteHeader(statusCode)
		w.Write([]byte(errorResponse))
	}
}
