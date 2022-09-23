package model

type Links struct {
	BaseModel
	UserId string
	User   User   `gorm:"foreignKey:user_id"`
	Name   string `gorm:"name"`
	Link   string `gorm:"link"`
}
