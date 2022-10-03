package person

import (
	"net/http"

	"github.com/gabriellmandelli/family-tree/internal/business/person"
	util "github.com/gabriellmandelli/family-tree/internal/foundation/context"
	"github.com/labstack/echo/v4"
)

//Person struct
type Person struct {
	personService person.PersonService
}

//NewPersonHttp return Person p
func NewPersonHttp(personService person.PersonService) *Person {
	return &Person{
		personService: personService,
	}
}

const (
	personBaseUrl = "/person"
)

//Register register controllers
func (controller *Person) Register(server *echo.Echo) {
	v1 := server.Group("v1")
	v1.GET(personBaseUrl, controller.findAll)
	v1.POST(personBaseUrl, controller.addPerson)
	v1.PUT(personBaseUrl+"/:personID", controller.addPerson)
}

func (p *Person) findAll(echoCtx echo.Context) error {
	ctx, _, errx := util.InitializeContext(echoCtx, nil)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	person, errx := p.personService.FindAllPerson(ctx, "")

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	return echoCtx.JSON(http.StatusOK, person)
}

func (p *Person) addPerson(echoCtx echo.Context) error {
	requestBody := person.Person{}

	ctx, _, errx := util.InitializeContext(echoCtx, &requestBody)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	person, errx := p.personService.AddPerson(ctx, &requestBody)

	if errx != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	return echoCtx.JSON(http.StatusOK, person)
}
