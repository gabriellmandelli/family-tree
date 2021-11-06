package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gabriellmandelli/family-tree/internal/api"
	"github.com/gabriellmandelli/family-tree/internal/config"
	"github.com/gabriellmandelli/family-tree/internal/router"
)

func main() {

	server := router.NewRouter()

	api.NewPersonAPI().Register(server)

	go func() {
		server.Logger.Fatal(server.Start(config.GetConfig().AppPort))
	}()

	health := router.NewRouter()

	api.NewHealthCheckAPI().Register(health)

	go func() {
		health.Logger.Fatal(health.Start(config.GetConfig().HealthPort))
	}()

	osSignalChan := make(chan os.Signal, 2)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGTERM)
	switch <-osSignalChan {
	case os.Interrupt:
	case syscall.SIGTERM:
	}
}
