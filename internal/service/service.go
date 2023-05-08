package service

import (
	"context"
	"user-microservice/internal/model"
	"user-microservice/internal/repository"

	"github.com/golang-jwt/jwt/v4"
)

type User interface {
	CreateUser(ctx context.Context, user *model.User) (id uint, err error)
	DeleteUser(ctx context.Context, id, pass string, t *jwt.Token) error
	GenerateToken(ctx context.Context, user *model.User) (string, error)
	GetUser(ctx context.Context, id string) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
}

type City interface {
	GetAllByUserId(ctx context.Context, userId string) ([]*model.City, error)
	AddToFavorite(ctx context.Context, userId, cityName string, t *jwt.Token) error
	DeleteFromFavorite(ctx context.Context, userId, cityName string, t *jwt.Token) error
}

type Service struct {
	User
	City
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		User: NewUserService(r.User),
		City: NewCityService(r.City),
	}
}
