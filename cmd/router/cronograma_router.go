package router

import (
	"github.com/User0608/zeus_project_api/cmd/injectors"
	"github.com/labstack/echo/v4"
)

func cronogramaUpgrade(e *echo.Echo) {
	h := injectors.GetCronogramaHandler()
	g := e.Group("/cronograma")
	g.GET("", h.FindCronograms)
}
