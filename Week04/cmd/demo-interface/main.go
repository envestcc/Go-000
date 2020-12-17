package main

import (
	"Week04/internal/app/demo-interface/server"
	"Week04/internal/app/demo-interface/service"
	"Week04/internal/pkg/dao"
)

func main() {
	app := server.NewServer("", service.NewService(dao.NewAccount(nil)))
	app.Start()
}

func loadConfig() {

}
