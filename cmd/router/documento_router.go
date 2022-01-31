package router

import (
	"github.com/User0608/zeus_project_api/cmd/injectors"
	"github.com/labstack/echo/v4"
)

func documentoUpgrade(e *echo.Echo) {
	h := injectors.GetDocumentoHandler()
	g := e.Group("/documento")
	g.POST("/memorando", h.CrearMemorando)
	g.GET("/memorando", h.FindMemorandos)

	g.POST("/informe", h.CrearInforme)
	g.GET("/informe", h.FindInformes)

	g.POST("/oficio", h.CrearOficio)
	g.GET("/oficio", h.FindOficios)

	g.Static("/memorando/pdfs", "files/memorandos")
	g.Static("/oficio/pdfs", "files/oficios")
	g.Static("/informe/pdfs", "files/informes")
}
