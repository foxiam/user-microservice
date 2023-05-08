package service

import (
	"context"
	"errors"
	"strconv"
	"user-microservice/internal/model"
	"user-microservice/internal/repository"

	"github.com/golang-jwt/jwt/v4"
)

type CityService struct {
	repo repository.City
}

func NewCityService(r repository.City) *CityService {
	return &CityService{repo: r}
}

func (s *CityService) GetAllByUserId(ctx context.Context, userId string) ([]*model.City, error) {
	cities, err := s.repo.FindAllByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return cities, nil
}

func (s *CityService) AddToFavorite(ctx context.Context, userId, cityName string, t *jwt.Token) error {
	if !validToken(t, userId) {
		return errors.New("invalid token id")
	}

	city, err := s.repo.FindByName(ctx, cityName)
	if err != nil {
		return err
	}

	cityId := strconv.FormatUint(uint64(city.Id), 10)
	return s.repo.AddToFavorite(ctx, userId, cityId)
}

func (s *CityService) DeleteFromFavorite(ctx context.Context, userId, cityName string, t *jwt.Token) error {
	if !validToken(t, userId) {
		return errors.New("invalid token id")
	}

	city, err := s.repo.FindByName(ctx, cityName)
	if err != nil {
		return err
	}

	cityId := strconv.FormatUint(uint64(city.Id), 10)
	return s.repo.DeleteFromFavorite(ctx, userId, cityId)
}
