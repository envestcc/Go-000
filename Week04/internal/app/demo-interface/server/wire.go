// +build wireinject

package server

import (
	"Week04/internal/app/demo-interface/service"
	"Week04/internal/pkg/dao"
	"Week04/internal/pkg/server"
	"database/sql"

	"github.com/google/wire"
)

func NewServerWire(listen string, db *sql.DB) server.IServer {
	panic(wire.Build(server.NewHttpServer, NewRouter, service.NewService, dao.NewAccount, NewOptions))
}
