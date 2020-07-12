# API Example with GIN

This is a simple API written using Golang and the [GIN Framework](https://gin-gonic.com).
The API allows a simple CRUD for a User entity. The ORM used to map the structs
to the SQLite3 Database is [GORM](https://gorm.io/).

The application also allows for Users to be authenticated to the API using JWT.
The [jwt-go](https://github.com/dgrijalva/jwt-go) package is used to generate the JWT Token.

## API Feature Status
- [x] Users Endpoints
  - [x] List all registered Users
  - [x] Register User
    - [x] Persists User on DB
    - [x] Must hash password
  - [x] Show specific User
  - [x] Update Users account
    - [x] Must be authenticated
  - [x] Delete Users account
    - [x] Must be authenticated
- [x] Sessions Endpoint
  - [x] Create session (Authenticate JWT)

## API Tests
- [ ] Unit Tests
  - [ ] TBD
- [ ] Integration Tests
  - [ ] TBD
