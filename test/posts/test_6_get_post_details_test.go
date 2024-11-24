package posts

import (
	"chirpify/controller"
	"chirpify/data/response"
	"chirpify/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupGetPostDetailsRouter() *gin.Engine {
	postService := service.NewPostService().(*service.PostServiceImpl)

	// Add mock posts to the in-memory data
	addMockPosts(postService)

	postController := controller.NewPostController(postService)

	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.GET("/posts/:postID/details", postController.GetPostDetails)
	}
	return r
}

func TestGetPostDetails(t *testing.T) {
	router := setupGetPostDetailsRouter()

	req, _ := http.NewRequest("GET", "/api/v1/posts/1/details", nil)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var respBody response.Response
	err := json.Unmarshal(resp.Body.Bytes(), &respBody)
	assert.NoError(t, err)
	assert.Equal(t, "Ok", respBody.Status)
}
