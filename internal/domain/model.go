package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID      `json:"id" gorm:"type:char(36);primary_key"`
	CreatedAt *time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" binding:"-"`
}
