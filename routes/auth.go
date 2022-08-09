package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rubichandrap/hello-go/controllers"
)

func AuthRoutes(router fiber.Router) {
	m := router.Group("/auth")
	c := controllers.AuthController{}

	m.Post("/register", c.Register)
	m.Post("/login", c.Login)

}
