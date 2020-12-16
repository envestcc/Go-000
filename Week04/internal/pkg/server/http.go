package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	Start()
	Stop()
}

type httpServer struct {
	srv             *http.Server
	gracefulTimeout time.Duration
}

type Router interface {
	Route(*gin.Engine)
}

type option func(*http.Server)

func NewReadTimeoutOption(timeout time.Duration) option {
	return func(srv *http.Server) {
		srv.ReadTimeout = timeout
	}
}

var (
	DefaultOptions = []option{
		NewReadTimeoutOption(600 * time.Second),
	}
)

func NewHttpServer(listen string, router Router, options ...option) IServer {
	engine := gin.New()
	router.Route(engine)

	srv := &http.Server{
		Addr:    listen,
		Handler: engine,
	}

	for i := range DefaultOptions {
		DefaultOptions[i](srv)
	}
	for i := range options {
		options[i](srv)
	}

	return &httpServer{
		srv: srv,
	}
}

func (s *httpServer) Start() {
	s.srv.ListenAndServe()
}

func (s *httpServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), s.gracefulTimeout)
	defer cancel()

	s.srv.Shutdown(ctx)
}
