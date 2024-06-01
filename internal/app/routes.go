package app

import (
	"html/template"
	"os"

	"github.com/gofiber/fiber/v2"
)

func (app *App) checkItsMe(c *fiber.Ctx) error {
	val, err := app.ch.Read(c)
	if err != nil || val != code {
		return c.Redirect("/")
	}

	return c.Next()
}

func (app *App) setRoutes() {
	app.srvr.Get("/", func(c *fiber.Ctx) error {
		return c.Render("ui/views/me", fiber.Map{
			"text": app.text,
		})
	})
	app.srvr.Get("/hi", func(c *fiber.Ctx) error {
		return c.Render("ui/views/hi", fiber.Map{})
	})
	app.srvr.Get("/in", func(c *fiber.Ctx) error {
		return app.signIn(c)
	})
	app.srvr.Get("/out", func(c *fiber.Ctx) error {
		return app.signOut(c)
	})

	edit := app.srvr.Group("/edit", app.checkItsMe)
	edit.Get("/", func(c *fiber.Ctx) error {
		return c.Render("ui/views/edit", fiber.Map{
			"text": app.text,
		})
	})
	edit.Post("/", func(c *fiber.Ctx) error {
		app.text = template.HTML(c.FormValue("text"))
		os.WriteFile(app.textFile, []byte(app.text), 0644)
		return c.Redirect("/edit")
	})
}
