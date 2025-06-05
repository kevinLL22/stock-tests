package repositories

import (
	"context"
	"strconv"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinLL22/stock-tests/internal/db"
	"github.com/kevinLL22/stock-tests/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestRatingTypeRepoIntegration(t *testing.T) {
	databaseURL := "postgresql://root@localhost:26257/defaultdb?sslmode=disable"
	if databaseURL == "" {
		t.Fatal("debes exportar DATABASE_URL con la conexión a tu BD de pruebas, e.g. postgresql://root@localhost:26257/testdb?sslmode=disable")
	}

	if err := db.RunMigrations(databaseURL); err != nil {
		t.Fatalf("migrations fallaron: %v", err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		t.Fatalf("no se pudo crear pgxpool: %v", err)
	}
	defer pool.Close()

	repo := NewRatingTypeRepository(pool)

	prev, err := repo.FindAll(ctx)
	if err != nil {
		t.Fatalf("FindAll inicial falló: %v", err)
	}
	baseCount := len(prev)

	rating := models.RatingType{
		ID:          67890,
		Code:        "TEST_RATING",
		Description: "Test Rating",
	}

	err = repo.Upsert(ctx, rating)
	if err != nil {
		t.Fatalf("Upsert(insert) falló: %v", err)
	}

	all, err := repo.FindAll(ctx)
	if err != nil {
		t.Fatalf("FindAll falló: %v", err)
	}
	if len(all) != baseCount+1 {
		t.Fatalf("esperaba %d rating types, encontré %d", baseCount+1, len(all))
	}

	created, err := repo.Get(ctx, strconv.FormatInt(rating.ID, 10))
	if err != nil {
		t.Fatalf("Get(%d) falló: %v", rating.ID, err)
	}
	assert.Equal(t, rating.Code, created.Code)
	assert.Equal(t, rating.Description, created.Description)

	created.Description = "Updated Rating"
	err = repo.Upsert(ctx, created)
	if err != nil {
		t.Fatalf("Upsert(update) falló: %v", err)
	}
	updated, err := repo.Get(ctx, strconv.FormatInt(rating.ID, 10))
	if err != nil {
		t.Fatalf("Get después de update falló: %v", err)
	}
	assert.Equal(t, "Updated Rating", updated.Description)

	err = repo.Delete(ctx, strconv.FormatInt(rating.ID, 10))
	if err != nil {
		t.Fatalf("Delete(%d) falló: %v", rating.ID, err)
	}

	all2, err := repo.FindAll(ctx)
	if err != nil {
		t.Fatalf("FindAll tras delete falló: %v", err)
	}
	if len(all2) != baseCount {
		t.Fatalf("esperaba %d rating types tras delete, encontré %d", baseCount, len(all2))
	}
	_, err = repo.Get(ctx, strconv.FormatInt(rating.ID, 10))
	assert.Error(t, err, "Get tras delete debería devolver error")
}
