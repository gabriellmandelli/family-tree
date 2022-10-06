package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	familyTreeHttp "github.com/gabriellmandelli/family-tree/internal/adapter/api/http/familytree"
	healthHttp "github.com/gabriellmandelli/family-tree/internal/adapter/api/http/health"
	personHttp "github.com/gabriellmandelli/family-tree/internal/adapter/api/http/person"
	relationShipHttp "github.com/gabriellmandelli/family-tree/internal/adapter/api/http/relationship"
	"github.com/gabriellmandelli/family-tree/internal/adapter/config"
	"github.com/gabriellmandelli/family-tree/internal/adapter/database"
	"github.com/gabriellmandelli/family-tree/internal/business/familytree"
	"github.com/gabriellmandelli/family-tree/internal/business/person"
	"github.com/gabriellmandelli/family-tree/internal/business/relationship"
	"github.com/gabriellmandelli/family-tree/internal/foundation/http/router"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
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

	//Router
	serviceHttp := router.NewRouter()
	metaHttp := router.NewRouter()

	registerServices(db, serviceHttp, metaHttp)

	go func() {
		serviceHttp.Logger.Fatal(serviceHttp.Start(cfg.AppPort))
	}()

	go func() {
		metaHttp.Logger.Fatal(metaHttp.Start(cfg.HealthPort))
	}()

	osSignalChan := make(chan os.Signal, 2)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGTERM)
	switch <-osSignalChan {
	case os.Interrupt:
	case syscall.SIGTERM:
	}
}

func registerServices(db *mongo.Database, serviceHttp *echo.Echo, metaHttp *echo.Echo) {

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

	healtCheckhHttp := healthHttp.NewHealthCheckHttp()

	//Register
	personHttp.Register(serviceHttp)
	relationshipHttp.Register(serviceHttp)
	familytreeHttp.Register(serviceHttp)

	healtCheckhHttp.Register(metaHttp)
}
