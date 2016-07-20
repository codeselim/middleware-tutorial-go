package routing

import (
	"github.com/codeselim/middleware-tutorial-go/handler"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", handler.NewHandlerWrapper(handler.SayHello))
	r.Handle("/users", handler.NewHandlerWrapper(handler.GetUsers))
	r.Handle("/users/", handler.NewHandlerWrapper(handler.GetUserById))
	r.Handle("/users/{Id}", handler.NewHandlerWrapper(handler.GetUserById))
	return r
}
