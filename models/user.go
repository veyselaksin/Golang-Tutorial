package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID `gorm:"primaryKey" json:"uuid"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time
}
