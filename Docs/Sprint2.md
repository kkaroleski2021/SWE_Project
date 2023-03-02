# Swampy Sells : Sprint 2

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
- Both front-end work has been successfully combined into one cohesive home page.
- Two Cypress tests have been applied to the frontend code and both passed.


## Front-End  

### Unit Tests
- write unit tests here  

### Cypress Tests
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



## Back-End
### Unit Tests
- write unit tests here

### API Documentation

#### Endpoints
| Path             | HTTP   |
|------------------|--------|
| "/users"         | GET    |
| "/users/{id}"    | GET    |
| "/users/{id}"    | POST   |
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


#### "/users/{id}" - POST 
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
