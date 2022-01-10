package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gabriellmandelli/family-tree/adapter/api/familytree"
	"github.com/gabriellmandelli/family-tree/adapter/api/health"
	"github.com/gabriellmandelli/family-tree/adapter/api/person"
	"github.com/gabriellmandelli/family-tree/adapter/api/relationship"
	"github.com/gabriellmandelli/family-tree/adapter/config"
	"github.com/gabriellmandelli/family-tree/adapter/router"
)

func main() {

	server := router.NewRouter()

	person.NewPersonAPI().Register(server)
	relationship.NewRelationShipAPI().Register(server)
	familytree.NewFamilyTreeAPI().Register(server)

	go func() {
		server.Logger.Fatal(server.Start(config.GetConfig().AppPort))
	}()

	healthCheck := router.NewRouter()

	health.NewHealthCheckAPI().Register(healthCheck)

	go func() {
		healthCheck.Logger.Fatal(healthCheck.Start(config.GetConfig().HealthPort))
	}()

	osSignalChan := make(chan os.Signal, 2)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGTERM)
	switch <-osSignalChan {
	case os.Interrupt:
	case syscall.SIGTERM:
	}
}
