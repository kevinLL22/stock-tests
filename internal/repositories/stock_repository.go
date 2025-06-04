package repositories

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinLL22/stock-tests/internal/models"
)

type StockRepository interface {
	Upsert(ctx context.Context, s models.Stock) error
	FindAll(ctx context.Context) ([]models.Stock, error)
	Get(ctx context.Context, id string) (models.Stock, error)
	Delete(ctx context.Context, id string) error
}

type stockRepo struct{ db *pgxpool.Pool }

func NewStockRepository(db *pgxpool.Pool) StockRepository {
	return &stockRepo{db: db}
}

// Métodos: implementations TBD
func (r *stockRepo) Upsert(ctx context.Context, s models.Stock) error {
	return nil
}
func (r *stockRepo) FindAll(ctx context.Context) ([]models.Stock, error) {
	return nil, nil
}
func (r *stockRepo) Get(ctx context.Context, id string) (models.Stock, error) {
	return models.Stock{}, nil
}
func (r *stockRepo) Delete(ctx context.Context, id string) error {
	return nil
}

// Implementations of StockRepository methods will go here.
