package models

import "time"

type Ad struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	Title        string    `json:"title"`
	AdText       string    `json:"ad_text"`
	ImageAddress string    `json:"image_address"`
	Price        float64   `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateAd struct {
	Title        string  `json:"title"`
	AdText       string  `json:"ad_text"`
	ImageAddress string  `json:"image_address"`
	Price        float64 `json:"price"`
}
