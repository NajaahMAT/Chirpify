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

func setupLikePostRouter() *gin.Engine {
	postService := service.NewPostService().(*service.PostServiceImpl)

	// Add mock posts to the in-memory data
	addMockPosts(postService)

	postController := controller.NewPostController(postService)

	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.POST("/posts/:postID/like", postController.LikePost)
	}
	return r
}

func TestLikePost(t *testing.T) {
	router := setupLikePostRouter()

	likeRequest := request.LikeRequest{
		PostID: 1,
		UserID: 126,
	}

	// Marshalling the request payload into JSON
	requestBody, err := json.Marshal(likeRequest)
	if err != nil {
		t.Fatalf("Could not marshal request: %v", err)
	}

	// Creating the request for the "POST" method
	req, _ := http.NewRequest("POST", "/api/v1/posts/1/like", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var respBody response.Response
	err = json.Unmarshal(resp.Body.Bytes(), &respBody)
	assert.NoError(t, err)
	assert.Equal(t, "Ok", respBody.Status)
}
