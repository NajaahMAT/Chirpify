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
    - To view api documentation:
    use  following url in your browser after executing the project.
    ```http://localhost:8081/swagger/index.html```

## **assumption made while implementing the system**
#### Post ID Generation:
The PostServiceImpl uses an internal counter NextPostID to generate unique post IDs. This assumes that the post ID doesn't need to be persisted across application restarts (i.e., it's generated in memory). If the app needs to scale or recover from a restart, a persistent database with auto-incrementing IDs would be more appropriate.
#### In-Memory Data Storage:
All data (posts, comments, likes) is stored in memory (map[int64]model.Post, map[int64][]model.Comment, etc.). This assumption means that the data is not persistent and will be lost if the server restarts. A real application would require a database (e.g., PostgreSQL, MySQL, or MongoDB) to store this data permanently.
#### Concurrency Control:
The PostServiceImpl uses a mutex (sync.Mutex) to manage concurrent access to in-memory data structures. This ensures that only one goroutine can modify the data at a time. While this is fine for small-scale usage, for a production system, distributed locks or database transaction handling might be needed to ensure data integrity across multiple instances or servers.
#### Basic Error Handling:
Error handling is done by returning error messages in JSON responses (e.g., ctx.JSON(http.StatusBadRequest, response.Response{...})). It assumes that these error messages are sufficient for debugging or informing the user. In a production environment, detailed error logs, monitoring, and more user-friendly messages might be necessary.
#### Lack of Authentication/Authorization:
The code doesn't include any authentication or authorization mechanisms. It assumes that the API is open or that a separate mechanism (e.g., OAuth2, JWT) will be used outside the current code structure. This is a critical area for securing the API, especially in a social media app.
#### Data Validation:
Input validation is done by the ShouldBindJSON method, but the assumptions are that the input format is always correct according to the request structures (e.g., CreatePostRequest, UpdatePostRequest). Additional validation, such as checking for the length of messages or ensuring links are valid URLs, might be needed in a real-world scenario.
#### Post Privacy Handling:
The Privacy field in the CreatePostRequest is assumed to be a simple enum or string-based value, with values such as "Public" or "Private". The actual implementation of privacy is not provided, and there's no enforcement on the backend, meaning privacy enforcement would need to be handled separately.
#### Scheduled Posts:
The app allows for scheduled posts, where a post can be scheduled for future publishing via ScheduledPublishTime. However, the logic for handling the timing (e.g., checking whether it's time to publish the post) is not implemented, and it is assumed that another part of the system will handle this.
#### No Caching or Indexing:
There is no caching layer or indexing of posts for faster retrieval, which might be necessary when the number of posts grows. For large-scale applications, using a database with indexing or caching solutions like Redis might be required to improve performance.
#### Assumption of Post-Related Entities:
The app assumes that each post can have likes and comments, and those are stored and associated with posts. There are no constraints or checks on the number of likes or comments a post can have, which could lead to performance issues in the future if the number grows significantly.
#### Response Structure:
The API responses follow a consistent format (response.Response{}), which includes a Code, Status, and Data. This is a simple, uniform structure for all responses, but in a more complex system, different types of responses may require more granular structures.
#### No Rate Limiting:
There is no rate limiting or throttling implemented for the API. For a production system, rate limiting would be needed to prevent abuse and ensure fair usage of the service.
#### No Pagination:
The GetAllPosts function retrieves all posts without any pagination. This could cause performance issues as the number of posts increases. In a production system, pagination or filtering would be required to manage large datasets effectively.
#### No Support for Media Storage:
The Picture field in the post creation request suggests that users can upload media, but there is no system for handling file uploads or media storage. In a real-world scenario, you'd need integration with a file storage service (e.g., AWS S3, Google Cloud Storage) or a local file system for handling media content.
#### Post Status Updates:
The UpdatePost function allows updating post details such as Message, Link, etc., but assumes that posts will always exist and that their status can be modified freely. There is no logic in place for limiting status changes (e.g., after a post is published).
#### Model Structure Assumptions:
The model structures like Post, Comment, and LikeRecord are assumed to be predefined and structured appropriately for the needs of the app. However, the exact details of these models are not given here, and it's assumed that they will be properly designed to store the necessary fields.

    
## **Suggestions for Future Enhancements**
- **Database Enhancements**:
  #### Normalize the data schema to handle relationships like `users`, `posts`, `comments`, and `likes`.
  #### Implement caching for frequently accessed data using Redis.

- **Features**:
  #### Add user profiles and the ability to follow/unfollow users.
  #### Add a search feature to find posts by tags or keywords.

- **Performance**:
  #### Integrate a task queue for background processing (e.g., Go worker queues).

- **Security**:
  #### Add role-based access control (RBAC).
  #### Secure sensitive endpoints using rate limiting and API key validation.

