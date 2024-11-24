package posts

import (
	"bytes"
	"chirpify/controller"
	"chirpify/data/request"
	"chirpify/data/response"
	"chirpify/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupAddCommentRouter() *gin.Engine {
	postService := service.NewPostService().(*service.PostServiceImpl)

	// Add mock posts to the in-memory data
	addMockPosts(postService)

	postController := controller.NewPostController(postService)

	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.POST("/posts/:postID/comments", postController.AddComment)
	}
	return r
}

func TestAddComment(t *testing.T) {
	router := setupAddCommentRouter()

	commentRequest := request.CommentRequest{
		UserID:          126,
		CommentText:     "This is a comment on the post uytrdfghjkljhghjkhghj.",
		ParentCommentID: 1,
		IsEdited:        false,
		Attachments:     []string{}, // Assuming attachments is an array of strings
	}

	// Creating the request body as per the updated CommentRequest struct
	body, err := json.Marshal(commentRequest)
	if err != nil {
		t.Fatalf("Could not marshal request: %v", err)
	}

	// Sending the request to the server
	req, _ := http.NewRequest("POST", "/api/v1/posts/1/comments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var respBody response.Response
	err = json.Unmarshal(resp.Body.Bytes(), &respBody)
	assert.NoError(t, err)
	assert.Equal(t, "Success", respBody.Status)

}
