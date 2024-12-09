package ent

import (
	"app/internal/pkg/ns/ent"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Member структура
type Member struct {
	gorm.Model
	ID         uuid.UUID        `gorm:"type:uuid;primaryKey"`
	UserID     uuid.UUID        `gorm:"type:uuid"`
	Namespaces []*ent.Namespace `gorm:"many2many:namespace_members;"`
}
