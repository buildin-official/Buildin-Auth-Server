package models

import (
	"time"

	guuid "github.com/google/uuid"
)

type User struct {
	ID        guuid.UUID `gorm:"primaryKey" json:"-"`
	Email     string     `gorm:"not null;size:320" json:"email"`
	Password  string     `gorm:"not null;size:500" json:"-"`
	Username  string     `gorm:"not null;size:20" json:"username"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"-" `
}

type UnvalidatedUser struct {
	ID        guuid.UUID `gorm:"primaryKey" json:"-"`
	Email     string     `gorm:"not null;size:320" json:"email"`
	Password  string     `gorm:"not null;size:500" json:"-"`
	Username  string     `gorm:"not null;size:20" json:"username"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"-" `
}

type ChangePasswordCode struct {
	ID        guuid.UUID `gorm:"primaryKey" json:"-"`
	UserID    guuid.UUID `gorm:"unique; not null" json:"-"`
	Password  string     `gorm:"not null;size:500" json:"-"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"-" `
}
