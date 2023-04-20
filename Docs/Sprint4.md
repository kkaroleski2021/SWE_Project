# Swampy Sells : Sprint 4

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
- When uploading an image, the backend code now accounts for files with the same names. File names are now given a randomized prefix and are checked against existing files. The extended name is now the official Fname and part of the path in the upFile table in the database rather than the original file name. 
- Created a client profile in the database so that clients can access the database under certain permissions.
- A testing program was created in order to work on routing and information retrieval from the GO API on a smaller scale.
- The json fields for adding a new listing have been extended in the backend to allow for more information to be stored. And the backend API has been adjusted accordingly.
- CORS headers were added to ensure frontend is able to receive information from the API.
- Routing on the frontend was changed from new local host tabs to components
- Buttons were fixed and succesfully load to the desired pages

## Front-End  


### New Unit Tests
beforeEach(() => {
  TestBed.configureTestingModule({
    imports: [
        AppRoutingModule.withRoutes(
        [{path: '', component: HomeComponent}, {path: 'Login', component: BComponent}, {path: 'signup', component: SComponent}]
      )
    ]
  });

})

it('should render title', () => { //6
    const fixture = TestBed.createComponent(Header2Component);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h1').textContent).contain('Log In to Swampy Sells!');
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
func Test_OrderedProduct(t *testing.T) {
	var jsonStr = []byte(`{"ID":4,"ProductID":"xyz","ProductQuantity":20,"OrderID":"123"}`)
	req, err := http.NewRequest("POST", "/order", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(product.AddOrder)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"ID":4,"ProductID":"xyz","ProductQuantity":20,"OrderID":"123"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_AddProduct(t *testing.T) {
	var jsonStr = []byte(`{"UserID":4,"UserPhoneNum":"5617277352","Street":"ABC","City":"WPB","State":"FL",
	"Zip":"33409", "Name":"Headphones", "Description":"Sony", "Condition":"Gently Used", "Tags":"Trendy, Modern",
	"Price": 250}`)
	req, err := http.NewRequest("POST", "/newlisting", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(product.AddProduct)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"UserID":4,"UserPhoneNum":"5617277352","Street":"ABC","City":"WPB","State":"FL",
	"Zip":"33409", "Name":"Headphones", "Description":"Sony", "Condition":"Gently Used", "Tags":"Trendy, Modern",
	"Price": 250}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_UploadImage(t *testing.T) {
	var jsonStr = []byte(`{"ProdID":4,"FName":"A","Fsize":"150","Ftype":"PNG","Path":"C:/Users/Downloads"}`)
	req, err := http.NewRequest("POST", "/newlisting/addimages", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(product.UploadImg)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"ProdID":4,"FName":"A","Fsize":"150","Ftype":"PNG","Path":"C:/Users/Downloads"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_DeleteUserProfile(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/users/updateprofile/delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("ID", "19")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.DeleteUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"ID":19,"firstname":"changed","lastname":"Product","email":"prod@email",
	"password":"$2a$10$.upfnBlQkLAONTEz.I9ITuROCbpy4yoJ0OLFsMrjTdVJVTYTfZuQDC"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}



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


### API Documentation

#### Endpoints
| Path             					| HTTP   |
|-----------------------------------|--------|
| "/users"         					| GET    |
| "/users/{id}"    					| GET    |
| "/users/updateprofile"    		| PUT    |
| "/users/updateprofile/delete"   	| DELETE |
| "/signup"        					| POST   |
| "/login"        					| POST   |
| "/newlisting"         			| POST   |
| "/newlisting/addimages"         	| POST   |
| "/search"        					| GET    |
| "/search"        					| POST   |
| "/order"                  | POST   |
| "/searchhistory" 					| GET    |


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
```

#### "/newlisting" - POST 
This method submits a new listing to the database. 
It is stored in a separate table with the user id of the user that submitted the listing.
Name and price are required fields.   

Request Format: 
```
{
    "phoneNumber":"1234567890",
    "street":"",
    "city":"gainesville",
    "state":"Florida",
    "zip":"32611",
    "name":"sink",
    "descrip":"a great sink",
    "condition":"good",
    "tags":"furniture",
    "price":30
}
```
Response Format:
```
{
    "ID": 26,
    "CreatedAt": "2023-04-19T14:30:17.47-04:00",
    "UpdatedAt": "2023-04-19T14:30:17.47-04:00",
    "DeletedAt": null,
    "UserId": 19,
    "phoneNumber": "1234567890",
    "street": "",
    "city": "gainesville",
    "state": "Florida",
    "zip": "32611",
    "name": "sink",
    "descrip": "a great sink",
    "condition": "good",
    "tags": "furniture",
    "price": 30
}
```

#### "/newlisting/addimages" - POST 
This method submits adds images to the database corresponding with a product. 
It is stored in a separate table with the product id that the images belong to.
The content type is of form-data and the key is "file".

Request Format: 
```
{
  	key = file 
	value = uploadedimg.jpg
}
```
Response Format:
```
{
    "ID": 2,
    "CreatedAt": "2023-03-29T21:35:39.413-04:00",
    "UpdatedAt": "2023-03-29T21:35:39.413-04:00",
    "DeletedAt": null,
    "ProdID": 19,
    "Fname": "Untitled.jpg",
    "Fsize": "",
    "Ftype": "",
    "Path": "./frontend/src/assets/uploadsUntitled.jpg"
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
### "/order" - POST
This method adds an order to the database for ordered products.

