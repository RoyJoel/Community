package model

import (
	"encoding/json"

	"github.com/zhangjiacheng-iHealth/IHCommunity/package/utils"
)

type Comment struct {
	Id        int64            `json:"id"`
	UserId      int64           `json:"userId"`
	PostId      int64           `json:"postId"`
	Date float64          `json:"date"`
	Photos   utils.IntMatrix           `json:"photos"`
	Videos     utils.IntMatrix            `json:"videos"`
	Content      utils.IntMatrix  `json:"content"`
}

func (Comment Comment) TableName() string {
	return "Comment"
}

func (Comment Comment) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":        Comment.Id,
		"userId":      Comment.UserId,
		"postId":      Comment.PostId,
		"date": Comment.Date,
		"photos":   Comment.Photos,
		"videos":     Comment.Videos,
		"content":      Comment.Content,
	})
}

// Redis类似序列化操作
func (Comment Comment) MarshalBinary() ([]byte, error) {
	return json.Marshal(Comment)
}

func (Comment Comment) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Comment)
}