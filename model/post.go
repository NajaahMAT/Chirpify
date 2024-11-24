package model

import "time"

type Post struct {
	ID                   int64           `json:"id"`
	UserID               int64           `json:"user_id"`
	Message              string          `json:"message,omitempty"`                // Text content of the post
	Link                 string          `json:"link,omitempty"`                   // URL to share
	Caption              string          `json:"caption,omitempty"`                // Caption for the link
	Description          string          `json:"description,omitempty"`            // Description of the link/media
	Picture              string          `json:"picture,omitempty"`                // URL to an image
	Name                 string          `json:"name,omitempty"`                   // Title of the link or media
	Tags                 []string        `json:"tags,omitempty"`                   // Array of tagged user IDs
	Source               string          `json:"source,omitempty"`                 // File upload URL or multipart data (for media)
	FileURL              string          `json:"file_url,omitempty"`               // Alternative to source for media upload
	Privacy              PrivacySettings `json:"privacy,omitempty"`                // Privacy settings for the post
	ScheduledPublishTime int64           `json:"scheduled_publish_time,omitempty"` // Unix timestamp for scheduled post
	AllowComments        bool            `json:"allow_cmt,omitempty"`              // Whether to allow comments
	Location             *GeoLocation    `json:"location,omitempty"`               // Geolocation details
	Status               string          `json:"status"`                           // Status of the post (e.g., "published", "draft")
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
}

type Comment struct {
	ID              int64     `json:"id"`
	UserID          int64     `json:"user_id"`
	CommentText     string    `json:"comment_text"`
	ParentCommentID int64     `json:"parent_comment_id"`
	IsEdited        bool      `json:"is_edited,omitempty"`
	Attachments     []string  `json:"attachments,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type PostDetail struct {
	Post        Post         `json:"post"`
	Comments    []Comment    `json:"comments"`
	LikesCount  int          `json:"likes_count"`
	LikeRecords []LikeRecord `json:"like_records"`
}

type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Place     string  `json:"place,omitempty"` // Place ID for the location
}

type PrivacySettings struct {
	Value string `json:"value"` // Options: "PUBLIC", "FRIENDS", "ONLY_ME", "CUSTOM"
}

// LikeRecord holds the user ID and post ID for each like action.
type LikeRecord struct {
	UserID int64 `json:"user_id"`
	PostID int64 `json:"post_id"`
}
