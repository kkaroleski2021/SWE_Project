/*BACKEND UNIT TESTS*/

package user

import (
	"encoding/json"
	"fmt"
	"go/router"
	"go/user"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	//"net/http/httptest"
	//"path/filepath"
	"testing"
)

// DecodeJSONRequest unmarshals a JSON from src and stores it in dst (must be a pointer).
// If the unmarshalling fails, a 400 with an error message is written on w.
func DecodeJSONRequest(dst interface{}, src io.Reader, w http.ResponseWriter) (succeeded bool) {
	if err := json.NewDecoder(src).Decode(dst); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad request JSON (%v).", err)
		return false
	}

	return true
}

// EncodeJSONResponse marshals an object to JSON and sends a 200 with the
// object, or sends a 500 with an error message is marshalling fails
func EncodeJSONResponse(w http.ResponseWriter, response interface{}) {
	if raw, err := json.Marshal(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, "Failed to encode query result to JSON (%v).", err)
	} else {
		w.Header().Add("Content-type", "application/json; charset=utf-8")
		w.Write(raw)
	}
}

func TestSearch(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	// req, err := http.NewRequest("GET", "/some-endpoint", nil)
	req, err := http.NewRequest("GET", "/search", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(router.Search)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", user.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", user.GetUser).Methods("GET")
	r.HandleFunc("/signup", user.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", user.ValidateMiddleware(user.UpdateUser)).Methods("PUT")
	r.HandleFunc("/users/{id}", user.ValidateMiddleware(user.DeleteUser)).Methods("DELETE")
	r.HandleFunc("/login", user.LogIn).Methods("POST")
	r.HandleFunc("/search", router.Search).Methods("GET")
	r.HandleFunc("/searchhistory", router.SearchHistory).Methods("GET")
	r.HandleFunc("/search", router.SearchPost).Methods("POST")

	return r
}
func Test_SearchHistory(t *testing.T) {
	request, err := http.NewRequest("GET", "/searchhistory", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(router.SearchHistory)

	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func Test_TrialGetUser(t *testing.T) {
	//var User user.User
	request, err := http.NewRequest("GET", "/users", nil)
	//request.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(user.GetUser)

	handler.ServeHTTP(response, request)

	//Router().ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func Test_GetUser(t *testing.T) {
	var User user.User
	request, _ := http.NewRequest("GET", "/users/{id}", nil)
	if request != nil {
		fmt.Println(User.ID)
	}
	response := httptest.NewRecorder()
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Println(User.ID)
}

func Test_GetUsers(t *testing.T) {
	r := Router()
	r.Get("/users", user.GetUsers)
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		fmt.Println("Help!")
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.GetUsers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[
		{
			"ID": 14,
			"CreatedAt": "2023-03-01T22:29:31.286Z",
			"UpdatedAt": "2023-03-01T22:29:31.286Z",
			"DeletedAt": null,
			"firstname": "natasha",
			"lastname": "sthilaire",
			"email": "mileyan25@gmail.com",
			"password": "$2a$10$JJ7gDSb9vN7UiwBRdRZK0OzESazKN.R758b.Je1KkfB2qrQnOYmrm"
		},
		{
			"ID": 17,
			"CreatedAt": "2023-03-02T02:22:33.383Z",
			"UpdatedAt": "2023-03-02T02:27:34.779Z",
			"DeletedAt": null,
			"firstname": "first1",
			"lastname": "last",
			"email": "email1@email",
			"password": "$2a$10$ADd6OUj5h3Y19MfD5HJYkuaYEYCiCOhOB8nDTuWqPHuVtBs.B9soS"
		},
		{
			"ID": 19,
			"CreatedAt": "2023-03-28T14:55:24.478Z",
			"UpdatedAt": "2023-03-28T14:55:24.478Z",
			"DeletedAt": null,
			"firstname": "Adding",
			"lastname": "Product",
			"email": "prod@email",
			"password": "$2a$10$.upfnBlQkLAONTEz.I9ITuRoCbpy4yoJ0OLsMrjTdVJVTYVfZuQDC"
		}
	]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	fmt.Println(rr.Body.String())
	/*request.Header.Set("Content-Type", "application/json")


	defer request.Body.Close()
	response := httptest.NewRecorder()
	//Response Writer, *Request
	handler := http.HandlerFunc(user.GetUsers)
	handler.ServeHTTP(response, request) //throws an exception here
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	fmt.Println("Response: ", response)
	fmt.Println("response code: ", response.Code)*/

}

func Test_CreateUser(t *testing.T) {
	request, _ := http.NewRequest("GET", "/signup", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	//require.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}

func Test_UpdateUser(t *testing.T) {
	request, _ := http.NewRequest("PUT", "/users/{id}", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}

func Test_DeleteUser(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/users/{id}", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}

func Test_LogIn(t *testing.T) {
	data := url.Values{}
	data.Set("logIn", "User")
	request, _ := http.NewRequest("POST", "/login", strings.NewReader(data.Encode()))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}

func Test_Search(t *testing.T) {
	data := url.Values{}
	data.Set("search", "pc")
	request, _ := http.NewRequest("POST", "/search", strings.NewReader(data.Encode()))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}
