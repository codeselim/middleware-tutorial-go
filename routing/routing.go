package routing

import (
	"github.com/codeselim/middleware-tutorial-go/handler"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.SayHello)
	return r
}
