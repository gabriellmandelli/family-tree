package person

import (
	"net/http"

	"github.com/gabriellmandelli/family-tree/business/person"
	util "github.com/gabriellmandelli/family-tree/foundation/context"
	"github.com/labstack/echo/v4"
)

//PersonAPI struct
type PersonAPI struct {
	personService person.PersonService
}

//NewPersonAPI return Person api
func NewPersonAPI() *PersonAPI {
	return &PersonAPI{
		personService: person.NewPersonService(),
	}
}

const (
	personBaseUrl = "/person"
)

//Register register controllers
func (controller *PersonAPI) Register(server *echo.Echo) {
	v1 := server.Group("v1")
	v1.GET(personBaseUrl, controller.findAll)
	v1.POST(personBaseUrl, controller.addPerson)
	v1.PUT(personBaseUrl+"/:personID", controller.addPerson)
}

func (api *PersonAPI) findAll(echoCtx echo.Context) error {
	ctx, _, errx := util.InitializeContext(echoCtx, nil)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	person, errx := api.personService.FindAllPerson(ctx, "")

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	return echoCtx.JSON(http.StatusOK, person)
}

func (api *PersonAPI) addPerson(echoCtx echo.Context) error {
	requestBody := person.Person{}

	ctx, _, errx := util.InitializeContext(echoCtx, &requestBody)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	person, errx := api.personService.AddPerson(ctx, &requestBody)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	return echoCtx.JSON(http.StatusOK, person)
}
