package models

import "time"

type RoomType struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
