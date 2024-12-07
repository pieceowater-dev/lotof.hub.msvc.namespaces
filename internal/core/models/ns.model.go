package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Namespace struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string
	Slug        string
	Description string
	Owner       uuid.UUID `gorm:"type:uuid"`                    // User ID
	Services    []Service `gorm:"constraint:OnDelete:CASCADE;"` // Relation to Services
	Members     []Member  `gorm:"many2many:namespace_members;"` // Many-to-Many relation with Members
}
