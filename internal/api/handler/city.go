package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (h *Handler) GetAllByUserId(c *fiber.Ctx) error {
	userId := c.Params("id")

	favorites, err := h.services.City.GetAllByUserId(context.Background(), userId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Cities found", "data": favorites})
}

func (h *Handler) AddToFavorite(c *fiber.Ctx) error {
	userId, cityName, err := getInputBody(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed body parse", "data": err.Error()})
	}
	token := c.Locals("user").(*jwt.Token)

	err = h.services.City.AddToFavorite(context.Background(), userId, cityName, token)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed added city to favorite", "data": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "City added to favorite", "data": nil})
}

func (h *Handler) DeleteFromFavorite(c *fiber.Ctx) error {
	userId, cityName, err := getInputBody(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed body parse", "data": err.Error()})
	}
	token := c.Locals("user").(*jwt.Token)

	err = h.services.City.DeleteFromFavorite(context.Background(), userId, cityName, token)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed deleted city from favorite", "data": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "City deleted from favorite", "data": nil})
}

func getInputBody(c *fiber.Ctx) (string, string, error) {
	type Input struct {
		UserId   string `json:"user_id"`
		CityName string `json:"city_name"`
	}

	in := new(Input)
	if err := c.BodyParser(in); err != nil {
		return "", "", err
	}

	return in.UserId, in.CityName, nil
}
