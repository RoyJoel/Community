package model

import (
	"encoding/json"
)

type Privileges struct {
	Id                int64   `json:"id"`
	ProposalId int64   `json:"proposalId"`
	Content          string   `json:"content"`
}

func (Privileges Privileges) TableName() string {
	return "Privileges"
}

func (Privileges Privileges) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":                Privileges.Id,
		"proposalId":          Privileges.ProposalId,
		"content":           Privileges.Content,
	})
}

// Redis类似序列化操作
func (Privileges Privileges) MarshalBinary() ([]byte, error) {
	return json.Marshal(Privileges)
}

func (Privileges Privileges) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Privileges)
}

