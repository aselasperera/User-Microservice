package main

import (
	"user-microservice/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the User Microservice")
	})

	// User routes
	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUser)
	app.Post("/users", handlers.CreateUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)

	app.Listen(":3000")
}
