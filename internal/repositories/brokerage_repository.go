package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinLL22/stock-tests/internal/models"
	"strconv"
)

type BrokerageRepository interface {
	Upsert(ctx context.Context, brokerageModel models.Brokerage) error
	FindAll(ctx context.Context) ([]models.Brokerage, error)
	Get(ctx context.Context, id string) (models.Brokerage, error)
	Delete(ctx context.Context, id string) error
}

type BrokerageRepo struct{ db *pgxpool.Pool }

func NewBrokerageRepository(db *pgxpool.Pool) BrokerageRepository {
	return &BrokerageRepo{db: db}
}

func (bRepo BrokerageRepo) Upsert(ctx context.Context, brokerageModel models.Brokerage) error {
	const query = `
                UPSERT INTO brokerages (brokerage_id, name)
                VALUES ($1, $2)
        `
	_, err := bRepo.db.Exec(ctx, query, brokerageModel.ID, brokerageModel.Name)
	return err
}

func (bRepo BrokerageRepo) FindAll(ctx context.Context) ([]models.Brokerage, error) {
	const query = `
                SELECT brokerage_id, name
                FROM brokerages
                ORDER BY brokerage_id
        `
	rows, err := bRepo.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brokerageModelsList []models.Brokerage
	for rows.Next() {
		var brokerage models.Brokerage
		if err := rows.Scan(&brokerage.ID, &brokerage.Name); err != nil {
			return nil, err
		}
		brokerageModelsList = append(brokerageModelsList, brokerage)
	}
	return brokerageModelsList, rows.Err()
}

func (bRepo BrokerageRepo) Get(ctx context.Context, id string) (models.Brokerage, error) {
	brokerageID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.Brokerage{}, fmt.Errorf("invalid brokerage ID: %w", err)
	}

	const query = `
                SELECT brokerage_id, name
                FROM brokerages
                WHERE brokerage_id = $1
        `
	var brokerage models.Brokerage
	err = bRepo.db.QueryRow(ctx, query, brokerageID).Scan(&brokerage.ID, &brokerage.Name)
	if err != nil {
		return models.Brokerage{}, err
	}
	return brokerage, nil
}

func (bRepo BrokerageRepo) Delete(ctx context.Context, id string) error {
	brokerageID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid brokerage ID: %w", err)
	}

	const query = `
                DELETE FROM brokerages
                WHERE brokerage_id = $1
        `
	_, err = bRepo.db.Exec(ctx, query, brokerageID)
	return err
}
