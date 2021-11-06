package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//HealthCheckAPI struct
type HealthCheckAPI struct {
}

//NewHealthCheckAPI return HealthCheckAPI
func NewHealthCheckAPI() *HealthCheckAPI {
	return &HealthCheckAPI{}
}

const (
	healthCheckBaseURL = "/health"
)

//Register register controllers
func (controller *HealthCheckAPI) Register(server *echo.Echo) {
	v1 := server.Group("v1")
	v1.GET(healthCheckBaseURL, controller.getHealthCheck)
}

func (controller *HealthCheckAPI) getHealthCheck(context echo.Context) error {
	statusOk := make(map[string]string, 0)
	statusOk["status"] = "ok"
	return context.JSON(http.StatusOK, statusOk)
}
