package router

import (
	"github.com/User0608/zeus_project_api/cmd/injectors"
	"github.com/labstack/echo/v4"
)

func moduloesbasUpgrade(e *echo.Echo) {
	h := injectors.GetModuloesbasHandler()
	g := e.Group("/curso")
	g.GET("/esbas", h.Find)
	g.GET("/ciclo", h.FindCiclo)
}
