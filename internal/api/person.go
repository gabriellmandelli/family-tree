package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//PersonAPI struct
type PersonAPI struct {
}

//NewPersonAPI return Person api
func NewPersonAPI() *PersonAPI {
	return &PersonAPI{}
}

const (
	personBaseUrl = "/person"
)

//Register register controllers
func (controller *PersonAPI) Register(server *echo.Echo) {
	v1 := server.Group("v1")
	v1.GET(personBaseUrl, controller.statusOk)
}

func (controller *PersonAPI) statusOk(context echo.Context) error {
	return context.JSON(http.StatusOK, nil)
}
