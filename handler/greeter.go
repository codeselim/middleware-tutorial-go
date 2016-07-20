package handler

import (
	"fmt"
	"github.com/codeselim/middleware-tutorial-go/common"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	usersApiUrl = "http://jsonplaceholder.typicode.com/users"
)

// handler function has the handler type signature
func SayHello(w http.ResponseWriter, r *http.Request) error {
	timeNow := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, "Hello, now it is the %s", timeNow)
	return nil
}

func GetUsers(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Received a get users request")
	response, err := getUsers(usersApiUrl)
	if err != nil {
		return err
	}
	w.Write([]byte(response)) //http status code defaults to 200
	return nil
}

func GetUserById(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Received a get users by id request")

	userId, err := getValidId(r)
	if err != nil {
		return err
	}

	requestUrl := fmt.Sprintf("%s/%s", usersApiUrl, userId)
	response, err := getUsers(requestUrl)
	if err != nil {
		return err
	}

	w.Write([]byte(response)) //http status code defaults to 200
	return nil
}

func getUsers(requestUrl string) (string, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Fatal("Couldn't create request", err.Error())
		return "", err
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Couldn't send http client request", err.Error())
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Couldn't read response body c", err.Error())
		return "", err
	}

	bodyStringified := string(body)
	return bodyStringified, err
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
