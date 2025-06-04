package services

import (
	"context"
	"fmt"
	"github.com/kevinLL22/stock-tests/internal/models"
	"github.com/kevinLL22/stock-tests/internal/repositories"
)

type CompanySvc struct {
	repo repositories.CompanyRepository
}

func NewCompanyService(repo repositories.CompanyRepository) *CompanySvc {
	return &CompanySvc{repo: repo}
}

func (companyService *CompanySvc) CreateOrUpdate(ctx context.Context, companyModel models.Company) error {
	//validations
	if companyModel.Ticker == "" {
		return fmt.Errorf("ticket cannot be empty")
	}
	if companyModel.Name == "" {
		return fmt.Errorf("need to provide a name for the company")
	}

	// store
	return companyService.repo.Upsert(ctx, companyModel)
}

func (companyService *CompanySvc) ListAll(ctx context.Context) ([]models.Company, error) {
	return companyService.repo.FindAll(ctx)
}

func (companyService *CompanySvc) GetByID(ctx context.Context, id string) (models.Company, error) {

	return companyService.repo.Get(ctx, id)
}

func (companyService *CompanySvc) DeleteByID(ctx context.Context, id string) error {

	_, err := companyService.repo.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("there is not a company wit id %s: %w", id, err)
	}

	return companyService.repo.Delete(ctx, id)
}
