package router

import (
	"github.com/gofiber/fiber/v2"
)

// USER handles all the user routes
var USER fiber.Router

// SetupRoutes setups all the Routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	USER = api.Group("/user")
	SetupUserRoutes()
}
