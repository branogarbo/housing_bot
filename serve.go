package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func serveLastHTMLResponse() {
	s := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	s.Get("/", func(c *fiber.Ctx) error {
		if strings.TrimSpace(lastResponse)[:1] == "<" {
			c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		} else {
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		}

		return c.SendString(lastResponse)
	})

	s.Listen(":3000")
}
