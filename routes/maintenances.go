package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rubichandrap/hello-go/controllers"
	"github.com/rubichandrap/hello-go/middleware"
)

func MaintenancesRoutes(router fiber.Router) {
	m := router.Group("/maintenances", middleware.JWTProtected())
	c := controllers.UserController{}

	m.Get("/", c.Get)
	m.Get("/:userID", c.GetByID)
	m.Post("/", c.Store)
	m.Put("/:userID", c.Update)
	m.Put("/change-password/:userID", c.UpdatePassword)
}
