package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Ads       []Ad      `json:"ads"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterUserModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
