package services

import (
	"context"
	"fmt"
	"github.com/kevinLL22/stock-tests/internal/models"
	"github.com/kevinLL22/stock-tests/internal/repositories"
)

type BrokerageSvc struct {
	repo repositories.BrokerageRepository
}

func NewBrokerageService(repo repositories.BrokerageRepository) *BrokerageSvc {
	return &BrokerageSvc{repo: repo}
}

func (svc *BrokerageSvc) CreateOrUpdate(ctx context.Context, brokerageModel models.Brokerage) error {
	if brokerageModel.Name == "" {
		return fmt.Errorf("need to provide a name for the brokerage")
	}
	return svc.repo.Upsert(ctx, brokerageModel)
}

func (svc *BrokerageSvc) ListAll(ctx context.Context) ([]models.Brokerage, error) {
	return svc.repo.FindAll(ctx)
}

func (svc *BrokerageSvc) GetByID(ctx context.Context, id string) (models.Brokerage, error) {
	return svc.repo.Get(ctx, id)
}

func (svc *BrokerageSvc) DeleteByID(ctx context.Context, id string) error {
	if _, err := svc.repo.Get(ctx, id); err != nil {
		return fmt.Errorf("there is not a brokerage with id %s: %w", id, err)
	}
	return svc.repo.Delete(ctx, id)
}
