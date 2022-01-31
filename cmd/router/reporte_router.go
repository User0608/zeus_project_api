package router

import (
	"github.com/User0608/zeus_project_api/cmd/injectors"
	"github.com/labstack/echo/v4"
)

func reportesUpgrade(e *echo.Echo) {
	h := injectors.GetReporteHandler()
	g := e.Group("/reporte")
	g.GET("/actividad", h.Actividades)
}
