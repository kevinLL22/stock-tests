package services

import (
	"context"
	"fmt"

	"github.com/kevinLL22/stock-tests/internal/models"
	"github.com/kevinLL22/stock-tests/internal/repositories"
)

// ActionTypeSvc provides business logic for action types.
type ActionTypeSvc struct {
	repo repositories.ActionTypeRepository
}

func NewActionTypeService(repo repositories.ActionTypeRepository) *ActionTypeSvc {
	return &ActionTypeSvc{repo: repo}
}

func (s *ActionTypeSvc) CreateOrUpdate(ctx context.Context, m models.ActionType) error {
	if m.Code == "" {
		return fmt.Errorf("code cannot be empty")
	}
	return s.repo.Upsert(ctx, m)
}

func (s *ActionTypeSvc) ListAll(ctx context.Context) ([]models.ActionType, error) {
	return s.repo.FindAll(ctx)
}

func (s *ActionTypeSvc) GetByID(ctx context.Context, id string) (models.ActionType, error) {
	return s.repo.Get(ctx, id)
}

func (s *ActionTypeSvc) DeleteByID(ctx context.Context, id string) error {
	if _, err := s.repo.Get(ctx, id); err != nil {
		return fmt.Errorf("there is not an action type with id %s: %w", id, err)
	}
	return s.repo.Delete(ctx, id)
}
