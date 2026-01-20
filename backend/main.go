package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"github.com/samuriot/track-me/db"
)

func main() {
	db.Init()
	app := fiber.New()
	
	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	log.Fatal(app.Listen(":3000"))
}