package main

import (
	"github.com/kevinLL22/stock-tests/internal/db"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/kevinLL22/stock-tests/internal/config"
	"github.com/kevinLL22/stock-tests/internal/controllers"
	"github.com/kevinLL22/stock-tests/internal/repositories"
	"github.com/kevinLL22/stock-tests/internal/services"
)

func main() {
	// 1. Leer configuración (puerto, conexión, etc.)
	cfg := config.Load()

	// 2. Inicializar conexión a CockroachDB (todavía placeholder)
	pool, err := db.NewPool(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("db init: %v", err)
	}
	defer pool.Close()

	// 3. Instanciar capa de servicio (repo + URL remota)
	stockRepo := repositories.NewStockRepository(pool)
	stockSvc := services.NewStockService(
		stockRepo,
		"https://example.com/stocks-endpoint.json", // TODO: move to cfg
	)

	// 4. Configurar router
	r := gin.Default()
	controllers.RegisterRoutes(r, stockSvc)

	// 5. Arrancar
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server: %v", err)
	}
}
