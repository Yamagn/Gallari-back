package Model

import "time"

type User struct {
	Id          int    `gorm:"primary_key" json:"id"`
	DisplayName string `gorm:"size:128" json:"display_name"`
	UserId      string `gorm:"unique" json:"user_id"`
	Email       string `gorm:"auto_now" json:"email"`
	Password    string
	CreatedAt time.Time `json:"createdAt" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAT" sql:"DEFAULT:current_timestamp on update current_timestamp"`
}
