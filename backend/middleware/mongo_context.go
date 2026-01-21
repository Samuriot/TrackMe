package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
)

func MongoContextMiddleware(timeout time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		c.Locals("mongoCtx", ctx)
		return c.Next()
	}
}
