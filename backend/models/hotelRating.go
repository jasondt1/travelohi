package models

import "time"

type HotelRating struct {
	ID            uint   `json:"id"`
	Type          string `json:"type"`
	Value         int    `json:"value"`
	HotelReviewID uint   `json:"reviewId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Review HotelReview `gorm:"foreignKey:HotelReviewID"`
}
