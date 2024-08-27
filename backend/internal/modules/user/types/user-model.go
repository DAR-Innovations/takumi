package types

import "time"

type User struct {
	ID             int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username       string    `gorm:"unique;not null" json:"username"`
	FirstName      string    `gorm:"not null" json:"firstName"`
	LastName       string    `gorm:"not null" json:"lastName"`
	Email          string    `gorm:"unique;not null" json:"email"`
	Gender         string    `json:"gender"`
	BirthDate      time.Time `json:"birthDate"`
	Password       string    `gorm:"not null" json:"password"`
	Role           string    `gorm:"default:USER" json:"role"`
	CreatedAt      time.Time `json:"createdAt"`
	Coins          int       `gorm:"default:0" json:"coins"`
	ProfilePicture string    `gorm:"default:''" json:"profilePicture"`
}
