package router

import (
	"chirpify/controller"
	"chirpify/service"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	PostController *controller.PostController
}

func InitializeDependencies() *Dependencies {
	postService := service.NewPostService()
	postController := controller.NewPostController(postService)

	return &Dependencies{
		PostController: postController,
	}
}

func NewRouter(deps *Dependencies) *gin.Engine {
	router := gin.Default()

	baseRouter := router.Group("/api/v1")

	// Route for creating a new post
	baseRouter.POST("/posts", deps.PostController.Create)

	// Route for updating an existing post
	baseRouter.PUT("/posts/:postID", deps.PostController.Update)

	// Route for retrieving all posts
	baseRouter.GET("/posts", deps.PostController.GetAllPosts)

	// Route for liking a post
	baseRouter.POST("/posts/:postID/like", deps.PostController.LikePost)

	// Route for adding a comment to a post
	baseRouter.POST("/posts/:postID/comments", deps.PostController.AddComment)

	// Route for retrieving details of a post, including comments and likes
	baseRouter.GET("/posts/:postID/details", deps.PostController.GetPostDetails)

	// Return the Gin router with all routes set up
	return router
}
