package handlers

import (
	"fmt"

	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/handlers/response"
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/services"
	"github.com/User0608/zeus_project_api/utils"
	"github.com/labstack/echo/v4"
)

type PersonaHandler struct {
	binder echo.DefaultBinder
	serv   *services.PersonaService
}

func NewPersonaHandler(s *services.PersonaService) *PersonaHandler {
	return &PersonaHandler{serv: s, binder: echo.DefaultBinder{}}
}

func (h *PersonaHandler) InfoBaseForEntity(c echo.Context) error {
	info, err := h.serv.FindInfoBaseForEntity()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, info)
}
func (h *PersonaHandler) RegistrarPrimerJefe(c echo.Context) error {
	jefe := models.PrimerJefeEntity{}
	if err := h.binder.BindBody(c, &jefe); err != nil {
		return errores.JsonErrorResponse(c, fmt.Errorf("%s:%w", "handlers.RegistrarPrimerJefe", err))
	}
	res, err := h.serv.RegistrarPrimerJefe(&jefe)
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, res)
}
func (h *PersonaHandler) PrimerJefeInfo(c echo.Context) error {
	jefe, err := h.serv.LoadPrimerJefeInfo()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, jefe)
}

/////////=======///Segundo jefe
func (h *PersonaHandler) RegistrarSegundoJefe(c echo.Context) error {
	jefe := models.SegundoJefeEntity{}
	if err := h.binder.BindBody(c, &jefe); err != nil {
		return errores.JsonErrorResponse(c, fmt.Errorf("%s:%w", "handlers.RegistrarSegundoJefe", err))
	}
	res, err := h.serv.RegistrarSegundoJefe(&jefe)
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, res)
}
func (h *PersonaHandler) SegundoJefeInfo(c echo.Context) error {
	jefe, err := h.serv.LoadSegundoJefe()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, jefe)
}

//////////
func (h *PersonaHandler) ResumenPrimerSegunJefe(c echo.Context) error {
	primerJefe, segunJefe, err := h.serv.ResumenPrimerJefeSegundoJefe()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, &response.ResumenPrimerSegunJefe{
		PrimerJefe:  primerJefe,
		SegundoJefe: segunJefe,
	})
}

////////////////////////////////////// jefe de instruccion
func (h *PersonaHandler) RegistrarJefeInstruccion(c echo.Context) error {
	j := &models.JefeInstruccionEntity{}
	if err := h.binder.BindBody(c, j); err != nil {
		return errores.JsonErrorResponse(c, err)
	}
	jefe, err := h.serv.RegistrarJefeInstruccion(j)
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, jefe)
}
func (h *PersonaHandler) FindJefeInstruccion(c echo.Context) error {
	jefe, err := h.serv.LoadJefeInstruccion()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, jefe)
}

//////////////////// Instructor

func (h *PersonaHandler) RegistrarInstructor(c echo.Context) error {
	instructor := &models.InstructorEntity{}
	if err := h.binder.BindBody(c, instructor); err != nil {
		return errores.JsonErrorResponse(c, err)
	}
	if err := h.serv.RegistrarInstructor(instructor); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, instructor)
}
func (h *PersonaHandler) UpdateInstructor(c echo.Context) error {
	instructor := &models.InstructorEntity{}
	if err := h.binder.BindBody(c, instructor); err != nil {
		return errores.JsonErrorResponse(c, err)
	}
	if err := h.serv.UpdateInstructor(instructor); err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, instructor)
}

func (h *PersonaHandler) FindInstructorByDNI(c echo.Context) error {
	dni := c.Param("dni")
	instructor, err := h.serv.FindInstructorByDNI(dni)
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, instructor)
}

func (h *PersonaHandler) FindAllInstructores(c echo.Context) error {
	instructores, err := h.serv.FindAllInstructores()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, instructores)
}
func (h *PersonaHandler) ResumenAllInstructores(c echo.Context) error {
	jefe, err := h.serv.LoadJefeInstruccion()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	instructores, err := h.serv.FindAllInstructores()
	if err != nil {
		return errores.ErrorResponse(c, err)
	}
	return utils.OKResponse(c, &response.ResumenInstructor{
		JefeInstruccion: jefe,
		Instructores:    instructores,
	})
}
