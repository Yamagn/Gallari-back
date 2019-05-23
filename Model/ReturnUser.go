package Model

import "time"

type ReturnUser struct {
	Id int `json:"id"`
	DisplayName string `json:"display_name"`
	UserId      string `json:"user_id"`
	Email       string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

