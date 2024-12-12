package svc

import (
	"app/internal/pkg/ns/ent"
	"fmt"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type NSService struct {
	db gossiper.Database
}

func NewNSService(db gossiper.Database) *NSService {
	return &NSService{db: db}
}

func (s NSService) GetNamespaces(filter gossiper.Filter[string]) (*[]ent.Namespace, int64, error) {
	var namespaces []ent.Namespace
	var count int64

	query := s.db.GetDB().Model(&ent.Namespace{})

	// Apply search filters
	if filter.Search != "" {
		search := "%" + filter.Search + "%"
		query = query.Where("title LIKE ?", search, search)
	}

	// Count total records
	if err := query.Count(&count).Error; err != nil {
		return &[]ent.Namespace{}, 0, fmt.Errorf("failed to count: %w", err)
	}

	// Apply pagination
	query = query.Offset((filter.Pagination.Page - 1) * filter.Pagination.Length).Limit(filter.Pagination.Length)

	// Apply sorting dynamically
	if field := filter.Sort.Field; field != "" && gossiper.IsFieldValid(&ent.Namespace{}, field) {
		query = query.Order(fmt.Sprintf("%s %s", gossiper.ToSnakeCase(field), filter.Sort.Direction))
	}

	// Fetch data
	if err := query.Find(&namespaces).Error; err != nil {
		return &[]ent.Namespace{}, 0, fmt.Errorf("failed to fetch: %w", err)
	}

	return &namespaces, count, nil
}

func (s NSService) GetNamespace(id string) (*ent.Namespace, error) {
	var namespace *ent.Namespace

	if err := s.db.GetDB().
		Model(&ent.Namespace{}).
		Where("id = ?", id).
		First(&namespace).Error; err != nil {
		return nil, err
	}

	return namespace, nil
}

func (s NSService) CreateNamespace(input *ent.Namespace) (*ent.Namespace, error) {
	if err := s.db.GetDB().Create(input).Error; err != nil {
		return nil, fmt.Errorf("failed to create namespace: %w", err)
	}
	return input, nil
}

func (s NSService) UpdateNamespace(namespace *ent.Namespace) (*ent.Namespace, error) {
	if err := s.db.GetDB().Save(namespace).Error; err != nil {
		return nil, fmt.Errorf("failed to update namespace: %w", err)
	}
	return namespace, nil
}
