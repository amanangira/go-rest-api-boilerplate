package repository

import (
	"boilerplate/store/entity"
	"context"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	Create(ctx context.Context, user entity.User) (string, error)
	Get(ctx context.Context, id string) (entity.User, error)
}

type UserRepository struct {
	dbClient *sqlx.DB
}

func NewUserRepository(dbClient *sqlx.DB) IUserRepository {
	return UserRepository{
		dbClient,
	}
}

func (u UserRepository) Create(ctx context.Context, user entity.User) (string, error) {
	var id string
	err := u.dbClient.GetContext(
		ctx,
		&id,
		`INSERT INTO users (name,email,password) values ($1, $2, $3) returning id;`,
		user.Name,
		user.Email,
		user.Password)

	return id, err
}

func (u UserRepository) Get(ctx context.Context, id string) (entity.User, error) {
	var user entity.User

	err := u.dbClient.GetContext(
		ctx,
		&user,
		`SELECT id, name, email, created_at, updated_at, deleted_at FROM users where id = $1`, id)

	if err != nil {
		return user, err
	}

	return user, nil
}
