package model

import (
	"encoding/json"

	"github.com/zhangjiacheng-iHealth/IHCommunity/package/utils"
)

type VersionLog struct {
	Id                 int64              `json:"id"`
	PostId             int64  `json:"postId"`
	Date          float64             `json:"date"`
	BuildNum           string             `json:"buildNum"`
	IosUrl               string             `json:"iosUrl"`
	AndroidUrl               string             `json:"androidUrl"`
}

func (VersionLog VersionLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":                 VersionLog.Id,
		"postId": VersionLog.PostId,
		"date":          VersionLog.LoginName,
		"buildNum":           VersionLog.BuildNum,
		"iosUrl":               VersionLog.IosUrl,
		"androidUrl":               VersionLog.AndroidUrl,
	})
}

// Redis类似序列化操作
func (VersionLog VersionLog) MarshalBinary() ([]byte, error) {
	return json.Marshal(VersionLog)
}

func (VersionLog VersionLog) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &VersionLog)
}
