package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi")
	})

	app.Get("/surfboard.conf", func(c *fiber.Ctx) error {
		return c.SendFile("./surfboard.conf")
	})

	app.Get("/clash.yaml", func(c *fiber.Ctx) error {
		return c.SendFile("./clash.yaml")
	})

	app.Listen(":3000")
}
