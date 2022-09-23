package model

import (
	"time"

	"gorm.io/gorm"
)

// gorm.Model definition
// gorm:"default:uuid_generate_v4()
type BaseModel struct {
	ID        string `gorm:"primaryKey,default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
