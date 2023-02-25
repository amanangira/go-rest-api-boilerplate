package repository

import (
	"boilerplate/dtos"
	"context"
	"github.com/jmoiron/sqlx"
)

type ICompanyRepository interface {
	Create(ctx context.Context, user *dtos.CompanyRequestDTO) (string, error)
}

type CompanyRepository struct {
	dbClient *sqlx.DB
}

func NewCompanyRepository(dbClient *sqlx.DB) ICompanyRepository {
	return CompanyRepository{
		dbClient,
	}
}

func (u CompanyRepository) Create(ctx context.Context, user *dtos.CompanyRequestDTO) (string, error) {
	// TODO - handle gracefully
	//query := ""
	//var row pgx.Row

	//user.SetCreateTimeStamp()
	//user.SetUpdateTimestamp()
	//
	//query = `INSERT INTO users
	//		(email, password_hash, first_name, last_name, created_at, updated_at)
	//		values ($1, $2, $3, $4, $5, $6) returning id;`
	//row = u.dbClient.QueryRowxContext(
	//	ctx,
	//	query,
	//	user.Email,
	//	user.PasswordHash,
	//	user.FirstName,
	//	user.LastName,
	//	user.CreatedAt,
	//	user.UpdatedAt)
	//
	//var id string
	//if err := row.Scan(&id); err != nil {
	//	return "", err
	//}

	return "", nil
}
