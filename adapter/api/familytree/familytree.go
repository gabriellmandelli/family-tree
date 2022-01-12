package familytree

import (
	"net/http"

	"github.com/gabriellmandelli/family-tree/business/familytree"
	util "github.com/gabriellmandelli/family-tree/foundation/context"
	"github.com/labstack/echo/v4"
)

//FamilyTreeAPI struct
type FamilyTreeAPI struct {
	rsService familytree.FamilyTreeService
}

//NewFamilyTreeAPI return Person api
func NewFamilyTreeAPI(familytreeService familytree.FamilyTreeService) *FamilyTreeAPI {
	return &FamilyTreeAPI{
		rsService: familytreeService,
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
