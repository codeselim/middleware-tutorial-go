package handler

import (
	"encoding/json"
	"fmt"
	"github.com/codeselim/middleware-tutorial-go/common"
	"github.com/codeselim/middleware-tutorial-go/connection"
	"github.com/codeselim/middleware-tutorial-go/contract/usersapi"
	"github.com/codeselim/middleware-tutorial-go/mapper"
	"net/http"
	"time"
)

type UsersHandler struct {
	UserFacade connection.UserConnection
	UserMapper mapper.User
}

func NewUsersHandler(uf connection.UserConnection, um mapper.User) UsersHandler {
	return UsersHandler{
		UserFacade: uf,
		UserMapper: um,
	}
}

// handler function has the handler type signature
func (uh UsersHandler) SayHello(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	timeNow := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, "Hello, now it is the %s", timeNow)
	return nil
}

func (uh UsersHandler) GetUsers(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	fmt.Println("Received a get users request")
	response, err := uh.UserFacade.GetUsers()
	if err != nil {
		return err
	}
	w.Write([]byte(response)) //http status code defaults to 200
	return nil
}

func (uh UsersHandler) GetUsersWithPhotos(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	response, err := uh.UserFacade.GetUserWithPhotos()
	if err != nil {
		return err
	}
	w.Write([]byte(response)) //http status code defaults to 200
	return nil
}

func (uh UsersHandler) GetUserById(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	fmt.Println("Received a get users by id request")

	userId, err := getValidId(vars)
	if err != nil {
		return err
	}

	response, err := uh.UserFacade.GetUserById(userId)
	if err != nil {
		return err
	}

	//Parse JSON data
	data := usersapi.User{}
	b := []byte(response)
	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	//get the domain user object from the remote user object
	domainUser := mapper.NewUserMapper().GetDomainUser(data, "DATA-FROM-DB")

	//marshal the result and send JSON to the client
	marshaledResponse, err := json.Marshal(domainUser)
	if err != nil {
		return err
	}

	w.Write([]byte(marshaledResponse)) //http status code defaults to 200
	return nil
}

func getValidId(vars map[string]string) (string, error) {
	userId := vars["Id"]
	userIdPresent := len(userId) > 0
	if !userIdPresent {
		return "", common.BadRequest(http.StatusText(400), common.UserIdMissing) // custom errors
	}
	return userId, nil
}
