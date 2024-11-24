package service

import (
	"chirpify/data/request"
	"chirpify/data/response"
	"chirpify/helper"
	"chirpify/model"
	"errors"
	"fmt"
	"sync"
	"time"
)

type PostServiceImpl struct {
	Posts      map[int64]model.Post
	Comments   map[int64][]model.Comment
	Likes      map[int64][]model.LikeRecord // postID -> userID -> like status
	NextPostID int64
	Mutex      sync.Mutex
}

func NewPostService() PostService {
	return &PostServiceImpl{
		Posts:      make(map[int64]model.Post),
		Comments:   make(map[int64][]model.Comment),
		Likes:      make(map[int64][]model.LikeRecord),
		NextPostID: 1,
	}
}

func (s *PostServiceImpl) CreatePost(req request.CreatePostRequest) (int64, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	postID := s.NextPostID
	s.Posts[postID] = model.Post{
		ID:                   postID,
		UserID:               int64(req.UserID),
		Message:              helper.DerefString(req.Message),
		Link:                 helper.DerefString(req.Link),
		Caption:              helper.DerefString(req.Caption),
		Description:          helper.DerefString(req.Description),
		Picture:              helper.DerefString(req.Picture),
		Name:                 helper.DerefString(req.Name),
		Tags:                 helper.DerefStringArray(req.Tags),
		Source:               helper.DerefString(req.Source),
		FileURL:              helper.DerefString(req.FileURL),
		Privacy:              derefPrivacy(req.Privacy),
		ScheduledPublishTime: helper.DerefInt64(req.ScheduledPublishTime),
		AllowComments:        helper.DerefBool(req.AllowComments),
		Location:             req.Location,
		Status:               helper.DerefString(req.Status),
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	s.NextPostID++
	return postID, nil
}

func (s *PostServiceImpl) UpdatePost(postID int64, req request.UpdatePostRequest) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	post, exists := s.Posts[postID]
	if !exists {
		return errors.New("post not found")
	}

	// Update fields conditionally
	if req.Message != nil {
		post.Message = *req.Message
	}
	if req.Link != nil {
		post.Link = *req.Link
	}
	if req.Caption != nil {
		post.Caption = *req.Caption
	}
	if req.Description != nil {
		post.Description = *req.Description
	}
	if req.Picture != nil {
		post.Picture = *req.Picture
	}
	if req.Name != nil {
		post.Name = *req.Name
	}
	if req.Tags != nil {
		post.Tags = *req.Tags
	}
	if req.Source != nil {
		post.Source = *req.Source
	}
	if req.FileURL != nil {
		post.FileURL = *req.FileURL
	}
	if req.Privacy != nil {
		post.Privacy = *req.Privacy
	}
	if req.ScheduledPublishTime != nil {
		post.ScheduledPublishTime = *req.ScheduledPublishTime
	}
	if req.AllowComments != nil {
		post.AllowComments = *req.AllowComments
	}
	if req.Location != nil {
		post.Location = req.Location
	}
	if req.Status != nil {
		post.Status = *req.Status
	}

	post.UpdatedAt = time.Now()
	s.Posts[postID] = post
	return nil
}

func (s *PostServiceImpl) GetAllPosts() ([]model.Post, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	var posts []model.Post
	for _, post := range s.Posts {
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *PostServiceImpl) LikePost(request request.LikeRequest) (string, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	// Initialize Likes map for the post if it doesn't exist
	if _, exists := s.Likes[request.PostID]; !exists {
		s.Likes[request.PostID] = []model.LikeRecord{}
	}

	// Check if the user has already liked the post
	for _, like := range s.Likes[request.PostID] {
		if like.UserID == request.UserID {
			return "Post already liked by the same user", nil // Return success message when the user has already liked the post
		}
	}

	// Add like if the user hasn't liked the post
	likeRecord := model.LikeRecord{
		UserID: request.UserID,
		PostID: request.PostID,
	}
	s.Likes[request.PostID] = append(s.Likes[request.PostID], likeRecord)

	// Optionally, log the like action if you want to track user likes
	fmt.Printf("User %d liked Post %d\n", likeRecord.UserID, likeRecord.PostID)

	// Return success message and nil error
	return "Post liked successfully", nil
}

func (s *PostServiceImpl) AddComment(postID int64, request request.CommentRequest) (int64, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	// Check if the post exists
	_, exists := s.Posts[postID]
	if !exists {
		return 0, errors.New("post not found")
	}

	// Create a new comment
	commentID := int64(len(s.Comments[postID]) + 1)
	comment := model.Comment{
		ID:              commentID,
		UserID:          request.UserID,
		CommentText:     request.CommentText,
		ParentCommentID: request.ParentCommentID,
		IsEdited:        false,
		Attachments:     request.Attachments,
		CreatedAt:       time.Now(),
	}

	// Add the comment to the post's comment list
	s.Comments[postID] = append(s.Comments[postID], comment)

	return commentID, nil
}

func (s *PostServiceImpl) GetPostDetails(postID int64) (response.PostDetail, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	// Retrieve the post details
	post, exists := s.Posts[postID]
	if !exists {
		return response.PostDetail{}, errors.New("post not found")
	}

	// Get comments for the post
	comments := s.Comments[postID]

	// Retrieve the like records for the post
	likeRecords := s.Likes[postID]
	likesCount := len(likeRecords) // Count of likes

	return response.PostDetail{
		Post:        post,
		Comments:    comments,
		LikesCount:  likesCount,
		LikeRecords: likeRecords,
	}, nil
}

func derefPrivacy(ptr *model.PrivacySettings) model.PrivacySettings {
	if ptr != nil {
		return *ptr
	}
	return model.PrivacySettings{}
}
