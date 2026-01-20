package db

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

const mongoCtxKey string = "mongoCtx"

func GetMongoContext(c *fiber.Ctx) context.Context {
	ctx := c.Locals(mongoCtxKey)
	if ctx == nil {
		// fallback safety net
		return context.Background()
	}
	return ctx.(context.Context)
}
