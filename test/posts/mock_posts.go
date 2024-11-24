package posts

import (
	"chirpify/model"
	"chirpify/service"
)

// Helper function to add mock posts
func addMockPosts(postService *service.PostServiceImpl) {
	// Initializing the pointer values before assignment
	message := "Good Morning Friends!"
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

	// Directly accessing the Posts map in PostServiceImpl
	postService.Posts[1] = model.Post{
		ID:                   1,
		UserID:               126,
		Message:              message,
		Link:                 link,
		Caption:              caption,
		Description:          description,
		Picture:              picture,
		Name:                 name,
		Tags:                 tags,
		Source:               source,
		FileURL:              fileURL,
		Privacy:              privacy,
		ScheduledPublishTime: scheduledPublishTime,
		AllowComments:        allowComments,
		Location:             &location,
		Status:               status,
	}
}
