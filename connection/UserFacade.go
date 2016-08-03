package connection

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	usersApiUrl  = "http://jsonplaceholder.typicode.com/users"
	photosApiUrl = "http://jsonplaceholder.typicode.com/photos"
)

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

type UserFacade struct{}

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

func (uf UserFacade) GetUserWithPhotos() (string, error) {
	endpointsUrls := []string{
		usersApiUrl,
		photosApiUrl,
	}
	responses := fireAsync(endpointsUrls)
	var responsesString string
	for _, response := range responses {
		bodyBytes, _ := ioutil.ReadAll(response.response.Body)
		response.response.Body.Close()
		responsesString += string(bodyBytes)
	}
	return responsesString, nil
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

func fireAsync(endpointsUrls []string) []*HttpResponse {
	ch := make(chan *HttpResponse)
	responses := []*HttpResponse{}
	client := http.Client{}

	for _, endpointUrl := range endpointsUrls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := client.Get(url)
			ch <- &HttpResponse{url, resp, err}
			if err != nil && resp != nil && resp.StatusCode == http.StatusOK {
				resp.Body.Close()
			}
		}(endpointUrl)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			if r.err != nil {
				fmt.Println("with an error", r.err)
			}
			responses = append(responses, r)
			if len(responses) == len(endpointsUrls) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
	return responses
}
