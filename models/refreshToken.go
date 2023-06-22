package models

import (
	"time"

	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        guuid.UUID `gorm:"primaryKey" json:"-"`
	UserID    guuid.UUID `gorm:"not null" json:"-"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"-" `
}

func FindRefreshToken(db *gorm.DB, refreshToken guuid.UUID) (error, RefreshToken) {
	foundRefreshToken := RefreshToken{}
	refreshTokenQuery := RefreshToken{ID: refreshToken}
	if err := db.First(&foundRefreshToken, &refreshTokenQuery).Error; err == gorm.ErrRecordNotFound {
		return err, RefreshToken{}
	}
	return nil, foundRefreshToken
}
