package handler

import (
	"fmt"
	"github.com/codeselim/middleware-tutorial-go/common"
	"github.com/codeselim/middleware-tutorial-go/contract/api"
	"github.com/codeselim/middleware-tutorial-go/contract/usersapi"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var userFixture = api.User{
	Id:       2,
	Name:     "Ervin Howell",
	Data:     "DATA-FROM-DB",
	Email:    "Shanna@melissa.tv",
	Username: "Antonette",
}

const userJsonOutput = "{\"id\":2,\"name\":\"Ervin Howell\",\"username\":\"Antonette\",\"email\":\"Shanna@melissa.tv\",\"data\":\"DATA-FROM-DB\"}"
const remoteUserJsonOutput = "{\"id\": 2,\"name\": \"Ervin Howell\",\"username\": \"Antonette\",\"email\": \"Shanna@melissa.tv\",\"address\": {\"street\": \"Victor Plains\",\"suite\": \"Suite 879\", \"city\": \"Wisokyburgh\",\"zipcode\": \"90566-7771\"}}"

// \build a UserFacade Mock
type fakeUserFacade struct{}

func (ff fakeUserFacade) GetUsers() (string, error) {
	fmt.Println("GetUsers mocked method called...")
	return "", nil
}

func (ff fakeUserFacade) GetUserById(id string) (string, error) {
	fmt.Println("GetUserById mocked method called...")
	fmt.Println("returning: " + remoteUserJsonOutput)
	return remoteUserJsonOutput, nil
}

// build a UserMapper Mock
type fakeUserMapper struct{}

func (fm fakeUserMapper) GetDomainUser(remoteUser usersapi.User, data string) api.User {
	return userFixture
}

// helpers methods
func getMuxVarsWithId(id string) map[string]string {
	muxVars := make(map[string]string)
	muxVars["Id"] = id
	return muxVars
}

func TestGetUsersById(t *testing.T) {

	cases := []struct {
		fu               fakeUserFacade
		fm               fakeUserMapper
		requestUrl       string
		muxVars          map[string]string
		expectedResponse string
		expectedError    error
		method           string
	}{
		{
			fu:               fakeUserFacade{},
			fm:               fakeUserMapper{},
			requestUrl:       "/users/", //non-valid request
			muxVars:          nil,
			expectedResponse: "",
			expectedError:    common.BadRequest(http.StatusText(http.StatusBadRequest), common.UserIdMissing),
			method:           http.MethodGet,
		},
		{
			fu:               fakeUserFacade{},
			fm:               fakeUserMapper{},
			requestUrl:       "/users/2/", //valid request
			muxVars:          getMuxVarsWithId("2"),
			expectedResponse: userJsonOutput,
			expectedError:    nil,
			method:           http.MethodGet,
		},
	}

	for _, c := range cases {
		// Create a request to pass it to our handler.
		req, err := http.NewRequest(c.method, c.requestUrl, nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		// Our handlerWrapper satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		usersHandler := NewUsersHandler(c.fu, c.fm) //create the OfferHandler with the mocked MobileWSFacade

		err = usersHandler.GetUserById(rr, req, c.muxVars)

		if !reflect.DeepEqual(err, c.expectedError) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedError, err)
		}

		response, err := ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Fatal(err)
		}

		if c.expectedResponse != string(response) {
			t.Errorf("Expected %q but got %q", c.expectedResponse, string(response))
		}
	}

}
