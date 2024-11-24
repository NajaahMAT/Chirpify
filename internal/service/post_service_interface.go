package service

import (
	"chirpify/http/request"
	"chirpify/http/response"
	"chirpify/internal/model"
)

type PostService interface {
	CreatePost(req request.CreatePostRequest) (int64, error)
	UpdatePost(postID int64, req request.UpdatePostRequest) error
	GetAllPosts() ([]model.Post, error)
	LikePost(request request.LikeRequest) (string, error)
	AddComment(postID int64, request request.CommentRequest) (int64, error)
	GetPostDetails(postID int64) (response.PostDetail, error)
}
