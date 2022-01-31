package handlers

import (
	"strconv"

	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/services"
	"github.com/User0608/zeus_project_api/utils"
	"github.com/labstack/echo/v4"
)

type CronogramaHandler struct {
	binder echo.DefaultBinder
	ser    *services.CronogramaService
}

func NewCronogramaHandler(s *services.CronogramaService) *CronogramaHandler {
	return &CronogramaHandler{ser: s, binder: echo.DefaultBinder{}}
}

func (h *CronogramaHandler) FindCronograms(c echo.Context) error {
	cronos, err := h.ser.FindCronogramas()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, cronos)
}

func (h *CronogramaHandler) FindProgramacion(c echo.Context) error {
	cronid, err := strconv.Atoi(c.Param("cronograma_id"))
	if err != nil || cronid < 0 {
		return errores.ErrorResponse(c, errores.NewBadRequestf(nil, "cronograma id invalido"))
	}
	programaciones, err := h.ser.FindProgramaciones(uint(cronid))
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, &programaciones)
}

func (h *CronogramaHandler) FindActividades(c echo.Context) error {
	cronid, err := strconv.Atoi(c.Param("programacion_id"))
	if err != nil || cronid < 0 {
		return errores.ErrorResponse(c, errores.NewBadRequestf(nil, "programacion id invalido"))
	}
	actividades, err := h.ser.FindActividaes(uint(cronid))
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, &actividades)
}

func (h *CronogramaHandler) FindCronogramsOnly(c echo.Context) error {
	cronos, err := h.ser.FindCronogramasOnly()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, cronos)
}

func (h *CronogramaHandler) CreateCrono(c echo.Context) error {
	crono := models.Cronograma{}
	crono.CreatedBy = "73491346"
	if err := h.binder.BindBody(c, &crono); err != nil {
		return errores.JsonErrorResponse(c, nil)
	}
	if err := h.ser.CrearCronograma(&crono); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, crono)
}

func (h *CronogramaHandler) DeleteCrono(c echo.Context) error {
	cronid, err := strconv.Atoi(c.Param("cronograma_id"))
	if err != nil || cronid < 0 {
		return errores.ErrorResponse(c, errores.NewBadRequestf(nil, "cronograma_id invalido"))
	}
	if err := h.ser.DeleteCronograma(uint(cronid)); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKSuccess(c)
}

func (h *CronogramaHandler) CrearProgramacionForCronograma(c echo.Context) error {
	cronid, err := strconv.Atoi(c.Param("cronograma_id"))
	if err != nil || cronid < 0 {
		return errores.ErrorResponse(c, errores.NewBadRequestf(nil, "cronograma_id invalido"))
	}
	p := models.Programacion{}
	if err := h.binder.BindBody(c, &p); err != nil {
		return errores.JsonErrorResponse(c, nil)
	}
	if err := h.ser.CrearProgramacionForCronograma(uint(cronid), &p); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, p)
}

func (h *CronogramaHandler) DeleteProgramacion(c echo.Context) error {
	programacionid, err := strconv.Atoi(c.Param("programacion_id"))
	if err != nil || programacionid < 0 {
		return errores.ErrorResponse(c, errores.NewBadRequestf(nil, "programacion_id invalido"))
	}
	if err := h.ser.DeleteProgramacin(uint(programacionid)); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKSuccess(c)
}

func (h *CronogramaHandler) CreateActividadForProgramacion(c echo.Context) error {
	programacionid, err := strconv.Atoi(c.Param("programacion_id"))
	if err != nil || programacionid < 0 {
		return errores.ErrorResponse(c, errores.NewBadRequestf(nil, "cronograma_id invalido"))
	}
	a := models.Actividad{}
	if err := h.binder.BindBody(c, &a); err != nil {
		return errores.JsonErrorResponse(c, nil)
	}
	if err := h.ser.CreateActividadForProgramacion(uint(programacionid), &a); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, &a)
}
func (h *CronogramaHandler) DeleteActividad(c echo.Context) error {
	actividadid, err := strconv.Atoi(c.Param("actividad_id"))
	if err != nil || actividadid < 0 {
		return errores.ErrorResponse(c, errores.NewBadRequestf(nil, "cronograma_id invalido"))
	}
	if err := h.ser.DeleteActivida(uint(actividadid)); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKSuccess(c)
}
