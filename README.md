# Chirpify
Chirpify is a lightweight backend API for a mini social media platform built using GoLang, designed to demonstrate RESTful API principles with clean architecture.

## Features Covered
- **Post Management**: Create, update, and retrieve posts.
- **Likes and Comments**: Add likes and comments to posts.
- **Post Details**: Retrieve details of a post, including comments and likes.

## Tech Stack
- **Backend**: Golang
- **Framework**: Gin
- **Database**: In-memory map (can be replaced with PostgreSQL or MongoDB)
- **API Docs**: Swagger

## Implementation Details
- **Main Code**: Located in the `main.go` file.
- **Router**: API endpoints are defined in `/http/router.go`.
- **Controllers**: Business logic for posts is handled in `/http/controller`.
- **Services**: Data manipulation and in-memory storage logic are implemented in `/internal/service`.
- **API Documentation**: Swagger Api Documentations implemented in `/docs/swagger.json`
- **Postman Collection**: Postman collections included in `Chipify.postman_collection.json`


## How to Run
1. Clone the repository.
2. Install dependencies:
   ```bash
   go mod tidy
3. Start the application:
   ```go run main.go```
4. To run the test:
    ```go test -v ./test/posts/```
5. To Install and Build Swagger API Document:
    - Install swagger (Windows OS):
    ```go install github.com/swaggo/swag/cmd/swag@latest```
    - Generate swagger:
    ```swag init -g cmd/main.go --output ./docs --parseDependency --parseInternal```
