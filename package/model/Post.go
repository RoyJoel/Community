package model

import (
	"encoding/json"

	"github.com/zhangjiacheng-iHealth/IHCommunity/package/utils"
)

type Post struct {
	Id                  int64            `json:"id"`
	Title               string           `json:"title"`
	Content             string           `json:"content"`
	Priority              int64            `json:"priority"`
	State           int64            `json:"state"`
	StartDate           float64          `json:"startDate"`
	EndDate             float64          `json:"endDate"`
}

func (Post Post) TableName() string {
	return "Post"
}

func (Post Post) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":                  Post.Id,
		"title":               Post.Title,
		"content":             Post.Content,
		"priority":              Post.Priority,
		"state":             Post.State,
		"startDate":           Post.StartDate,
		"endDate":             Post.EndDate,
	})
}

// Redis类似序列化操作
func (Post Post) MarshalBinary() ([]byte, error) {
	return json.Marshal(Post)
}

func (Post Post) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Post)
}

