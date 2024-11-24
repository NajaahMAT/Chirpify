package service

import (
	"chirpify/data/request"
	"chirpify/data/response"
	"chirpify/model"
)

type PostService interface {
	CreatePost(req request.CreatePostRequest) (int64, error)
	UpdatePost(postID int64, req request.UpdatePostRequest) error
	GetAllPosts() ([]model.Post, error)
	LikePost(userID, postID int64) (string, error)
	AddComment(postID int64, request request.CommentRequest) (int64, error)
	GetPostDetails(postID int64) (response.PostDetail, error)
}
