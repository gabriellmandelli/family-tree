package relationship

import (
	"net/http"

	"github.com/gabriellmandelli/family-tree/business/relationship"
	util "github.com/gabriellmandelli/family-tree/foundation/context"
	"github.com/labstack/echo/v4"
)

//RelationShipAPI struct
type RelationShipAPI struct {
	rsService relationship.RelationShipService
}

//NewRelationShipAPI return Person api
func NewRelationShipAPI(relationShipService relationship.RelationShipService) *RelationShipAPI {
	return &RelationShipAPI{
		rsService: relationShipService,
	}
}

const (
	relationShipBaseUrl = "/relationship"
)

//Register register controllers
func (controller *RelationShipAPI) Register(server *echo.Echo) {
	v1 := server.Group("v1")
	v1.GET(relationShipBaseUrl, controller.findAll)
	v1.POST(relationShipBaseUrl, controller.add)
}

func (api *RelationShipAPI) findAll(echoCtx echo.Context) error {
	ctx, _, errx := util.InitializeContext(echoCtx, nil)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	person, errx := api.rsService.FindAll(ctx, "", "")

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	return echoCtx.JSON(http.StatusOK, person)
}

func (api *RelationShipAPI) add(echoCtx echo.Context) error {
	requestBody := relationship.RelationShip{}

	ctx, _, errx := util.InitializeContext(echoCtx, &requestBody)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	relationShip, errx := api.rsService.Add(ctx, &requestBody)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	return echoCtx.JSON(http.StatusOK, relationShip)
}
