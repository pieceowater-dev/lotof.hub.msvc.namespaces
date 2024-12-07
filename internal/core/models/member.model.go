package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	ID         uuid.UUID   `gorm:"type:uuid;primaryKey"`
	UserID     uuid.UUID   `gorm:"type:uuid;primaryKey"`
	Namespaces []Namespace `gorm:"many2many:namespace_members;"` // Many-to-Many relation with Namespace
	Services   []Service   `gorm:"many2many:service_members;"`   // Many-to-Many relation with Services
}
