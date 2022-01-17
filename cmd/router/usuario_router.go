package router

import (
	"github.com/User0608/zeus_project_api/cmd/auth"
	"github.com/User0608/zeus_project_api/cmd/injectors"
	"github.com/labstack/echo/v4"
)

func usuarioUpgrade(e *echo.Echo) {
	h := injectors.GetUsuarioHandler()
	e.POST("/login", h.Login)
	g := e.Group("/usuario")
	g.Use(auth.JWTMiddleware)
	g.POST("", h.CreateUser)
	g.GET("", h.FindAll)
	g.GET("/free", h.FreeUsers)
	g.DELETE("", h.Delete)
}
