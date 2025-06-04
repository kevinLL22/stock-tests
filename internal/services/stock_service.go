package services

import (
	"context"
	"github.com/kevinLL22/stock-tests/internal/models"
	"github.com/kevinLL22/stock-tests/internal/repositories"
)

type StockService interface {
	RefreshFromAPI(ctx context.Context) error
	BestToBuy(ctx context.Context) (models.Stock, error)
	repositories.StockRepository // embed CRUD
}

type stockSvc struct {
	repositories.StockRepository
	remoteURL string
}

func NewStockService(repo repositories.StockRepository, remoteURL string) StockService {
	return &stockSvc{repo, remoteURL}
}

// Métodos: TBD

func (s *stockSvc) RefreshFromAPI(ctx context.Context) error {
	// Aquí iría la lógica para refrescar los datos desde la API remota.
	// Por ahora, solo un placeholder.
	return nil
}

func (s *stockSvc) BestToBuy(ctx context.Context) (models.Stock, error) {
	// Aquí iría la lógica para determinar cuál es la mejor acción para comprar.
	// Por ahora, solo un placeholder que devuelve un stock vacío.
	return models.Stock{}, nil
}

// Implementations of StockService methods will go here.
