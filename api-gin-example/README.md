# API Example with GIN

This is a simple API written using Golang and the [GIN Framework](https://gin-gonic.com).
The API allows a simple CRUD for a User entity. The ORM used to map the structs
to the SQLite3 Database is [GORM](https://gorm.io/).

The application also allows for Users to be authenticated to the API using JWT.
The [...](https://127.0.0.1) package is used to generate the JWT Token.

## API Feature Status
- [ ] Users Endpoints
  - [ ] List all registered Users
  - [ ] Register User
  - [ ] Show specific User
  - [ ] Update Users account
    - [ ] Must be authenticated
  - [ ] Delete Users account
    - [ ] Must be authenticated
- [x] Sessions Endpoint
  - [x] Create session (Authenticate JWT)

## API Tests
- [ ] Unit Tests
  - [ ] TBD
- [ ] Integration Tests
  - [ ] TBD
