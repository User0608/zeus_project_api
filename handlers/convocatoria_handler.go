package handlers

import (
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/services"
	"github.com/User0608/zeus_project_api/utils"
	"github.com/labstack/echo/v4"
)

type ConvocatoriaHandler struct {
	binder  echo.DefaultBinder
	service *services.ConvocatoriaService
}

func NewConvocatoriaHandler(s *services.ConvocatoriaService) *ConvocatoriaHandler {
	return &ConvocatoriaHandler{service: s, binder: echo.DefaultBinder{}}
}

func (h *ConvocatoriaHandler) Create(c echo.Context) error {
	convocatoria := &models.Convocatoria{}
	if err := h.binder.BindBody(c, convocatoria); err != nil {
		return errores.JsonErrorResponse(c, nil)
	}
	if err := h.service.Create(convocatoria); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, convocatoria)
}
func (h *ConvocatoriaHandler) FindAll(c echo.Context) error {
	convocatorias, err := h.service.FindAll()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, convocatorias)
}

func (h *ConvocatoriaHandler) FindByNombre(c echo.Context) error {
	nombre := c.Param("nombre")
	convocatoria, err := h.service.FindByName(nombre)
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, convocatoria)
}

func (h *ConvocatoriaHandler) Update(c echo.Context) error {
	convocatoria := &models.Convocatoria{}
	if err := h.binder.BindBody(c, convocatoria); err != nil {
		return errores.JsonErrorResponse(c, nil)
	}
	if err := h.service.Update(convocatoria); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, convocatoria)
}
func (h *ConvocatoriaHandler) Delete(c echo.Context) error {
	nombre := c.Param("nombre")
	if err := h.service.Delete(nombre); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKMessage(c, "operacion completada")
}
