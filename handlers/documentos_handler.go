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
func (h *DocumentoHandler) CrearInforme(c echo.Context) error {
	informe := models.Informe{}
	if err := h.binder.BindBody(c, &informe); err != nil {
		return errores.JsonErrorResponse(c, nil)
	}
	if err := h.ser.CreateInforme(&informe); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, &informe)
}
func (h *DocumentoHandler) CrearOficio(c echo.Context) error {
	oficio := models.Oficio{}
	if err := h.binder.BindBody(c, &oficio); err != nil {
		return errores.JsonErrorResponse(c, nil)
	}
	if err := h.ser.CreateOficio(&oficio); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, &oficio)
}
func (h *DocumentoHandler) FindMemorandos(c echo.Context) error {
	memorandos, err := h.ser.FindMemorandos()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, memorandos)
}
func (h *DocumentoHandler) FindInformes(c echo.Context) error {
	informes, err := h.ser.FindInformes()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, informes)
}
func (h *DocumentoHandler) FindOficios(c echo.Context) error {
	oficios, err := h.ser.FindOficios()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, oficios)
}
