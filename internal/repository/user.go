package repository

import (
	"context"
	"fmt"

	"user-microservice/internal/model"
	"user-microservice/pkg/database"
)

type UserRepository struct {
	db database.PGXQuerier
}

func NewUserPostgres(db database.PGXQuerier) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindById(ctx context.Context, id string) (*model.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", userTable)

	var user model.User
	err := r.db.QueryRow(ctx, query, id).Scan(&user.Id, &user.Email, &user.Password)
	return &user, err
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", userTable)

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", userTable)

	var user model.User
	err := r.db.QueryRow(ctx, query, email).Scan(&user.Id, &user.Email, &user.Password)
	return &user, err
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) (id uint, err error) {
	query := fmt.Sprintf("INSERT INTO %s (email, password) VALUES ($1, $2) RETURNING id", userTable)
	err = r.db.QueryRow(ctx, query, user.Email, user.Password).Scan(&id)
	return id, err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", userTable)
	_, err := r.db.Exec(ctx, query, id)
	return err
}
