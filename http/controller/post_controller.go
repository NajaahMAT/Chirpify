package controller

import (
	"chirpify/http/request"
	"chirpify/http/response"
	"chirpify/internal/service"
	"log"
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

// @Summary Create a new post
// @Description Create a new post for a user
// @Tags Posts
// @Accept json
// @Produce json
// @Param post body request.CreatePostRequest true "Create Post Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /posts [post]
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

// @Summary Update an existing post
// @Description Update an existing post with new data
// @Tags Posts
// @Accept json
// @Produce json
// @Param postID path int64 true "Post ID"
// @Param post body request.UpdatePostRequest true "Update Post Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /posts/{postID} [put]
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

// @Summary Retrieve all posts
// @Description Get a list of all posts
// @Tags Posts
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /posts [get]
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

// @Summary Like a specific post
// @Description Like a post by providing the like request
// @Tags Posts
// @Accept json
// @Produce json
// @Param likeRequest body request.LikeRequest true "Like Post Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /posts/like [post]
func (controller *PostController) LikePost(ctx *gin.Context) {
	var likeRequest request.LikeRequest

	// Bind the incoming JSON request body to the LikeRequest struct
	if err := ctx.ShouldBindJSON(&likeRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	msg, err := controller.postService.LikePost(likeRequest)
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

// @Summary Add a comment to a specific post
// @Description Add a comment to an existing post
// @Tags Posts
// @Accept json
// @Produce json
// @Param postID path int64 true "Post ID"
// @Param commentRequest body request.CommentRequest true "Comment Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /posts/{postID}/comments [post]
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

	// Log the commentID to ensure it's being returned correctly
	log.Printf("Comment ID: %d", commentID)

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

// @Summary Retrieve all details of a specific post, including comments and likes
// @Description Get all details of a specific post by its ID
// @Tags Posts
// @Produce json
// @Param postID path int64 true "Post ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /posts/{postID}/details [get]
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
