package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gabriellmandelli/family-tree/adapter/api/familytree"
	"github.com/gabriellmandelli/family-tree/adapter/api/health"
	"github.com/gabriellmandelli/family-tree/adapter/api/person"
	"github.com/gabriellmandelli/family-tree/adapter/api/relationship"
	"github.com/gabriellmandelli/family-tree/adapter/config"
	"github.com/gabriellmandelli/family-tree/adapter/database"
	"github.com/gabriellmandelli/family-tree/adapter/router"
	familytreeBusiness "github.com/gabriellmandelli/family-tree/business/familytree"
	personBusiness "github.com/gabriellmandelli/family-tree/business/person"
	relationshipBusiness "github.com/gabriellmandelli/family-tree/business/relationship"
)

func main() {

	//--Dependencies--
	//Config
	cfg := config.GetConfig()

	//Db
	db, _ := database.NewMongoDbClient(context.TODO(), cfg)

	//Repository
	personRepository := personBusiness.NewPersonRepository(db)
	relationshipRepository := relationshipBusiness.NewRelationShipRepository(db)

	//Service
	personService := personBusiness.NewPersonService(personRepository)
	relationshipService := relationshipBusiness.NewRelationShipService(relationshipRepository)
	familyTreeService := familytreeBusiness.NewFamilyTreeService(personService, relationshipService)

	//Http
	personApi := person.NewPersonAPI(personService)
	relationshipApi := relationship.NewRelationShipAPI(relationshipService)
	familytreeApi := familytree.NewFamilyTreeAPI(familyTreeService)

	//Router
	server := router.NewRouter()

	//Register
	personApi.Register(server)
	relationshipApi.Register(server)
	familytreeApi.Register(server)

	go func() {
		server.Logger.Fatal(server.Start(cfg.AppName))
	}()

	healthCheck := router.NewRouter()

	health.NewHealthCheckAPI().Register(healthCheck)

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
