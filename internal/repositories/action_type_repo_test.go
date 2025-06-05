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

func TestActionTypeRepoIntegration(t *testing.T) {
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

	repo := NewActionTypeRepository(pool)

	prev, err := repo.FindAll(ctx)
	if err != nil {
		t.Fatalf("FindAll inicial falló: %v", err)
	}
	baseCount := len(prev)

	action := models.ActionType{
		ID:          12345,
		Code:        "TEST_ACTION",
		Description: "Test Action",
	}

	err = repo.Upsert(ctx, action)
	if err != nil {
		t.Fatalf("Upsert(insert) falló: %v", err)
	}

	all, err := repo.FindAll(ctx)
	if err != nil {
		t.Fatalf("FindAll falló: %v", err)
	}
	if len(all) != baseCount+1 {
		t.Fatalf("esperaba %d action types, encontré %d", baseCount+1, len(all))
	}

	created, err := repo.Get(ctx, strconv.FormatInt(action.ID, 10))
	if err != nil {
		t.Fatalf("Get(%d) falló: %v", action.ID, err)
	}
	assert.Equal(t, action.Code, created.Code)
	assert.Equal(t, action.Description, created.Description)

	created.Description = "Updated Action"
	err = repo.Upsert(ctx, created)
	if err != nil {
		t.Fatalf("Upsert(update) falló: %v", err)
	}
	updated, err := repo.Get(ctx, strconv.FormatInt(action.ID, 10))
	if err != nil {
		t.Fatalf("Get después de update falló: %v", err)
	}
	assert.Equal(t, "Updated Action", updated.Description)

	err = repo.Delete(ctx, strconv.FormatInt(action.ID, 10))
	if err != nil {
		t.Fatalf("Delete(%d) falló: %v", action.ID, err)
	}

	all2, err := repo.FindAll(ctx)
	if err != nil {
		t.Fatalf("FindAll tras delete falló: %v", err)
	}
	if len(all2) != baseCount {
		t.Fatalf("esperaba %d action types tras delete, encontré %d", baseCount, len(all2))
	}
	_, err = repo.Get(ctx, strconv.FormatInt(action.ID, 10))
	assert.Error(t, err, "Get tras delete debería devolver error")
}
