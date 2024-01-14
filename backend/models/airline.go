package models

import "time"

type Airline struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
