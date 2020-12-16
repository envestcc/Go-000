package server

import (
	"Week04/internal/app/demo-interface/service"
	"Week04/internal/pkg/server"
	"time"

	"github.com/gin-gonic/gin"
)

func NewServer(listen string, srv service.IService) server.IServer {
	return server.NewHttpServer(listen, &Router{
		srv: srv,
	}, server.NewReadTimeoutOption(60*time.Second))
}

type Router struct {
	srv service.IService
}

func (r *Router) Route(engine *gin.Engine) {
	engine.GET("accounts", r.srv.ListAccount)
}
