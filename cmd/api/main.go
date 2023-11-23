package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/larnTechGeeks/pc-test/web/routes"
)

func main() {

	log.Print("Starting starter template gateway....")

	appRouter := routes.BuildRouter()

	server := &http.Server{
		Addr:    ":" + "9000",
		Handler: appRouter,
	}

	done := make(chan struct{})

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("Server shut down error: %v", err)
		}

		log.Print("Process terminated...shutting down")

		close(done)
	}()

	defer shutdown()

	log.Printf("server listening on port :%v", 9000)

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("Server shut down")
		} else {
			log.Fatal("Server shut down unexpectedly!", err)
		}
	}

	// Give server 30 seconds to shutdown gracefully
	timeout := 30 * time.Second

	// Allow user to forcibly terminate during graceful shutdown
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

	code := 0
	select {
	case <-sigint:
		code = 1
		log.Print("Process forcibly terminated")
	case <-time.After(timeout):
		code = 1
		log.Print("Shutdown timeout. Forcibly shutting down...")
	case <-done:
		log.Print("Shutdown completed...")

	}

	os.Exit(code)
}

func shutdown() {

	_, cancelTimeout := context.WithTimeout(context.Background(), time.Second*30)

	defer cancelTimeout()

	time.Sleep(time.Millisecond * 200)

	log.Print("bye bye")
}
