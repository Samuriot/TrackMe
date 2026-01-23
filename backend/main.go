package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samuriot/track-me/db"
	"github.com/samuriot/track-me/handlers"
	"github.com/samuriot/track-me/middleware"
	"github.com/samuriot/track-me/repository"
	"github.com/samuriot/track-me/routes"
	"github.com/samuriot/track-me/services"
)

func main() {
	mongodb, err := db.Init()
	if err != nil {
		log.Fatal("err: no mongoDB connection")
	}

	// Defer close to ensure MongoDB disconnects on app exit
	defer db.Close()

	app := fiber.New()
	app.Use(middleware.MongoContextMiddleware(5 * time.Second))

	UserRepository := repository.NewMongoUserRepository(mongodb)
	userService := services.NewUserService(UserRepository)
	userHandler := handlers.NewUserHandler(userService)

	routes.SetupProductRoutes(app, userHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("Shutting down server...")
		if err := app.Shutdown(); err != nil {
			log.Printf("Error shutting down server: %v", err)
		}
	}()

	log.Fatal(app.Listen(":3000"))
}
