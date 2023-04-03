package main

import "github.com/gofiber/fiber/v2"

func serveLastHTMLResponse() {
	s := fiber.New(fiber.Config{
		GETOnly:               true,
		DisableStartupMessage: true,
	})

	s.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

		return c.SendString(lastHTMLResponse)
	})

	s.Listen(":3000")
}
