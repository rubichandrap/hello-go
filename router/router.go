package router

import (
	"github.com/rubichandrap/hello-go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	api := app.Group("/api", logger.New())

	routes.AuthRoutes(api)
	routes.MaintenancesRoutes(api)
}
