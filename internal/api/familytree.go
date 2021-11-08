package api

import (
	"net/http"

	"github.com/gabriellmandelli/family-tree/internal/service"
	"github.com/gabriellmandelli/family-tree/internal/util"
	"github.com/labstack/echo/v4"
)

//FamilyTreeAPI struct
type FamilyTreeAPI struct {
	rsService service.FamilyTreeService
}

//NewFamilyTreeAPI return Person api
func NewFamilyTreeAPI() *FamilyTreeAPI {
	return &FamilyTreeAPI{
		rsService: service.NewFamilyTreeService(),
	}
}

const (
	familyTreeBaseUrl = "/familytree"
)

//Register register controllers
func (api *FamilyTreeAPI) Register(server *echo.Echo) {
	v1 := server.Group("v1")
	v1.GET(familyTreeBaseUrl+"/:personID", api.findAll)
}

func (api *FamilyTreeAPI) findAll(echoCtx echo.Context) error {
	ctx, _, errx := util.InitializeContext(echoCtx, nil)

	personID := echoCtx.Param("personID")

	if errx != nil || personID == "" {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	person, errx := api.rsService.GetFamilyTree(ctx, personID)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	return echoCtx.JSON(http.StatusOK, person)
}
