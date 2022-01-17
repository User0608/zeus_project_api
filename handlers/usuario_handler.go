package handlers

import (
	"fmt"

	"github.com/User0608/zeus_project_api/cmd/auth"
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/services"
	"github.com/User0608/zeus_project_api/utils"
	"github.com/labstack/echo/v4"
)

type UsuarioHandler struct {
	binder echo.DefaultBinder
	serv   *services.UsuarioService
}

func NewUsuarioHandler(s *services.UsuarioService) *UsuarioHandler {
	return &UsuarioHandler{serv: s, binder: echo.DefaultBinder{}}
}
func (h *UsuarioHandler) Login(c echo.Context) error {
	request := models.LogginRequest{}
	if err := h.binder.BindBody(c, &request); err != nil {
		return errores.JsonErrorResponse(c, fmt.Errorf("%s:%w", "handlers.Logging", err))
	}
	usuario, err := h.serv.Login(request)
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	if token, err := auth.GenerageToken(*usuario); err != nil {
		return errores.ErrorResponse(c, err)
	} else {
		return utils.OKToken(c, token, usuario)
	}
}
func (h *UsuarioHandler) CreateUser(c echo.Context) error {
	request := models.PostUsuario{}
	if err := h.binder.BindBody(c, &request); err != nil {
		return errores.JsonErrorResponse(c, fmt.Errorf("%s:%w", "handlers.CreateUser", err))
	}
	usuario, err := h.serv.CreateUser(request)
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, &usuario)
}

func (h *UsuarioHandler) FindAll(c echo.Context) error {
	usuarios, err := h.serv.FindAll()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, usuarios)
}

func (h *UsuarioHandler) FreeUsers(c echo.Context) error {
	usuarios, err := h.serv.FreeUsers()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, usuarios)
}

func (h *UsuarioHandler) Delete(c echo.Context) error {
	username := c.QueryParam("username")
	if err := h.serv.Delete(username); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKMessage(c, SUCCESS_OPERATION)
}
