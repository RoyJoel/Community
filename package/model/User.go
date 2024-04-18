package model

import (
	"encoding/json"
)

type User struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	Avatar             string `json:"avatar"`
	Role     int64 `json:"role"`
}

func (User User) TableName() string {
	return "User"
}

func (User User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":              User.Id,
		"name":            User.Name,
		"avatar":             User.Avatar,
		"role":     User.Role,
	})
}

// Redis类似序列化操作
func (User User) MarshalBinary() ([]byte, error) {
	return json.Marshal(User)
}

func (User User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &User)
}