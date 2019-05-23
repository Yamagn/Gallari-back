package Model

import "time"

type Product struct {
	Id int `gorm:"primary key" json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Link        string `json:"link"`
	IsPublic    bool   `json:"is_public"`
	CreatedUserId      int  `json:"created_user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedUt time.Time `json:"updated_ut"`
}
