package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinLL22/stock-tests/internal/models"
	"github.com/kevinLL22/stock-tests/internal/services"
	"net/http"
	"strconv"
)

type CompanyController struct {
	service *services.CompanySvc
}

func NewCompanyController(svc *services.CompanySvc) *CompanyController {
	return &CompanyController{service: svc}
}

func (ctrl *CompanyController) RegisterRoutes(router *gin.Engine) {
	grp := router.Group("/companies")
	{
		grp.POST("", ctrl.CreateOrUpdate)
		grp.GET("", ctrl.ListAll)
		grp.GET("/:id", ctrl.GetByID)
		grp.DELETE("/:id", ctrl.DeleteByID)
	}
}

func (ctrl *CompanyController) CreateOrUpdate(context *gin.Context) {

	// Parsing JSON to DTO
	var dto struct {
		Ticker string `json:"ticker" binding:"required"`
		Name   string `json:"name"   binding:"required"`
		ID     string `json:"id"` // optional, can be empty
	}
	if err := context.ShouldBindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// build the model from DTO
	company := models.Company{
		Ticker: dto.Ticker,
		Name:   dto.Name,
	}

	if dto.ID != "" {
		id, err := strconv.ParseInt(dto.ID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "ID invalid, must be a number"})
			return
		}
		company.ID = id
	}

	// call service
	if err := ctrl.service.CreateOrUpdate(context.Request.Context(), company); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Response
	if company.ID == 0 {

		context.Status(http.StatusCreated)
	} else {
		context.Status(http.StatusOK)
	}
}

func (ctrl *CompanyController) ListAll(c *gin.Context) {
	companies, err := ctrl.service.ListAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}

func (ctrl *CompanyController) GetByID(c *gin.Context) {

	//validate parameter
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter id is required"})
		return
	}

	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id, must be a number"})
		return
	}

	company, err := ctrl.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, company)
}

// DeleteByID handles the deletion of a company by its ID.
func (ctrl *CompanyController) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter id is required"})
		return
	}
	// Validar formato numérico
	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id, must be a number"})
		return
	}

	if err := ctrl.service.DeleteByID(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
