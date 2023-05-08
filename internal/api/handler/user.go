package handler

import (
	"context"

	"user-microservice/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (h *Handler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.services.User.GetUser(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.services.User.GetAllUsers(context.Background())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Users found", "data": users})
}

func (h *Handler) SignIn(c *fiber.Ctx) error {

	input := new(model.User)

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err.Error()})
	}

	tokenString, err := h.services.User.GenerateToken(context.Background(), input)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": tokenString})
}

func (h *Handler) SingUp(c *fiber.Ctx) error {
	type NewUser struct {
		Id    uint   `json:"id"`
		Email string `json:"email"`
	}

	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	id, err := h.services.User.CreateUser(context.Background(), user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to create user", "data": err.Error()})
	}

	newUser := NewUser{
		Id:    id,
		Email: user.Email,
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	type PasswordInput struct {
		Password string `json:"password"`
	}

	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	err := h.services.User.DeleteUser(context.Background(), id, pi.Password, token)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}
