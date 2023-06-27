// Package main runs the server and connects the database
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {

	db := getDb()

	// Create the web server
	app := fiber.New()

	app.Use(limiter.New())

	// Initialize middleware
	authMiddleware := getAuthMiddleware()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/api/desk/distance", authMiddleware, func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		d := new(DeskDistance)

		if err := c.BodyParser(d); err != nil {
			return err
		}

		db.Create(&DeskDistance{Distance: d.Distance})

		return c.SendString("Posted distance")
	})

	app.Listen(":3000")
}
