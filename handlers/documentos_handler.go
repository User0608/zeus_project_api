package handlers

import (
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/services"
	"github.com/User0608/zeus_project_api/utils"
	"github.com/labstack/echo/v4"
)

type DocumentoHandler struct {
	ser    *services.DocumentoService
	binder echo.DefaultBinder
}

func NewDocumentoHandler(s *services.DocumentoService) *DocumentoHandler {
	return &DocumentoHandler{ser: s, binder: echo.DefaultBinder{}}
}
func (h *DocumentoHandler) CrearMemorando(c echo.Context) error {
	memorando := models.Memorando{}
	if err := h.binder.BindBody(c, &memorando); err != nil {
		return errores.JsonErrorResponse(c, nil)
	}
	if err := h.ser.CreateMemorando(&memorando); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, &memorando)
}
func (h *DocumentoHandler) FindMemorandos(c echo.Context) error {
	memorandos, err := h.ser.FindMemorandos()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, memorandos)
}
