package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Store struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	StoreName string         `gorm:"size:255;not null"`
	Location  string         `gorm:"size:255"`
}

func (Store) TableName() string {
	return "store"
}
