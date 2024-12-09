package ent

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NSMember struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

type Namespace struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string
	Slug        string
	Description string
	Owner       uuid.UUID   `gorm:"type:uuid"`
	Members     []*NSMember `gorm:"many2many:namespace_members;"`
}
