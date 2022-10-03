package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gabriellmandelli/family-tree/internal/adapter/config"
	"github.com/gabriellmandelli/family-tree/internal/adapter/database"
	familyTreeHttp "github.com/gabriellmandelli/family-tree/internal/adapter/http/familytree"
	"github.com/gabriellmandelli/family-tree/internal/adapter/http/health"
	personHttp "github.com/gabriellmandelli/family-tree/internal/adapter/http/person"
	relationShipHttp "github.com/gabriellmandelli/family-tree/internal/adapter/http/relationship"
	"github.com/gabriellmandelli/family-tree/internal/adapter/router"
	"github.com/gabriellmandelli/family-tree/internal/business/familytree"
	"github.com/gabriellmandelli/family-tree/internal/business/person"
	"github.com/gabriellmandelli/family-tree/internal/business/relationship"
)

func main() {

	//--Dependencies--
	//Config
	cfg := config.GetConfig()

	//Db
	db, err := database.NewMongoDbClient(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}

	//Repository
	personRepository := person.NewPersonRepository(db)
	relationshipRepository := relationship.NewRelationShipRepository(db)

	//Service
	personService := person.NewPersonService(personRepository)
	relationshipService := relationship.NewRelationShipService(relationshipRepository)
	familyTreeService := familytree.NewFamilyTreeService(personService, relationshipService)

	//Http
	personHttp := personHttp.NewPersonHttp(personService)
	relationshipHttp := relationShipHttp.NewRelationShipHttp(relationshipService)
	familytreeHttp := familyTreeHttp.NewFamilyTreeHttp(familyTreeService)

	//Router
	server := router.NewRouter()

	healthCheck := router.NewRouter()

	//Register
	personHttp.Register(server)
	relationshipHttp.Register(server)
	familytreeHttp.Register(server)

	health.NewHealthCheckHttp().Register(healthCheck)

	go func() {
		server.Logger.Fatal(server.Start(cfg.AppPort))
	}()

	go func() {
		healthCheck.Logger.Fatal(healthCheck.Start(cfg.HealthPort))
	}()

	osSignalChan := make(chan os.Signal, 2)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGTERM)
	switch <-osSignalChan {
	case os.Interrupt:
	case syscall.SIGTERM:
	}
}
