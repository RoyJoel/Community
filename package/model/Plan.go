package model

import (
	"encoding/json"
)

type Plan struct {
	Id                int64   `json:"id"`
	ProposalId int64   `json:"proposalId"`
	Content          string   `json:"content"`
}

func (Plan Plan) TableName() string {
	return "Plan"
}

func (Plan Plan) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":                Plan.Id,
		"proposalId":          Plan.ProposalId,
		"content":           Plan.Content,
	})
}

// Redis类似序列化操作
func (Plan Plan) MarshalBinary() ([]byte, error) {
	return json.Marshal(Plan)
}

func (Plan Plan) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Plan)
}

