package model

import "encoding/json"

type Class struct {
	ID         int64  `json:"id"`
	ClassName  string `json:"className"`
	TotalVotes int    `json:"total_votes" default:"0"`
}

type ClassWithoutTotalVotes struct {
	Class
}

func (c ClassWithoutTotalVotes) MarshalJSON() ([]byte, error) {
	type Alias struct {
		ID        int64  `json:"id"`
		ClassName string `json:"className"`
	}
	return json.Marshal(Alias{ID: c.Class.ID, ClassName: c.Class.ClassName})
}
