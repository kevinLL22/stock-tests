package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinLL22/stock-tests/internal/config"
	"github.com/kevinLL22/stock-tests/internal/controllers"
	"github.com/kevinLL22/stock-tests/internal/repositories"
	"github.com/kevinLL22/stock-tests/internal/services"
)

func main() {

	cfg := config.Load()

	router := gin.Default()

	// inits

	// pool
	pool, _ := pgxpool.New(context.Background(), cfg.DatabaseURL)

	//repos
	companyRepo := repositories.NewCompanyRepository(pool)

	//services
	companySvc := services.NewCompanyService(companyRepo)

	// 3) controllers and routes
	companyCtrl := controllers.NewCompanyController(companySvc)
	companyCtrl.RegisterRoutes(router)

	router.Run(":8080")

}
