package router

import (
	"github.com/User0608/zeus_project_api/cmd/injectors"
	"github.com/labstack/echo/v4"
)

func personaUpgrade(e *echo.Echo) {
	h := injectors.GetPersonaHandler()
	g := e.Group("/persona")
	g.GET("/info", h.InfoBaseForEntity)
	g.POST("/primerjefe", h.RegistrarPrimerJefe)
	g.GET("/primerjefe", h.PrimerJefeInfo)
	g.POST("/segundojefe", h.RegistrarSegundoJefe)
	g.GET("/segundojefe", h.SegundoJefeInfo)
	g.GET("/resumenps", h.ResumenPrimerSegunJefe)

	g.POST("/jefeinstruccion", h.RegistrarJefeInstruccion)
	g.GET("/jefeinstruccion", h.FindJefeInstruccion)
	/////
	g.POST("/instructor", h.RegistrarInstructor)
	g.PUT("/instructor", h.UpdateInstructor)
	g.GET("/instructor/:dni", h.FindInstructorByDNI)
	g.GET("/instructor", h.FindAllInstructores)
	g.GET("/instructor/all", h.ResumenAllInstructores)
}
