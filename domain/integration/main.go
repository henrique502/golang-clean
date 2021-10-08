package integration

import "time"

const TableName = "integrations"

type Integration struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Enabled   bool   `json:"enabled"`
	Type      string `json:"type"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New() Integration {
	return Integration{}
}
