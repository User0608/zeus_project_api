package router

import (
	"github.com/User0608/zeus_project_api/cmd/injectors"
	"github.com/labstack/echo/v4"
)

func convocatoriaUpgrade(e *echo.Echo) {
	h := injectors.GetConvocatoria()
	g := e.Group("/convocatoria")
	g.GET("", h.FindAll)
	g.GET("/:nombre", h.FindByNombre)
	g.POST("", h.Create)
	g.PUT("", h.Update)
	g.DELETE("", h.Delete)
}
