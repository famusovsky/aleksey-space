package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func (app *App) signIn(c *fiber.Ctx) error {
	password := c.Query("pswd", "")
	realPassword, _ := os.ReadFile("password.txt")

	if password != string(realPassword) {
		return c.SendString(`password is wrong`)
	}

	app.ch.Set(c, code)

	return c.Redirect("/edit")
}

func (app *App) signOut(c *fiber.Ctx) error {
	app.ch.Remove(c)

	return c.Redirect("/")
}
