package main

import (
	"flag"

	"github.com/famusovsky/aleksey-space/internal/app"
)

func main() {
	addr := flag.String("addr", ":8080", "addr")
	flag.Parse()
	application := app.Get(*addr)
	application.Run()
}
