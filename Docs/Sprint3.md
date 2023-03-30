# Swampy Sells : Sprint 2

## Video Link: 

## Contents
- [Work Completed](#work-completed)
- [Front-End](#front-end)  
	- [New Unit Tests](#new-unit-tests)
    - [Unit tests](#unit-tests) 
    - [New Cypress Tests](#new-cypress-tests)
	- [Cypress tests](#cypress-tests)
- [Back-End](#back-end) 
    - [New Unit Tests](#new-unit-tests-1)
	- [Unit tests](#unit-tests-1)
    - [API Documentation](#api-documentation)


## Work Completed
- The ability to upload products to a product table in the database was added.
- The product currently has several fields and it linked to the respective user by id (more on that below).  
- Due to the complexity of JWT authentication, we have opted for a more simplistic cookie session method of authenticating users. This method supports url paths without the need for user id to be included. 
	- When a user logs in or signs up, they will be given a cookie. This cookie tells the system that user id #x is logged in. 
	- This cookie grants access to updating and deleting the user associated with that id. 
	- It also allows user to upload products which will be saved with their respective id.
-  The ability to upload multiple images for a product was implemented. 
	- These images will be saved to a file in 
	```/frontend/src/assets/uploads```. Several pieces of information will be stored including the filepath and the associated product ID in a table in the database.  
- Further work on areas surrounding uploading and displaying product images will be handled in the future.
- a login page and a sign up page was created 
	- the login page takes a user's password and username as input 
	- the sign up page takes in the user's first name, last name, email, and password
- the login/signup page has been linked to the home page through different buttons

## Front-End  


### New Unit Tests
it('should render a button as a Link, checks for href attribute and primary class', () => {
  render(<ButtonAsLink />)
  const buttonAsLink = screen.getByRole('link', { name: /link/i })
  expect(buttonAsLink).toHaveClass('primary')
  expect(buttonAsLink).toHaveAttribute('href', '/')
});

it(`should have as title 'signup'`, () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;
    expect(app.title).toEqual('signup');
  });
  
  it(`should have as title 'login2'`, () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;
    expect(app.title).toEqual('login2');
  });

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





## Back-End



### New Unit Tests
Tavern tests:

test_name: "Integration test with Tavern"

stages:

- name: "Test GetUser"
  request:
    url: "http://localhost:9000/users/14"
    method: GET
  response:
    status_code: 200

test_name: "Test GetUsers"

stages:

- name: "Test GetUsers"
  request:
    url: "http://localhost:9000/users"
    method: GET
  response:
    status_code: 200


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


### API Documentation

#### Endpoints
| Path             					| HTTP   |
|-----------------------------------|--------|
| "/users"         					| GET    |
| "/users/{id}"    					| GET    |
| "/users/updateprofile"    		| PUT    |
| "/users/updateprofile/delete"   	| DELETE |
| "/signup"        		| POST   |
| "/login"        		| POST   |
| "/newlisting"         | POST   |
| "/search"        		| GET    |
| "/search"        		| POST   |
| "/searchhistory" 		| GET    |


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


#### "/users/updateprofile" - PUT 
This method updates the user by id which is stored in a cookie.  
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


#### "/users/updateprofile/delete" - DELETE 
This method deletes the user by id.   
The id is stored and retrieved from a cookie. 
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

#### "/newlisting" - POST 
This method submits a new listing to the database. 
It is stored in a separate table with the user id of the user that submitted the listing.
Name and price are required fields.   

Request Format: 
```
{
    "name":"desk",
    "tags":"furniture",
    "price":"15"
}
```
Response Format:
```
{
    "name":"desk",
    "tags":"furniture",
    "price":"15"
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

### "/chat" 
This function will serve as a means for users to communicate with other users about potential items that they will look to sell or look into. This feature involves implementing a basic WebSocket server which will listen for messages and write them back to via the same WebSocket. 