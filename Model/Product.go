package Model

import "time"

type Product struct {
	ProductId          int    `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Link        string `json:"link"`
	IsPublic    bool   `json:"is_public"`
	CreatedUserId      int  `json:"created_user_id"`
	CreatedAt time.Time `json:"created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" sql:"DEFAULT:current_timestamp on update current_timestamp"`
}
