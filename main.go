package main

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"hello": "world",
		})
	})

	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        3,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusTooManyRequests).JSON(&fiber.Map{
				"error": "Too many requests",
			})
		},
	}))

	app.Get("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.JSON(&fiber.Map{
			"hello": name,
		})
	})

	app.Listen(":3000")
}
