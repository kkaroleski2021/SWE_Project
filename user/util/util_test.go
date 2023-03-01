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

func TestSomeHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	// req, err := http.NewRequest("GET", "/some-endpoint", nil)
	req, err := http.NewRequest("GET", "/users", nil)
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
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", user.GetUser).Methods("GET")
	return router
}

/*create tests for
-getUsers
-getUser
-CreateUser
-updateUser
-deleteUser
-LogIn*/

func Test_GetUsers(t *testing.T) {
	request, _ := http.NewRequest("GET", "/users/{id}", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}
