package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinLL22/stock-tests/internal/services"
	"net/http"
)

// Acepta el servicio, no el pool.
func RegisterRoutes(r *gin.Engine, svc services.StockService) {
	// Health-check
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// → Ejemplos de endpoints (stubs):
	//
	// r.POST("/stocks/refresh", func(c *gin.Context) {
	//	   if err := svc.RefreshFromAPI(c.Request.Context()); err != nil {
	//	       c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	       return
	//	   }
	//	   c.Status(http.StatusNoContent)
	// })
	//
	// r.GET("/stocks", func(c *gin.Context) {
	//	   items, _ := svc.FindAll(c.Request.Context())
	//	   c.JSON(http.StatusOK, items)
	// })
}
