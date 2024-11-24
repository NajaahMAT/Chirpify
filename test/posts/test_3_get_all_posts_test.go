package posts

import (
	"chirpify/http/controller"
	"chirpify/http/response"
	"chirpify/internal/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupGetAllPostsRouter() *gin.Engine {
	postService := service.NewPostService().(*service.PostServiceImpl)

	// Add mock posts to the in-memory data
	addMockPosts(postService)

	postController := controller.NewPostController(postService)

	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.GET("/posts", postController.GetAllPosts)
	}
	return r
}

func TestGetAllPosts(t *testing.T) {
	router := setupGetAllPostsRouter()

	req, _ := http.NewRequest("GET", "/api/v1/posts", nil)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var respBody response.Response
	err := json.Unmarshal(resp.Body.Bytes(), &respBody)
	assert.NoError(t, err)
	assert.Equal(t, "Ok", respBody.Status)
}
