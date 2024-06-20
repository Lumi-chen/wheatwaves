package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MODEL struct {
	ID         uuid.UUID `gorm:"type:varchar(36)"`
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime gorm.DeletedAt `gorm:"index"`
}

func (u *MODEL) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
