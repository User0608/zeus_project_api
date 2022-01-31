package router

import (
	"github.com/User0608/zeus_project_api/cmd/injectors"
	"github.com/labstack/echo/v4"
)

func cronogramaUpgrade(e *echo.Echo) {
	h := injectors.GetCronogramaHandler()
	g := e.Group("/cronograma")
	g.GET("", h.FindCronograms)
	g.GET("/only", h.FindCronogramsOnly)
	g.POST("", h.CreateCrono)
	g.DELETE("/:cronograma_id", h.DeleteCrono)
	g.POST("/programacion/:cronograma_id", h.CrearProgramacionForCronograma)
	g.GET("/programacion/:cronograma_id", h.FindProgramacion)
	e.DELETE("/programacion/:programacion_id", h.DeleteProgramacion)
	g.POST("/actividad/:programacion_id", h.CreateActividadForProgramacion)
	g.GET("/actividad/:programacion_id", h.FindActividades)
	e.DELETE("/actividad/:actividad_id", h.DeleteActividad)
}
