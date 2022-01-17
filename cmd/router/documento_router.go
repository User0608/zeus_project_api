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
	g.Static("/memorando/pdfs", "files/memorandos")
}
