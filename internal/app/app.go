package app

import (
	"github.com/gofiber/fiber/v3"
)

type App struct {
	srvr *fiber.App
	addr string
}

func Get(addr string) *App {
	res := fiber.New(
		fiber.Config{},
	)

	res.Get("/", func(c fiber.Ctx) error {
		// TODO
		return nil
	})

	return &App{
		srvr: res,
		addr: addr,
	}
}

func (app *App) Run() {
	app.srvr.Listen(app.addr)
}

func (app *App) Shutdown() {
	app.srvr.Shutdown()
}
