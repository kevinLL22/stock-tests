package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinLL22/stock-tests/internal/models"
)

// RatingTypeRepository defines CRUD operations for RatingType
// entities.
type RatingTypeRepository interface {
	Upsert(ctx context.Context, m models.RatingType) error
	FindAll(ctx context.Context) ([]models.RatingType, error)
	Get(ctx context.Context, id string) (models.RatingType, error)
	Delete(ctx context.Context, id string) error
}

// RatingTypeRepo implements RatingTypeRepository using pgx.
type RatingTypeRepo struct{ db *pgxpool.Pool }

func NewRatingTypeRepository(db *pgxpool.Pool) RatingTypeRepository {
	return &RatingTypeRepo{db: db}
}

func (r RatingTypeRepo) Upsert(ctx context.Context, m models.RatingType) error {
	const query = `
                UPSERT INTO rating_types (rating_id, code, description)
                VALUES ($1, $2, $3)
        `
	_, err := r.db.Exec(ctx, query, m.ID, m.Code, m.Description)
	return err
}

func (r RatingTypeRepo) FindAll(ctx context.Context) ([]models.RatingType, error) {
	const query = `
                SELECT rating_id, code, description
                FROM rating_types
                ORDER BY rating_id
        `
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.RatingType
	for rows.Next() {
		var rt models.RatingType
		if err := rows.Scan(&rt.ID, &rt.Code, &rt.Description); err != nil {
			return nil, err
		}
		list = append(list, rt)
	}
	return list, rows.Err()
}

func (r RatingTypeRepo) Get(ctx context.Context, id string) (models.RatingType, error) {
	ratingID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.RatingType{}, fmt.Errorf("invalid rating type ID: %w", err)
	}
	const query = `
                SELECT rating_id, code, description
                FROM rating_types
                WHERE rating_id = $1
        `
	var rt models.RatingType
	err = r.db.QueryRow(ctx, query, ratingID).Scan(&rt.ID, &rt.Code, &rt.Description)
	if err != nil {
		return models.RatingType{}, err
	}
	return rt, nil
}

func (r RatingTypeRepo) Delete(ctx context.Context, id string) error {
	ratingID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid rating type ID: %w", err)
	}
	const query = `
                DELETE FROM rating_types
                WHERE rating_id = $1
        `
	_, err = r.db.Exec(ctx, query, ratingID)
	return err
}
