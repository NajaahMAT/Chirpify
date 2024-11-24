package controller

import (
	"chirpify/data/request"
	"chirpify/data/response"
	"chirpify/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService service.PostService
}

func NewPostController(service service.PostService) *PostController {
	return &PostController{
		postService: service,
	}
}

// Create - Create a new post
func (controller *PostController) Create(ctx *gin.Context) {
	var createPostRequest request.CreatePostRequest
	err := ctx.ShouldBindJSON(&createPostRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	postID, err := controller.postService.CreatePost(createPostRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data: map[string]int64{
			"post_id": postID,
		},
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Update - Update an existing post
func (controller *PostController) Update(ctx *gin.Context) {
	postID := ctx.Param("postID")
	id, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Invalid post ID format",
		})
		return
	}

	var updatePostRequest request.UpdatePostRequest
	err = ctx.ShouldBindJSON(&updatePostRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	err = controller.postService.UpdatePost(id, updatePostRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   "Post updated successfully",
	})
}

// GetAllPosts - Retrieve all posts
func (controller *PostController) GetAllPosts(ctx *gin.Context) {
	posts, err := controller.postService.GetAllPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   posts,
	})
}

// LikePost - Like a specific post
func (controller *PostController) LikePost(ctx *gin.Context) {
	postID := ctx.Param("postID")
	id, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Invalid post ID format",
		})
		return
	}

	userID := ctx.GetInt64("userID") // Assuming userID is extracted from JWT or context

	msg, err := controller.postService.LikePost(userID, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   msg,
	})
}

// AddComment - Add a comment to a specific post
func (controller *PostController) AddComment(ctx *gin.Context) {
	postID := ctx.Param("postID")
	id, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Invalid post ID format",
		})
		return
	}

	var commentRequest request.CommentRequest
	err = ctx.ShouldBindJSON(&commentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	commentID, err := controller.postService.AddComment(id, commentRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	// Return success response with the comment ID and success message
	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data: map[string]interface{}{
			"message":    "Comment added successfully",
			"comment_id": commentID,
		},
	})
}

// GetPostDetails - Retrieve all details of a specific post, including comments and likes
func (controller *PostController) GetPostDetails(ctx *gin.Context) {
	postID := ctx.Param("postID")
	id, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Invalid post ID format",
		})
		return
	}

	postDetails, err := controller.postService.GetPostDetails(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   postDetails,
	})
}
