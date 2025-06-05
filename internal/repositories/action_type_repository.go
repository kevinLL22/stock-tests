package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinLL22/stock-tests/internal/models"
)

// ActionTypeRepository defines CRUD operations for ActionType
// entities.
type ActionTypeRepository interface {
	Upsert(ctx context.Context, m models.ActionType) error
	FindAll(ctx context.Context) ([]models.ActionType, error)
	Get(ctx context.Context, id string) (models.ActionType, error)
	Delete(ctx context.Context, id string) error
}

// ActionTypeRepo implements ActionTypeRepository using pgx.
type ActionTypeRepo struct{ db *pgxpool.Pool }

func NewActionTypeRepository(db *pgxpool.Pool) ActionTypeRepository {
	return &ActionTypeRepo{db: db}
}

func (r ActionTypeRepo) Upsert(ctx context.Context, m models.ActionType) error {
	const query = `
                UPSERT INTO action_types (action_id, code, description)
                VALUES ($1, $2, $3)
        `
	_, err := r.db.Exec(ctx, query, m.ID, m.Code, m.Description)
	return err
}

func (r ActionTypeRepo) FindAll(ctx context.Context) ([]models.ActionType, error) {
	const query = `
                SELECT action_id, code, description
                FROM action_types
                ORDER BY action_id
        `
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.ActionType
	for rows.Next() {
		var at models.ActionType
		if err := rows.Scan(&at.ID, &at.Code, &at.Description); err != nil {
			return nil, err
		}
		list = append(list, at)
	}
	return list, rows.Err()
}

func (r ActionTypeRepo) Get(ctx context.Context, id string) (models.ActionType, error) {
	actionID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.ActionType{}, fmt.Errorf("invalid action type ID: %w", err)
	}
	const query = `
                SELECT action_id, code, description
                FROM action_types
                WHERE action_id = $1
        `
	var at models.ActionType
	err = r.db.QueryRow(ctx, query, actionID).Scan(&at.ID, &at.Code, &at.Description)
	if err != nil {
		return models.ActionType{}, err
	}
	return at, nil
}

func (r ActionTypeRepo) Delete(ctx context.Context, id string) error {
	actionID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid action type ID: %w", err)
	}
	const query = `
                DELETE FROM action_types
                WHERE action_id = $1
        `
	_, err = r.db.Exec(ctx, query, actionID)
	return err
}
