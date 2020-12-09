package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {

	sig := make(chan os.Signal)
	shutdown := make(chan int)
	exit := make(chan int, 2)
	defer close(exit)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() error {
		s := <-sig
		close(shutdown)
		return fmt.Errorf("signal=%+v", s)
	}()

	errs := &errgroup.Group{}
	errs.Go(func() error {
		if err := AppServer(shutdown, exit); err != nil {
			exit <- 1
			return err
		}
		return nil
	})
	errs.Go(func() error {
		if err := DebugServer(shutdown, exit); err != nil {
			exit <- 1
			return err
		}
		return nil
	})

	if err := errs.Wait(); err != nil {
		log.Printf("Main Exit %+v", err)
	}

}

func AppServer(shutdown, exitOuter chan int) error {
	server := &http.Server{
		Addr: ":8000",
	}

	exit := make(chan int)
	defer close(exit)
	go func() {
		select {
		case <-shutdown:
			server.Shutdown(context.Background())
		case <-exit:
		case <-exitOuter:
			server.Shutdown(context.Background())
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		log.Printf("AppServer Error: %+v", err)
		return fmt.Errorf("AppServer: %+v", err)
	}
	return nil
}

func DebugServer(shutdown, exitOuter chan int) error {
	server := &http.Server{
		Addr: ":9000",
	}

	exit := make(chan int)
	defer close(exit)
	go func() {
		select {
		case <-shutdown:
			server.Shutdown(context.Background())
		case <-exit:
		case <-exitOuter:
			server.Shutdown(context.Background())
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		log.Printf("DebugServer Error: %+v", err)
		return fmt.Errorf("DebugServer: %+v", err)
	}
	return nil
}
