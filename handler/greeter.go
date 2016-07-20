package handler

import (
	"fmt"
	"net/http"
	"time"
	"log"
	"io/ioutil"
	"github.com/gorilla/mux"
)

const (
	usersApiUrl = "http://jsonplaceholder.tyffpicode.com/users"
)

// handler function has the handler type signature
func SayHello(w http.ResponseWriter, r *http.Request) {
	timeNow := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, "Hello, now it is the %s", timeNow)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a get users request")
	response, _ := getUsers(usersApiUrl)
	w.Write([]byte(response)) //http status code defaults to 200
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a get users by id request")
	vars := mux.Vars(r)
	userId := vars["Id"]
	requestUrl := fmt.Sprintf("%s/%s", usersApiUrl, userId)
	response, _ := getUsers(requestUrl)
	w.Write([]byte(response)) //http status code defaults to 200
}

func getUsers(requestUrl string) (string, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", requestUrl, nil)
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