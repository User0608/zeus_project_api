package services

import (
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/repos"
	"github.com/ksaucedo002/kcheck"
)

type ConvocatoriaService struct {
	repo *repos.ConvocatoriaRepository
}

func NewConvocatoriaService(r *repos.ConvocatoriaRepository) *ConvocatoriaService {
	return &ConvocatoriaService{repo: r}
}

func (r *ConvocatoriaService) FindAll() ([]*models.Convocatoria, error) {
	return r.repo.FindAll()
}
func (r *ConvocatoriaService) Update(c *models.Convocatoria) error {
	if err := kcheck.Valid(c); err != nil {
		return errores.NewBadRequestf(nil, err.Error())
	}
	return r.repo.Update(c)
}
func (r *ConvocatoriaService) Create(c *models.Convocatoria) error {
	if err := kcheck.Valid(c); err != nil {
		return errores.NewBadRequestf(nil, err.Error())
	}
	return r.repo.Create(c)
}
func (r *ConvocatoriaService) FindByName(name string) (*models.Convocatoria, error) {
	if name == "" {
		return nil, errores.NewBadRequestf(nil, "nombre de la convocatoria invalido")
	}
	return r.FindByName(name)
}

func (r *ConvocatoriaService) Delete(name string) error {
	if name == "" {
		return errores.NewBadRequestf(nil, "nombre de la convocatoria invalido")
	}
	return r.Delete(name)
}
