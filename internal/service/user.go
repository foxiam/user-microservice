package service

import (
	"context"
	"errors"
	"net/mail"
	"strconv"
	"time"
	"user-microservice/internal/config"
	"user-microservice/internal/model"
	"user-microservice/internal/repository"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId uint `json:"user_id"`
}

type UserService struct {
	repo repository.User
}

func NewUserService(r repository.User) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) (uint, error) {
	if !validEmail(user.Email) {
		return 0, errors.New("invalid email format")
	}
	hash, err := generatePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hash
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id, pass string, t *jwt.Token) error {
	if !validToken(t, id) {
		return errors.New("invalid token id")
	}

	if !s.validUser(ctx, id, pass) {
		return errors.New("not valid user")
	}

	if err := s.repo.DeleteUser(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *UserService) GenerateToken(ctx context.Context, u *model.User) (string, error) {
	user, err := s.repo.FindByEmail(ctx, u.Email)
	if err != nil {
		return "", err
	}

	if err := checkPasswordHash(u.Password, user.Password); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	})

	return token.SignedString([]byte(config.EnvConfig.SigningKeyJwt))
}

func (s *UserService) GetUser(ctx context.Context, id string) (*model.User, error) {
	user, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func (s *UserService) validUser(ctx context.Context, id string, p string) bool {
	user, err := s.repo.FindById(ctx, id)
	if err != nil {
		return false
	}
	if err = checkPasswordHash(p, user.Password); err != nil {
		return false
	}
	return true
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
