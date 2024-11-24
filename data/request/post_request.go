package request

import (
	"chirpify/model"
)

// UpdatePostRequest represents the request payload for updating a post.
type UpdatePostRequest struct {
	Message              *string                `json:"message,omitempty"` // Pointer to differentiate between no update and clearing the value
	Link                 *string                `json:"link,omitempty"`
	Caption              *string                `json:"caption,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Picture              *string                `json:"picture,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Tags                 *[]string              `json:"tags,omitempty"`
	Source               *string                `json:"source,omitempty"`
	FileURL              *string                `json:"file_url,omitempty"`
	Privacy              *model.PrivacySettings `json:"privacy,omitempty"`
	ScheduledPublishTime *int64                 `json:"scheduled_publish_time,omitempty"`
	AllowComments        *bool                  `json:"allow_cmt,omitempty"`
	Location             *model.GeoLocation     `json:"location,omitempty"`
	Status               *string                `json:"status,omitempty"` // "published", "draft"
}

// CreatePostRequest represents the request payload for creating a post.
type CreatePostRequest struct {
	UserID               int                    `json:"user_id" binding:"required"`
	Message              *string                `json:"message,omitempty"` // Pointer to differentiate between no update and clearing the value
	Link                 *string                `json:"link,omitempty"`
	Caption              *string                `json:"caption,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Picture              *string                `json:"picture,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Tags                 *[]string              `json:"tags,omitempty"`
	Source               *string                `json:"source,omitempty"`
	FileURL              *string                `json:"file_url,omitempty"`
	Privacy              *model.PrivacySettings `json:"privacy,omitempty"`
	ScheduledPublishTime *int64                 `json:"scheduled_publish_time,omitempty"`
	AllowComments        *bool                  `json:"allow_cmt,omitempty"`
	Location             *model.GeoLocation     `json:"location,omitempty"`
	Status               *string                `json:"status,omitempty"` // "published", "draft"
}

// CommentRequest represents the request payload for adding or updating a comment on a post.
type CommentRequest struct {
	UserID          int64    `json:"user_id" binding:"required"`      // ID of the user adding the comment
	CommentText     string   `json:"comment_text" binding:"required"` // The content of the comment
	ParentCommentID int64    `json:"parent_comment_id,omitempty"`
	IsEdited        bool     `json:"is_edited,omitempty"`
	Attachments     []string `json:"attachments,omitempty"`
}

type LikeRequest struct {
	PostID int64 `json:"post_id"`
	UserID int64 `json:"user_id"`
}
