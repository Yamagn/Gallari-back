package Model

import "time"

type User struct {
	Id int `gorm:"primary key" json:"id"`
	DisplayName string `gorm:"size:128; not null" json:"display_name"`
	UserId      string `gorm:"unique; not null" json:"user_id"`
	Email       string `gorm:"not null" json:"email"`
	Password    string `gorm:"not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
