// Package main runs the server and connects the database
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	db := getDb()

	engine := html.New("./views", ".html")

	// Create the web server
	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/", "./public")

	// app.Use(limiter.New())

	// Initialize middleware
	authMiddleware := getAuthMiddleware()

	app.Get("/", func(c *fiber.Ctx) error {
		var distances []DeskDistance
		var latestDistance DeskDistance

		db.Find(&distances)
		db.Order("created_at desc").First(&latestDistance)

		return c.Render("index", fiber.Map{"Distances": distances, "LatestDistance": latestDistance})
	})

	app.Get("/api/desk/distance", func(c *fiber.Ctx) error {
		var distances []DeskDistance
		db.Find(&distances)
		return c.Status(200).JSON(distances)
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
