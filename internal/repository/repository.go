package repository

import (
	"context"
	"user-microservice/internal/model"
	"user-microservice/pkg/database"
)

const (
	userTable     = "public.user"
	cityTable     = "public.city"
	favoriteTable = "public.favorite"
)

type User interface {
	FindById(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (id uint, err error)
	DeleteUser(ctx context.Context, id string) error
}

type City interface {
	FindByName(ctx context.Context, name string) (*model.City, error)
	FindAllByUserId(ctx context.Context, userId string) ([]*model.City, error)
	AddToFavorite(ctx context.Context, userId, cityId string) error
	DeleteFromFavorite(ctx context.Context, userId, cityId string) error
}

type Repository struct {
	User
	City
}

func NewRepository(db database.PGXQuerier) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
		City: NewCityPostgres(db),
	}
}
