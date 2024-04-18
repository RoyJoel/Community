package model

import (
	"encoding/json"
)

type Proposal struct {
	PId    int64  `json:"pId"`
	Title  string `json:"title"`
	Des string `json:"des"`
}

func (Proposal Proposal) TableName() string {
	return "Proposal"
}

func (Proposal Proposal) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":    Proposal.Id,
		"title":  Proposal.Title,
		"des": Proposal.Des,
	})
}

// Redis类似序列化操作
func (Proposal Proposal) MarshalBinary() ([]byte, error) {
	return json.Marshal(Proposal)
}

func (Proposal Proposal) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Proposal)
}