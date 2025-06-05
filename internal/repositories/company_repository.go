package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinLL22/stock-tests/internal/models"
	"strconv"
)

type CompanyRepository interface {
	Upsert(ctx context.Context, companyModel models.Company) error
	FindAll(ctx context.Context) ([]models.Company, error)
	Get(ctx context.Context, id string) (models.Company, error)
	Delete(ctx context.Context, id string) error
}

type CompanyRepo struct{ db *pgxpool.Pool }

func NewCompanyRepository(db *pgxpool.Pool) CompanyRepository {
	return &CompanyRepo{db: db}
}

func (cRepo CompanyRepo) Upsert(ctx context.Context, companyModel models.Company) error {

	if companyModel.ID == 0 {
		const insertQuery = `
            INSERT INTO companies (ticker, name)
            VALUES ($1, $2)
            RETURNING company_id
        `
		return cRepo.db.QueryRow(ctx, insertQuery, companyModel.Ticker, companyModel.Name).
			Scan(&companyModel.ID)
	}

	const updateQuery = `
        UPSERT INTO companies (company_id, ticker, name)
        VALUES ($1, $2, $3)
    `
	_, err := cRepo.db.Exec(ctx, updateQuery, companyModel.ID, companyModel.Ticker, companyModel.Name)
	return err
}

func (cRepo CompanyRepo) FindAll(ctx context.Context) ([]models.Company, error) {

	const query = `
		SELECT company_id, ticker, name
		FROM companies
		ORDER BY company_id
	`

	rows, err := cRepo.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companyModelsList []models.Company
	for rows.Next() {
		var company models.Company
		if err := rows.Scan(&company.ID, &company.Ticker, &company.Name); err != nil {
			return nil, err
		}
		companyModelsList = append(companyModelsList, company)
	}
	return companyModelsList, rows.Err()
}

func (cRepo CompanyRepo) Get(ctx context.Context, id string) (models.Company, error) {
	companyID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.Company{}, fmt.Errorf("invalid company ID: %w", err)
	}

	const query = `
		SELECT company_id, ticker, name
		FROM companies
		WHERE company_id = $1
	`
	var company models.Company
	err = cRepo.db.QueryRow(ctx, query, companyID).Scan(&company.ID, &company.Ticker, &company.Name)
	if err != nil {
		return models.Company{}, err
	}
	return company, nil
}

func (cRepo CompanyRepo) Delete(ctx context.Context, id string) error {

	companyID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid company ID: %w", err)
	}

	const query = `
		DELETE FROM companies
		WHERE company_id = $1
	`
	_, err = cRepo.db.Exec(ctx, query, companyID)
	return err
}
