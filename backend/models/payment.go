package models

import "time"

type Payment struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
