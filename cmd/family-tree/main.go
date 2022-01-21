package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gabriellmandelli/family-tree/adapter/config"
	"github.com/gabriellmandelli/family-tree/adapter/database"
	familyTreeHttp "github.com/gabriellmandelli/family-tree/adapter/http/familytree"
	"github.com/gabriellmandelli/family-tree/adapter/http/health"
	personHttp "github.com/gabriellmandelli/family-tree/adapter/http/person"
	relationShipHttp "github.com/gabriellmandelli/family-tree/adapter/http/relationship"
	"github.com/gabriellmandelli/family-tree/adapter/router"
	"github.com/gabriellmandelli/family-tree/business/familytree"
	"github.com/gabriellmandelli/family-tree/business/person"
	"github.com/gabriellmandelli/family-tree/business/relationship"
)

func main() {

	//--Dependencies--
	//Config
	cfg := config.GetConfig()

	//Db
	db, _ := database.NewMongoDbClient(context.TODO(), cfg)

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

	//Register
	personHttp.Register(server)
	relationshipHttp.Register(server)
	familytreeHttp.Register(server)

	go func() {
		server.Logger.Fatal(server.Start(cfg.AppName))
	}()

	healthCheck := router.NewRouter()

	health.NewHealthCheckHttp().Register(healthCheck)

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
