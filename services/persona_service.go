package services

import (
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/repos"
	"github.com/ksaucedo002/kcheck"
)

type PersonaService struct {
	sign       string
	repository *repos.PersonaRepository
}

func NewPersonaService(r *repos.PersonaRepository) *PersonaService {
	return &PersonaService{repository: r, sign: "service.PersonaService"}
}

func (s *PersonaService) FindInfoBaseForEntity() (*models.InformacionBaseEntity, error) {
	return s.repository.FindInfoBase()
}
func (s *PersonaService) RegistrarPrimerJefe(j *models.PrimerJefeEntity) (*models.PrimerJefeEntity, error) {
	if err := kcheck.Valid(j.Entity); err != nil {
		return nil, errores.NewBadRequestf(nil, err.Error())
	}
	return s.repository.RegistrarPrimerJefe(j)
}
func (s *PersonaService) LoadPrimerJefeInfo() (*models.PrimerJefeEntity, error) {
	return s.repository.LoadPrimerJefe()
}

/////////////====================//////////////Segundo jefe

func (s *PersonaService) RegistrarSegundoJefe(j *models.SegundoJefeEntity) (*models.SegundoJefeEntity, error) {
	if err := kcheck.Valid(j.Entity); err != nil {
		return nil, errores.NewBadRequestf(nil, err.Error())
	}
	return s.repository.RegistrarSegundoJefe(j)
}
func (s *PersonaService) LoadSegundoJefe() (*models.SegundoJefeEntity, error) {
	return s.repository.LoadSegundoJefe()
}

////
func (s *PersonaService) ResumenPrimerJefeSegundoJefe() (*models.PrimerJefeEntity, *models.SegundoJefeEntity, error) {
	return s.repository.ResumenPrimerSegunJefe()
}

////////////===================///////////////////Jefe Instruccion
func (s *PersonaService) RegistrarJefeInstruccion(j *models.JefeInstruccionEntity) (*models.JefeInstruccionEntity, error) {
	if err := kcheck.Valid(j.Entity); err != nil {
		return nil, errores.NewBadRequestf(nil, err.Error())
	}
	return s.repository.RegistrarJefeInstruccion(j)
}
func (s *PersonaService) LoadJefeInstruccion() (*models.JefeInstruccionEntity, error) {
	return s.repository.LoadJefeInstruccion()
}

///////////////===========//////////////// Instructores

func (s *PersonaService) RegistrarInstructor(i *models.InstructorEntity) error {
	if err := kcheck.Valid(i.Entity); err != nil {
		return errores.NewBadRequestf(nil, err.Error())
	}
	return s.repository.RegistrarInstructor(i)
}
func (s *PersonaService) UpdateInstructor(i *models.InstructorEntity) error {
	if err := kcheck.Valid(i.Entity); err != nil {
		return errores.NewBadRequestf(nil, err.Error())
	}
	return s.repository.UpdateInstructor(i)
}
func (s *PersonaService) FindInstructorByDNI(dni string) (*models.InstructorEntity, error) {
	if len(dni) == 0 {
		return nil, errores.NewBadRequestf(nil, "param dni es necesario")
	}
	return s.repository.FindInstructorByDNI(dni)
}

func (s *PersonaService) FindAllInstructores() ([]*models.InstructorEntity, error) {
	return s.repository.FindAllInstructors()
}
