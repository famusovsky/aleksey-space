package app

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
)

const code string = "amogus is kinda sus"

type App struct {
	srvr         *fiber.App
	f            http.FileSystem
	addr         string
	text         template.HTML
	ch           cookieHandler
	textFile     string
	passwordFile string
}

func Get(addr, textFile, passwordFile string, f http.FileSystem) *App {
	res := &App{
		srvr: fiber.New(fiber.Config{
			Views: html.NewFileSystem(f, ".html")}),
		addr:         addr,
		ch:           getCookieHandler("user-info", "aleksey-space"),
		f:            f,
		textFile:     textFile,
		passwordFile: passwordFile,
	}

	txt, err := os.ReadFile(textFile)
	if err != nil {
		res.srvr.Shutdown()
		panic(err)
	}
	res.text = template.HTML(txt)

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
