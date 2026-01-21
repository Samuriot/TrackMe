package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samuriot/track-me/handlers"
)

// SetupProductRoutes configures all product-related routes
func SetupProductRoutes(app *fiber.App) {
	// Create a route group for products
	productGroup := app.Group("/api/products")
	productGroup.Get("/", handlers.GetAllUsers)
	productGroup.Post("/", handlers.CreateUser)
	productGroup.Get("/:id", handlers.GetUser)
	productGroup.Put("/:id", handlers.UpdateUser)
	productGroup.Delete("/:id", handlers.DeleteUser)
}