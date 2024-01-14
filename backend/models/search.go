package models

import "time"

type Search struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"userID"`
	Keyword string `json:"keyword"`
	Type    string `json:"type"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	User User `gorm:"foreignKey:UserID"`
}
