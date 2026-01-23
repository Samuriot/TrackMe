package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samuriot/track-me/handlers"
)

// SetupProductRoutes configures all product-related routes
func SetupProductRoutes(app *fiber.App, handler *handlers.UserHandler) {
	// Create a route group for products
	productGroup := app.Group("/api/products")
	productGroup.Get("/", handler.GetAllUsers)
	productGroup.Post("/", handler.CreateUser)
	productGroup.Get("/:id", handler.GetUser)
	productGroup.Put("/:id", handler.UpdateUser)
	productGroup.Delete("/:id", handler.DeleteUser)
}