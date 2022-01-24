package router

import "github.com/labstack/echo/v4"

func Upgrade(e *echo.Echo) {
	usuarioUpgrade(e)
	personaUpgrade(e)
	codigoUpgrade(e)
	convocatoriaUpgrade(e)
	moduloesbasUpgrade(e)
	documentoUpgrade(e)
	cronogramaUpgrade(e)
}
