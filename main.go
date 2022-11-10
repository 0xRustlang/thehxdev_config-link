package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi")
	})

	app.Get("/surfboard.conf", func(c *fiber.Ctx) error {
		return c.SendFile("/root/config-link/surfboard.conf")
	})

	app.Get("/clash.yaml", func(c *fiber.Ctx) error {
		return c.SendFile("/root/config-link/clash.yaml")
	})

	app.Get("/xray.txt", func(c *fiber.Ctx) error {
		return c.SendFile("/root/config-link/xray.txt")
	})

	app.Listen(":3000")
}
