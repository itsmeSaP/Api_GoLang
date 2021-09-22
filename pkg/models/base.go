package models

import (
	"time"
)

type Base struct {
	//giving tags with struct
	ID        string    `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
