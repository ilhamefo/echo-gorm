package model

import (
	"time"
)

type User struct {
	BaseModel
	Name            string
	Username        string    `gorm:"unique"`
	Email           string    `gorm:"unique"`
	EmailVerifiedAt time.Time `gorm:"email_verified_at"`
	Password        string
	BackgroundColor string `gorm:"background_color;default:#FFFFFF"`
	TextColor       string `gorm:"text_color;default:#000000"`
}
