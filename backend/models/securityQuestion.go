package models

import "time"

type SecurityQuestion struct {
	ID       uint   `json:"id"`
	Question string `json:"question"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
