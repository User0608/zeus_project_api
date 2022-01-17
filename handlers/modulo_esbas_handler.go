package handlers

import (
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/services"
	"github.com/User0608/zeus_project_api/utils"
	"github.com/labstack/echo/v4"
)

type ModuloEsbasHandler struct {
	serv *services.ModuloEsbasService
}

func NewModuloEsbasHanlder(s *services.ModuloEsbasService) *ModuloEsbasHandler {
	return &ModuloEsbasHandler{serv: s}
}

func (h *ModuloEsbasHandler) Find(c echo.Context) error {
	modulos, err := h.serv.GetAll()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, modulos)
}
func (h *ModuloEsbasHandler) FindCiclo(c echo.Context) error {
	ciclos, err := h.serv.FindCiclos()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, ciclos)
}
