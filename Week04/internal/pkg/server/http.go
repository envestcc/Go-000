package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	sigChan         chan os.Signal
}

type Router interface {
	Route(*gin.Engine)
}

type Option func(*http.Server)

func NewReadTimeoutOption(timeout time.Duration) Option {
	return func(srv *http.Server) {
		srv.ReadTimeout = timeout
	}
}

var (
	DefaultOptions = []Option{
		NewReadTimeoutOption(600 * time.Second),
	}
)

func NewHttpServer(listen string, router Router, options ...Option) IServer {
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
		srv:     srv,
		sigChan: make(chan os.Signal),
	}
}

func (s *httpServer) Start() {
	go s.handleSignals()
	s.srv.ListenAndServe()
}

func (s *httpServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), s.gracefulTimeout)
	defer cancel()

	s.srv.Shutdown(ctx)
}

func (s *httpServer) handleSignals() {
	signal.Notify(s.sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-s.sigChan
	switch sig {
	case syscall.SIGINT, syscall.SIGTERM:
		s.Stop()
	}
}
