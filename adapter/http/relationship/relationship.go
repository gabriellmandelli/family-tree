package relationship

import (
	"net/http"

	"github.com/gabriellmandelli/family-tree/business/relationship"
	util "github.com/gabriellmandelli/family-tree/foundation/context"
	"github.com/labstack/echo/v4"
)

//RelationShip struct
type RelationShip struct {
	rsService relationship.RelationShipService
}

//NewRelationShipHttp return RelationShip rs
func NewRelationShipHttp(relationShipService relationship.RelationShipService) *RelationShip {
	return &RelationShip{
		rsService: relationShipService,
	}
}

const (
	relationShipBaseUrl = "/relationship"
)

//Register register controllers
func (controller *RelationShip) Register(server *echo.Echo) {
	v1 := server.Group("v1")
	v1.GET(relationShipBaseUrl, controller.findAll)
	v1.POST(relationShipBaseUrl, controller.add)
}

func (rs *RelationShip) findAll(echoCtx echo.Context) error {
	ctx, _, errx := util.InitializeContext(echoCtx, nil)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	person, errx := rs.rsService.FindAll(ctx, "", "")

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	return echoCtx.JSON(http.StatusOK, person)
}

func (rs *RelationShip) add(echoCtx echo.Context) error {
	requestBody := relationship.RelationShip{}

	ctx, _, errx := util.InitializeContext(echoCtx, &requestBody)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	relationShip, errx := rs.rsService.Add(ctx, &requestBody)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	return echoCtx.JSON(http.StatusOK, relationShip)
}
