package app

import "github.com/gofiber/fiber/v2"

func (app *App) setRoutes() {
	app.srvr.Get("/", func(c *fiber.Ctx) error {
		return c.Render("ui/views/me", fiber.Map{})
	})

	app.srvr.Get("/hi", func(c *fiber.Ctx) error {
		return c.Render("ui/views/hi", fiber.Map{})
	})
}
