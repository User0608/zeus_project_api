package handlers

import (
	"time"

	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/services"
	"github.com/User0608/zeus_project_api/utils"
	"github.com/labstack/echo/v4"
)

type RepositoyHandler struct {
	service *services.ReporteService
}

func NewRepositoyHandler(s *services.ReporteService) *RepositoyHandler {
	return &RepositoyHandler{service: s}
}
func (h *RepositoyHandler) Actividades(c echo.Context) error {
	inicio, err := time.Parse("2006-01-02", c.QueryParam("fecha_inicio"))
	if err != nil {
		return errores.ErrorResponse(c, errores.NewBadRequestf(nil, "fecha de inicio incorrecta"))
	}
	fin, err := time.Parse("2006-01-02", c.QueryParam("fecha_fin"))
	if err != nil {
		return errores.ErrorResponse(c, errores.NewBadRequestf(nil, "fecha de fin incorrecta"))
	}
	data, err := h.service.ActividadesIntervaloMeses(inicio, fin)
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, data)
}
