package router

import (
	"github.com/User0608/zeus_project_api/cmd/injectors"
	"github.com/labstack/echo/v4"
)

func codigoUpgrade(e *echo.Echo) {
	h := injectors.GetCodigoHandler()
	g := e.Group("/estaticos")
	g.GET("/codigos", h.AllCodigos)
}
