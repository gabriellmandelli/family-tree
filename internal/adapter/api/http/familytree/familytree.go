package http

import (
	"net/http"

	"github.com/gabriellmandelli/family-tree/internal/business/familytree"
	util "github.com/gabriellmandelli/family-tree/internal/foundation/context"
	"github.com/labstack/echo/v4"
)

// FamilyTree struct
type FamilyTree struct {
	rsService familytree.FamilyTreeService
}

// NewFamilyTree return Person ft
func NewFamilyTreeHttp(familytreeService familytree.FamilyTreeService) *FamilyTree {
	return &FamilyTree{
		rsService: familytreeService,
	}
}

const (
	familyTreeBaseUrl = "/familytree"
)

// Register register controllers
func (ft *FamilyTree) Register(server *echo.Echo) {
	v1 := server.Group("v1")
	v1.GET(familyTreeBaseUrl+"/:personID", ft.findAll)
}

func (ft *FamilyTree) findAll(echoCtx echo.Context) error {
	ctx, _, errx := util.InitializeContext(echoCtx, nil)

	personID := echoCtx.Param("personID")

	if errx != nil || personID == "" {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	person, errx := ft.rsService.GetFamilyTree(ctx, personID)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	return echoCtx.JSON(http.StatusOK, person)
}
