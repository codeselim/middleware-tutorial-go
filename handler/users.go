package handler

import (
	"fmt"
	"github.com/codeselim/middleware-tutorial-go/common"
	"github.com/codeselim/middleware-tutorial-go/connection"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type UsersHandler struct {
	UserFacade connection.UserConnection
}

func NewUsersHandler() UsersHandler {
	return UsersHandler{
		UserFacade: connection.NewUserFacade(),
	}
}

// handler function has the handler type signature
func (uh UsersHandler) SayHello (w http.ResponseWriter, r *http.Request) error {
	timeNow := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, "Hello, now it is the %s", timeNow)
	return nil
}

func (uh UsersHandler) GetUsers (w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Received a get users request")
	response, err := uh.UserFacade.GetUsers()
	if err != nil {
		return err
	}
	w.Write([]byte(response)) //http status code defaults to 200
	return nil
}

func (uh UsersHandler) GetUserById(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Received a get users by id request")

	userId, err := getValidId(r)
	if err != nil {
		return err
	}

	response, err := uh.UserFacade.GetUserById(userId)
	if err != nil {
		return err
	}

	w.Write([]byte(response)) //http status code defaults to 200
	return nil
}

func getValidId(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	userId := vars["Id"]
	userIdPresent := len(userId) > 0
	if !userIdPresent {
		return "", common.BadRequest("Bad request", "The Id of the user should be supplied") // custom errors
	}
	return userId, nil
}
