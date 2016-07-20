package connection

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

const (
	usersApiUrl = "http://jsonplaceholder.typicode.com/users"
)

type UserFacade struct {}

func NewUserFacade() UserConnection {
	return &UserFacade{}
}

func (uf UserFacade) GetUsers() (string, error) {
	response, err := sendRequest(usersApiUrl)
	return response, err
}

func (uf UserFacade) GetUserById(userId string) (string, error) {
	requestUrl := fmt.Sprintf("%s/%s", usersApiUrl, userId)
	response, err := sendRequest(requestUrl)
	return response, err
}

func sendRequest(requestUrl string) (string, error) {
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