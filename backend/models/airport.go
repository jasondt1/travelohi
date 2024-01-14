package models

import "time"

type Airport struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	CityID uint   `json:"cityId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	City City `gorm:"foreignKey:CityID"`
}
