package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MODEL struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime gorm.DeletedAt `gorm:"index"`
}
