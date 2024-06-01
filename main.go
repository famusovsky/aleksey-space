package main

import (
	"embed"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/famusovsky/aleksey-space/internal/app"
	"golang.org/x/sync/errgroup"
)

//go:embed ui
var f embed.FS

func main() {
	addr := flag.String("addr", ":8080", "addr")
	txt := flag.String("txt", "./text", "txt file")
	pswd := flag.String("pswd", "./password", "pswd file")
	flag.Parse()
	application := app.Get(*addr, *txt, *pswd, http.FS(f))

	sigQuit := make(chan os.Signal, 2)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)
	eg := new(errgroup.Group)

	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			return fmt.Errorf("captured signal: %v", s)
		}
	})

	application.Run()

	if err := eg.Wait(); err != nil {
		application.Shutdown()
	}
}
