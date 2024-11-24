package posts

import (
	"bytes"
	"chirpify/http/controller"
	"chirpify/http/request"
	"chirpify/http/response"
	"chirpify/internal/model"
	"chirpify/internal/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupUpdatePostRouter() *gin.Engine {
	postService := service.NewPostService().(*service.PostServiceImpl)

	// Add mock posts to the in-memory data
	addMockPosts(postService)

	postController := controller.NewPostController(postService)

	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.PUT("/posts/:postID", postController.Update)
	}
	return r
}

func TestUpdatePost(t *testing.T) {
	router := setupUpdatePostRouter()

	// Initializing the pointer values before assignment
	message := "Good Morning!"
	link := "http://example.com"
	caption := "Example Caption"
	description := "This is a sample description"
	picture := "http://example.com/picture.jpg"
	name := "Sample Post"
	tags := []string{"tag1", "tag2"}
	source := "mobile"
	fileURL := "http://example.com/file.mp4"
	privacy := model.PrivacySettings{Value: "public"} // Assuming model.PrivacySettings is a struct
	scheduledPublishTime := int64(1672531200)
	allowComments := true
	location := model.GeoLocation{Latitude: 37.7749, Longitude: -122.4194} // Assuming model.GeoLocation is a struct
	status := "published"

	updatePostRequest := request.UpdatePostRequest{
		Message:              &message,
		Link:                 &link,
		Caption:              &caption,
		Description:          &description,
		Picture:              &picture,
		Name:                 &name,
		Tags:                 &tags,
		Source:               &source,
		FileURL:              &fileURL,
		Privacy:              &privacy,
		ScheduledPublishTime: &scheduledPublishTime,
		AllowComments:        &allowComments,
		Location:             &location,
		Status:               &status,
	}

	// Creating the request body in JSON format
	body, err := json.Marshal(updatePostRequest)
	if err != nil {
		t.Fatalf("Could not marshal request: %v", err)
	}

	req, _ := http.NewRequest("PUT", "/api/v1/posts/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var respBody response.Response
	err = json.Unmarshal(resp.Body.Bytes(), &respBody)
	assert.NoError(t, err)
	assert.Equal(t, "Ok", respBody.Status)
}
