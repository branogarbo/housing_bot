package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func serveLastResponse(c *fiber.Ctx) error {
	if strings.TrimSpace(lastResponse)[:1] == "<" {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	} else {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	}

	return c.SendString(lastResponse)
}

func (b Bot) initFiber() {
	s := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	s.Get("/", serveLastResponse)

	s.Get("/live", func(c *fiber.Ctx) error {
		err := b.checkAddress(false)
		if err != nil {
			return err
		}

		return serveLastResponse(c)
	})

	s.Listen(":3000")
}
