package routing

import (
	"github.com/codeselim/middleware-tutorial-go/handler"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	usersHandler := handler.NewUsersHandler()
	r := mux.NewRouter()
	r.Handle("/", handler.NewHandlerWrapper(usersHandler.SayHello))
	r.Handle("/users", handler.NewHandlerWrapper(usersHandler.GetUsers))
	r.Handle("/users/", handler.NewHandlerWrapper(usersHandler.GetUserById))
	r.Handle("/users/{Id}", handler.NewHandlerWrapper(usersHandler.GetUserById))
	return r
}
