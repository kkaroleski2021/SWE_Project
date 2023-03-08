# Swampy Sells : Sprint 2

## Video Link: 

## Contents
- [Work Completed](#work-completed)
- [Front-End](#front-end)  
    - [Unit tests](#unit-tests)
    - [New Unit Tests](#new-unit-tests)
    - [Cypress tests](#cypress-tests)
    - [New Cypress Tests](#new-cypress-tests)
- [Back-End](#back-end) 
    - [Unit tests](#unit-tests-1)
    - [New Unit Tests](#new-unit-tests-1)
    - [API Documentation](#api-documentation)


## Work Completed
- 

## Front-End  

### Unit Tests
- NOTE: Most of these are the default angular ones, but the one at the bottom used to see if the images rendered was homemade - not sure if it actually works though
```
import { TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AppComponent } from './app.component';

describe('AppComponent', () => {
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        RouterTestingModule
      ],
      declarations: [
        AppComponent
      ],
    }).compileComponents();
  });

  it('should create the app', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });

  it(`should have as title 'hello-world'`, () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;
    expect(app.title).toEqual('hello-world');
  });

  it('should render title', () => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('.content span')?.textContent).toContain('hello-world app is running!');
  });

  it('should set image path as expected', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const ele = fixture.debugElement.nativeElement.querySelectorAll('img');
    expect(ele[1]['src']).toContain('assets/images/stuffedAnimalChair.jpg');
    expect(ele[2]['src']).toContain('assets/images/clamshellVase.jpg');
    expect(ele[3]['src']).toContain('assets/images/grandmasCouch.jpg');
    expect(ele[4]['src']).toContain('assets/images/clamshellVase.jpg');
});
});
```

### New Unit Tests



### Cypress Tests
```
describe('My First Test', () => {
  it('visits the home page', () => {
    cy.visit('http://localhost:4200')
  })
})

describe('My Second Test', () => {
  it('finds the content "Tickets"', () => {
    cy.visit('http://localhost:4200')

    cy.contains('Tickets')
  })
})
```


### New Cypress Tests






## Back-End
### Unit Tests
```
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
	request, _ := http.NewRequest("GET", "/users", nil)

	if request != nil {
		fmt.Println("we are in here")
		fmt.Println(User.ID)
	}

	response := httptest.NewRecorder()
	//handler := http.HandlerFunc(user.GetUser)

	fmt.Println("Response: ", response)

	fmt.Println("response code: ", response.Code)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func Test_GetUsers(t *testing.T) {

	request, err := http.NewRequest("GET", "users/{id}", nil)
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	//Response Writer, *Request
	handler := http.HandlerFunc(user.GetUser)

	handler.ServeHTTP(response, request) //throws an exception here

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

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
	data.Set("login", "User")
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

```


### New Unit Tests






### API Documentation

#### Endpoints
| Path             | HTTP   |
|------------------|--------|
| "/users"         | GET    |
| "/users/{id}"    | GET    |
| "/users/{id}"    | PUT    |
| "/users/{id}"    | DELETE |
| "/signup"        | POST   |
| "/login"         | POST   |
| "/search"        | GET    |
| "/search"        | POST   |
| "/searchhistory" | GET    |


#### "/users" - GET 
This method returns a list of all of the users in the database.   
Request Format: none  
Response Format: 
```
[
    {
        "ID": "10",
        "CreatedAt": "yr-mo-day...",
        "UpdatedAt": "yr-mo-day...",
        "DeletedAt": null,
        "firstname": "first",
        "lastname": "last",
        "email": "email3@email",
        "password": "hashed password"
    },
]
```


#### "/users/{id}" - GET 
This method returns a user by their ID.    
Request Format: none
Response Format: 
```
[
    {
        "ID": "10",
        "CreatedAt": "yr-mo-day...",
        "UpdatedAt": "yr-mo-day...",
        "DeletedAt": null,
        "firstname": "first",
        "lastname": "last",
        "email": "email3@email",
        "password": "hashed password"
    },
]
```


#### "/users/{id}" - PUT 
This method updates the user by id.  
Must have a JWT authorization token.  
Request Format: anything that is intended to be changed.
```
{
    "firstname":"firstChanged2",
}
``` 
Response Format: 

```
{
    "ID": 10,
    "CreatedAt": "2023-02-17T03:42:48.752Z",
    "UpdatedAt": "2023-03-01T21:06:30.288-05:00",
    "DeletedAt": null,
    "firstname": "firstChanged2",
    "lastname": "last3",
    "email": "email3@email",
    "password": "hashed password"
}
"The user has been updated"
```


#### "/users/{id}" - DELETE 
This method deletes the user by id.   
Must have a JWT authorization token.  
Request Format: None   
Response Format: 

```
"The user has been successfully deleted!"
```




#### "/signup" - POST
This method adds a user to the database.  
Email and Password are required fields.    
Request Format: 
```
{
    "firstname":"firstChanged",
    "lastname":"last3",
    "email":"email3@email",
    "password":"something3"
}
``` 
Response Format: 
```
{
        "ID": "10",
        "CreatedAt": "yr-mo-day...",
        "UpdatedAt": "yr-mo-day...",
        "DeletedAt": null,
        "firstname": "first",
        "lastname": "last",
        "email": "email3@email",
        "password": "hashed password"
}
```




#### "/login" - POST 
This method attempts to log a user in by comparing the inputed password and email to the password associated with the email in the database.    

Request Format: 
```
{
    "email":"email3@email",
    "password":"something3"
}
```
Response Format:
```
"The user has been successfully logged in"
{
    "data": {
        "access-token": "eyJhbG....."
    },
    "message": "success",
    "status": true
}
```


#### "/search" - GET 
This method retrieves the last search input.   
Request Format: None   
Response Format: 

```
[{"Search":"pc"}]
```



#### "/search" - POST 
This method saves a search input.   
Request Format: 
```
{
    "search":"pc"
}
```
Response Format: none


#### "/searchhistory" - POST 
This method gets all previous searches.    
Request Format: None   
Response Format: 

```
[{"Search":"pc"},{"Search":"books"}]
```
