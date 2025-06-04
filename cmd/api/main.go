package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinLL22/stock-tests/internal/config"
	"github.com/kevinLL22/stock-tests/internal/controllers"
	"github.com/kevinLL22/stock-tests/internal/db"
	"github.com/kevinLL22/stock-tests/internal/repositories"
	"github.com/kevinLL22/stock-tests/internal/services"
	"log"
)

func main() {

	cfg := config.Load()

	// inits

	// pool
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}
	defer pool.Close()

	if err := db.RunMigrations(cfg.DatabaseURL); err != nil {
		log.Fatal("cannot run migrations", err)
	}

	router := gin.Default()

	//repos
	companyRepo := repositories.NewCompanyRepository(pool)

	//services
	companySvc := services.NewCompanyService(companyRepo)

	// 3) controllers and routes
	companyCtrl := controllers.NewCompanyController(companySvc)
	companyCtrl.RegisterRoutes(router)

	router.Run(":8080")

}
