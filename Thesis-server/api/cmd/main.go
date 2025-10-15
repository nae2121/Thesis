package main

import (
	"log"
	"paperagent/app"
)
func main() {
	srv, err := app.New()
	if err != nil { log.Fatal(err) }
	if err := srv.Run(); err != nil { log.Fatal(err) }
}