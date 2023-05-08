package repository

import (
	"context"
	"fmt"
	"user-microservice/internal/model"
	"user-microservice/pkg/database"
)

type CityRepository struct {
	db database.PGXQuerier
}

func NewCityPostgres(db database.PGXQuerier) *CityRepository {
	return &CityRepository{db: db}
}

func (r *CityRepository) FindByName(ctx context.Context, name string) (*model.City, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = $1", cityTable)
	fmt.Println(query)
	var city model.City
	err := r.db.QueryRow(ctx, query, name).Scan(&city.Id, &city.Name)
	return &city, err
}

func (r *CityRepository) FindAllByUserId(ctx context.Context, userId string) ([]*model.City, error) {
	query := fmt.Sprintf("SELECT id, name FROM %s JOIN %s c ON c.id = city_id WHERE user_id = $1",
		favoriteTable, cityTable)

	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []*model.City
	for rows.Next() {
		var city model.City
		err = rows.Scan(&city.Id, &city.Name)
		if err != nil {
			return nil, err
		}
		cities = append(cities, &city)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cities, nil
}

func (r *CityRepository) AddToFavorite(ctx context.Context, userId, cityId string) error {
	query := fmt.Sprintf("INSERT INTO %s(user_id, city_id) VALUES ($1, $2)", favoriteTable)

	_, err := r.db.Exec(ctx, query, userId, cityId)
	return err
}

func (r *CityRepository) DeleteFromFavorite(ctx context.Context, userId, cityId string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 and city_id = $2", favoriteTable)

	_, err := r.db.Exec(ctx, query, userId, cityId)
	return err
}
