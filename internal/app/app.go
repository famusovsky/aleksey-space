package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	srvr *fiber.App
	addr string
}

func Get(addr string) *App {
	res := fiber.New(
		fiber.Config{},
	)

	res.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	return &App{
		srvr: res,
		addr: addr,
	}
}

func (app *App) Run() {
	go func() {
		err := app.srvr.Listen(app.addr)
		log.Fatal(err)
	}()
}

func (app *App) Shutdown() {
	log.Println("Shutting down the server")
	app.srvr.Shutdown()
}
