package handler

import "net/http"

type HandlerWrapperFunc func(w http.ResponseWriter, r *http.Request)
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
	hw.Handler(w, r)
}
