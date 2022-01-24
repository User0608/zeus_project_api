package handlers

import (
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/services"
	"github.com/User0608/zeus_project_api/utils"
	"github.com/labstack/echo/v4"
)

type CronogramaHandler struct {
	ser *services.CronogramaService
}

func NewCronogramaHandler(s *services.CronogramaService) *CronogramaHandler {
	return &CronogramaHandler{ser: s}
}

func (h *CronogramaHandler) FindCronograms(c echo.Context) error {
	cronos, err := h.ser.FindCronogramas()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, cronos)
}
