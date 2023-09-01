package models

import (
	"time"

	"github.com/google/uuid"
)

type UUIDBaseModel struct {
	UUID      uuid.UUID `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
