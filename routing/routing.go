package routing

import (
	"github.com/codeselim/middleware-tutorial-go/connection"
	"github.com/codeselim/middleware-tutorial-go/handler"
	"github.com/codeselim/middleware-tutorial-go/mapper"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	connection := connection.NewUserFacade()
	mapper := mapper.NewUserMapper()
	usersHandler := handler.NewUsersHandler(connection, mapper)
	r := mux.NewRouter()
	r.Handle("/", handler.NewHandlerWrapper(usersHandler.SayHello))
	r.Handle("/usersWithStations", handler.NewHandlerWrapper(usersHandler.GetUsersWithPhotos))
	r.Handle("/users", handler.NewHandlerWrapper(usersHandler.GetUsers))
	r.Handle("/users/", handler.NewHandlerWrapper(usersHandler.GetUserById))
	r.Handle("/users/{Id}", handler.NewHandlerWrapper(usersHandler.GetUserById))
	return r
}
