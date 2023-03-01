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


## Front-End
is there a difference between cypress and unit ? --- delete cypress if no difference  

### Unit Tests
- write unit tests here  

### Cypress Tests
- write cypress tests here  



## Back-End
### Unit Tests
- write unit tests here

### API Documentation
- [Search](#search)
- [Users](#user)

