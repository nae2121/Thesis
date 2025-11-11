package main


import (
"log"
"paperagent/pkg/app"
)


func main() {
srv, err := app.New()
if err != nil { log.Fatal(err) }
if err := srv.Run(); err != nil { log.Fatal(err) }
}