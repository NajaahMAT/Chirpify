package response

import "chirpify/model"

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type PostDetail struct {
	Post        model.Post         `json:"post"`
	Comments    []model.Comment    `json:"comments"`
	LikesCount  int                `json:"likes_count"`
	LikeRecords []model.LikeRecord `json:"like_records"`
}
