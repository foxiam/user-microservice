package router

import (
	"user-microservice/internal/api/handler"
	"user-microservice/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app     *fiber.App
	handler *handler.Handler
}

func NewServer(app *fiber.App, handler *handler.Handler) *Server {
	return &Server{app: app, handler: handler}
}

func (s *Server) Router() {

	api := s.app.Group("/api", logger.New())

	auth := api.Group("/auth")
	user := api.Group("/user")
	favorite := user.Group("/favorite-cities")

	//Auth
	auth.Post("/login", s.handler.SignIn)
	auth.Post("/registration", s.handler.SingUp)

	//Favorite
	favorite.Get("/:id", s.handler.GetAllByUserId)
	favorite.Post("", middleware.Protected(), s.handler.AddToFavorite)
	favorite.Delete("", middleware.Protected(), s.handler.DeleteFromFavorite)

	//User
	user.Get("/all", s.handler.GetAllUsers)
	user.Get("/:id", s.handler.GetUser)
	user.Delete("/:id", middleware.Protected(), s.handler.DeleteUser)
}
