package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//HealthCheck struct
type HealthCheck struct {
}

//NewHealthCheckHttp return HealthCheck
func NewHealthCheckHttp() *HealthCheck {
	return &HealthCheck{}
}

const (
	healthCheckBaseURL = "/health"
)

//Register register hc
func (hc *HealthCheck) Register(server *echo.Echo) {
	v1 := server.Group("v1")
	v1.GET(healthCheckBaseURL, hc.getHealthCheck)
}

func (hc *HealthCheck) getHealthCheck(context echo.Context) error {
	statusOk := make(map[string]string, 0)
	statusOk["status"] = "ok"
	return context.JSON(http.StatusOK, statusOk)
}
