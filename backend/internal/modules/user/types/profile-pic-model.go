package types

import (
	"time"
)

type ProfilePicture struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `gorm:"not null" json:"user_id"`
	ImageData []byte    `gorm:"type:bytea;not null" json:"image_data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
