package routing

import (
	"github.com/codeselim/middleware-tutorial-go/handler"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.SayHello)
	r.HandleFunc("/users", handler.GetUsers)
	r.HandleFunc("/users/", handler.GetUsers)
	r.HandleFunc("/users/{Id}", handler.GetUserById)
	return r
}
