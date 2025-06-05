// internal/repositories/company_repo_test.go
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

// TestCompanyRepoIntegration verifica que CompanyRepo.Upsert, Get, FindAll y Delete realmente funcionen
// sobre una BD de pruebas llamada "testdb" en CockroachDB local.
func TestCompanyRepoIntegration(t *testing.T) {

	databaseURL := "postgresql://root@localhost:26257/defaultdb?sslmode=disable"
	if databaseURL == "" {
		t.Fatal("debes exportar DATABASE_URL con la conexión a tu BD de pruebas, e.g. postgresql://root@localhost:26257/testdb?sslmode=disable")
	}

	// 2) Ejecutar migraciones ANTES de abrir el pool, para crear tablas en testdb
	if err := db.RunMigrations(databaseURL); err != nil {
		t.Fatalf("migrations fallaron: %v", err)
	}

	// 3) Abrir pool de pgxpool
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		t.Fatalf("no se pudo crear pgxpool: %v", err)
	}
	defer pool.Close()

	// 4) Crear el repo
	repo := NewCompanyRepository(pool)

	// 5) Definir un Company de prueba
	company := models.Company{
		Ticker: "TEST",
		Name:   "Test Company",
	}

	//t.Log("currently company name", company.Name)

	// 6) Llamar a Upsert para insertar (ID=0 implica creación)
	err = repo.Upsert(ctx, company)
	if err != nil {
		t.Fatalf("Upsert(insert) falló: %v", err)
	}

	// 7) Ahora recuperarlo con FindAll, debería aparecer exactamente uno
	all, err := repo.FindAll(ctx)
	if err != nil {
		t.Fatalf("FindAll falló: %v", err)
	}
	if len(all) != 1 {
		t.Fatalf("esperaba 1 compañía, encontré %d", len(all))
	}

	// 8) Verificar campos y que se haya asignado un ID distinto de cero
	created := all[0]

	assert.NotZero(t, created.ID, "company.ID debe ser distinto de cero tras el insert")
	assert.Equal(t, "TEST", created.Ticker)
	assert.Equal(t, "Test Company", created.Name)

	// 9) Probar Get por ID
	fetched, err := repo.Get(ctx, strconv.FormatInt(created.ID, 10))
	if err != nil {
		t.Fatalf("Get(%d) falló: %v", created.ID, err)
	}
	assert.Equal(t, created.ID, fetched.ID)
	assert.Equal(t, created.Ticker, fetched.Ticker)
	assert.Equal(t, created.Name, fetched.Name)

	// 10) Probar Update: cambiar el nombre y volver a Upsert
	created.Name = "Test Company Updated"
	err = repo.Upsert(ctx, created)
	if err != nil {
		t.Fatalf("Upsert(update) falló: %v", err)
	}
	updated, err := repo.Get(ctx, strconv.FormatInt(created.ID, 10))
	if err != nil {
		t.Fatalf("Get después de update falló: %v", err)
	}
	assert.Equal(t, "Test Company Updated", updated.Name)

	// 11) Probar Delete
	err = repo.Delete(ctx, strconv.FormatInt(created.ID, 10))
	if err != nil {
		t.Fatalf("Delete(%d) falló: %v", created.ID, err)
	}

	// 12) Verificar que ya no exista: FindAll debe devolver 0, Get debe fallar
	all2, err := repo.FindAll(ctx)
	if err != nil {
		t.Fatalf("FindAll tras delete falló: %v", err)
	}
	if len(all2) != 0 {
		t.Fatalf("esperaba 0 compañías tras delete, encontré %d", len(all2))
	}
	_, err = repo.Get(ctx, strconv.FormatInt(created.ID, 10))
	assert.Error(t, err, "Get tras delete debería devolver error")
}
