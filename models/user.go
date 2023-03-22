package models

import (
	"time"

	guuid "github.com/google/uuid"
)

type User struct {
	ID        guuid.UUID `gorm:"primaryKey" json:"-"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"-" `
}

type UnvalidatedUser struct {
	ID        guuid.UUID `gorm:"primaryKey" json:"-"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"-" `
}
