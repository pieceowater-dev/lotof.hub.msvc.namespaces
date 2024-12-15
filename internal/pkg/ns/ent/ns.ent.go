package ent

import (
	"github.com/axgle/mahonia"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

//user -> []member -> []namespace

type NSMember struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Namespace uuid.UUID
}

func cleanSlug(input string) string {
	// Transliterate non-English characters to Latin letters (e.g., Cyrillic to Latin)
	encoder := mahonia.NewEncoder("utf-8")
	transliterated := encoder.ConvertString(input)

	// Replace non-English letters, numbers, and symbols with hyphens
	re := regexp.MustCompile("[^a-zA-Z]+")
	cleaned := re.ReplaceAllString(transliterated, "-")

	// Remove leading/trailing hyphens
	cleaned = strings.Trim(cleaned, "-")

	// If the cleaned slug is empty, return a default slug
	if cleaned == "" {
		return "default"
	}

	// Convert to lowercase
	return strings.ToLower(cleaned)
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

func (ns *Namespace) BeforeCreate(tx *gorm.DB) (err error) {
	ns.Slug = cleanSlug(ns.Slug) // Clean the slug before saving
	return nil
}
