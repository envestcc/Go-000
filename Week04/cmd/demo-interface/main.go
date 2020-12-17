package main

import (
	"Week04/internal/app/demo-interface/server"
)

func main() {
	app := server.NewServerWire("", nil)
	app.Start()
}

func loadConfig() {

}
