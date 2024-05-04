package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/famusovsky/aleksey-space/internal/app"
	"golang.org/x/sync/errgroup"
)

func main() {
	addr := flag.String("addr", ":8080", "addr")
	flag.Parse()
	application := app.Get(*addr)

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
