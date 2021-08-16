# article-api

## Prerequisites
* go modules

## Setup/Running the application
1. `go mod download`
2. `go run main.go`
3. `send requests to http://localhost:8080/`

## Endpoints
`POST /articles`

 Sample request body:
 ```
 {
     "title": "article title",
     "body": "article body",
     "date": "2021-01-01",
     "tags": ["tag1, "tag2, "tag3"]
 }
 ```

`GET /articles/{id}`
* `/articles/1`

`GET /tags/{tagName}/{date}`
* `{date}` is in `YYYY-MM-DD` format
* `/tags/tag1/2021-01-01`

## Description
Decided to use `go-sqlite3` -- file-based database for hassle-free and easy to run database solution.

Using gorilla/mux for HTTP router.

Structured the app using the `controller -> service -> repository` pattern. This pattern was used to abstract the data access layer and its implementation and to prevent tight coupling with application/business logic.

Error handling was approached from the perspective of providing the most meaningful error message for the given request.

In terms of testing, I tried to follow TDD as close as possible. A separate test.db was created to run tests on. The aim was to replicate real usage scenarios as close as possible. 

## Assumptions

* Assumed the client will not pass an article id. Instead the id will be generated by the DB and returned when an article is created.
* Two database tables: one for articles and the other for tags

## Future iterations
* Add more validation around user requests
* Add more complex scenario tests
* Simplify SQL queries where possible
* Implement caching

## Tests
`go test -v`
