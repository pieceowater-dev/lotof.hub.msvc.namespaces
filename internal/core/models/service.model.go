package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	App         string
	NamespaceID uuid.UUID // Foreign key for Namespace
	Namespace   Namespace `gorm:"constraint:OnDelete:CASCADE;"` // Relation to Namespace
	Members     []Member  `gorm:"many2many:service_members;"`   // Many-to-Many relation with Member
}
