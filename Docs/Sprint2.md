# Swampy Sells : Sprint 2

## Video Link: https://youtu.be/NJGJS-FYA2A

## Contents
- [Work Completed](#work-completed)
- [Front-End](#front-end)  
    - [Unit tests](#unit-tests)
    - [Cypress tests](#cypress-tests)
- [Back-End](#front-end) 
    - [Unit tests](#unit-tests)
    - [API Documentation](#api-documentation)


## Work Completed
- A mysql instance was created in Amazon's Web Service , RDS.  This allows the mysql database to be shared, accessed, and edited among users. Having the database set up allows information to be stored about users.
- Password hashing was implemented in order to protect user's passwords as they are stored in the database. 
- When a user attempts to log in with their email and password, the program will check with the database to ensure the correct log in credentials.
- JWT authorization has begun to be implemented. When a user logs in, they will receive a JWT which will allow them to access certain features such as modifying their password. Certain API requests will now be protected and require a valid JWT to access.
- The user can hover over a catagory in the menu to select a sub-catagory of listings.
- Work has begun creating a listing page that outlines the details of the listing, include an image, description from the seller, product rating based on user feedback, price, and item condition. In future sprints, we hope to both create the actual separate page instead of having all the components on the home page, and perhaps implement the mapping function so users can see how close they are to the seller. Additionally, future sprints will see us collecting the data (seller name, product image/price/condition/description/seller location, etc.) so we can stop hardcoding in example blurbs
- Both front-end work has been successfully combined into one cohesive home page.
- Two Cypress tests have been applied to the frontend code and both passed.


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
