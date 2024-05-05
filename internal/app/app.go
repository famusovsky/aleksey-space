package app

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
)

type App struct {
	srvr *fiber.App
	f    http.FileSystem
	addr string
}

func Get(addr string, f http.FileSystem) *App {
	res := &App{
		srvr: fiber.New(fiber.Config{
			Views: html.NewFileSystem(f, ".html")}),
		addr: addr,
		f:    f,
	}

	res.srvr.Use("/static", filesystem.New(filesystem.Config{
		Root:       f,
		PathPrefix: "ui/static",
		Browse:     true,
	}))

	res.setRoutes()

	return res
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
