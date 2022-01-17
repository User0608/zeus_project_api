package handlers

import (
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/services"
	"github.com/User0608/zeus_project_api/utils"
	"github.com/labstack/echo/v4"
)

type CodigoHandler struct {
	binder echo.DefaultBinder
	serv   *services.CodigoService
}

func NewCodigoHandler(s *services.CodigoService) *CodigoHandler {
	return &CodigoHandler{serv: s, binder: echo.DefaultBinder{}}
}

func (h *CodigoHandler) AllCodigos(c echo.Context) error {
	codigos, err := h.serv.FindAllCodigos()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, codigos)
}
