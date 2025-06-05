package services

import (
	"context"
	"fmt"

	"github.com/kevinLL22/stock-tests/internal/models"
	"github.com/kevinLL22/stock-tests/internal/repositories"
)

// RatingTypeSvc provides business logic for rating types.
type RatingTypeSvc struct {
	repo repositories.RatingTypeRepository
}

func NewRatingTypeService(repo repositories.RatingTypeRepository) *RatingTypeSvc {
	return &RatingTypeSvc{repo: repo}
}

func (s *RatingTypeSvc) CreateOrUpdate(ctx context.Context, m models.RatingType) error {
	if m.Code == "" {
		return fmt.Errorf("code cannot be empty")
	}
	return s.repo.Upsert(ctx, m)
}

func (s *RatingTypeSvc) ListAll(ctx context.Context) ([]models.RatingType, error) {
	return s.repo.FindAll(ctx)
}

func (s *RatingTypeSvc) GetByID(ctx context.Context, id string) (models.RatingType, error) {
	return s.repo.Get(ctx, id)
}

func (s *RatingTypeSvc) DeleteByID(ctx context.Context, id string) error {
	if _, err := s.repo.Get(ctx, id); err != nil {
		return fmt.Errorf("there is not a rating type with id %s: %w", id, err)
	}
	return s.repo.Delete(ctx, id)
}
