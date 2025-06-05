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

func TestBrokerageRepoIntegration(t *testing.T) {
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

	repo := NewBrokerageRepository(pool)

	prev, err := repo.FindAll(ctx)
	if err != nil {
		t.Fatalf("FindAll inicial falló: %v", err)
	}
	baseCount := len(prev)

	brokerage := models.Brokerage{
		ID:   54321,
		Name: "Test Brokerage",
	}

	err = repo.Upsert(ctx, brokerage)
	if err != nil {
		t.Fatalf("Upsert(insert) falló: %v", err)
	}

	all, err := repo.FindAll(ctx)
	if err != nil {
		t.Fatalf("FindAll falló: %v", err)
	}
	if len(all) != baseCount+1 {
		t.Fatalf("esperaba %d brokerages, encontré %d", baseCount+1, len(all))
	}

	created, err := repo.Get(ctx, strconv.FormatInt(brokerage.ID, 10))
	if err != nil {
		t.Fatalf("Get(%d) falló: %v", brokerage.ID, err)
	}
	assert.Equal(t, brokerage.Name, created.Name)

	created.Name = "Updated Brokerage"
	err = repo.Upsert(ctx, created)
	if err != nil {
		t.Fatalf("Upsert(update) falló: %v", err)
	}
	updated, err := repo.Get(ctx, strconv.FormatInt(brokerage.ID, 10))
	if err != nil {
		t.Fatalf("Get después de update falló: %v", err)
	}
	assert.Equal(t, "Updated Brokerage", updated.Name)

	err = repo.Delete(ctx, strconv.FormatInt(brokerage.ID, 10))
	if err != nil {
		t.Fatalf("Delete(%d) falló: %v", brokerage.ID, err)
	}

	all2, err := repo.FindAll(ctx)
	if err != nil {
		t.Fatalf("FindAll tras delete falló: %v", err)
	}
	if len(all2) != baseCount {
		t.Fatalf("esperaba %d brokerages tras delete, encontré %d", baseCount, len(all2))
	}
	_, err = repo.Get(ctx, strconv.FormatInt(brokerage.ID, 10))
	assert.Error(t, err, "Get tras delete debería devolver error")
}
