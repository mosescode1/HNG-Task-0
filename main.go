package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Use(cors.New())

	now := time.Now()

	isoformat := now.Format("2006-01-02T15:04:05Z")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"email":            "iammoses19@gmail.com",
			"current_datetime": isoformat,
			"github_url":       "https://github.com/mosescode1/HNG-Task-0.git",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
