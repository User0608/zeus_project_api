package services

import (
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/repos"
	"github.com/ksaucedo002/kcheck"
)

type CronogramaService struct {
	res *repos.CronogramaRepository
}

func NewCronogramaService(r *repos.CronogramaRepository) *CronogramaService {
	return &CronogramaService{res: r}
}

func (s *CronogramaService) CrearCronograma(c *models.Cronograma) error {
	if err := kcheck.Valid(c); err != nil {
		return errores.NewBadRequestf(nil, err.Error())
	}
	return s.res.CrearCronograma(c)
}

func (s *CronogramaService) DeleteCronograma(cronoid uint) error {
	if cronoid == 0 {
		return errores.NewBadRequestf(nil, "conograma id, invalido!")
	}
	return s.res.DeleteCronograma(cronoid)
}

func (s *CronogramaService) CrearProgramacionForCronograma(crono_id uint, p *models.Programacion) error {
	if crono_id == 0 {
		return errores.NewBadRequestf(nil, "conograma id, invalido!")
	}
	if err := kcheck.Valid(p); err != nil {
		return errores.NewBadRequestf(nil, err.Error())
	}
	return s.res.AddProgramacion(crono_id, p)
}

func (s *CronogramaService) DeleteProgramacin(programacionID uint) error {
	if programacionID == 0 {
		return errores.NewBadRequestf(nil, "programacion id, invalido!")
	}
	return s.res.DeleteProgramacion(programacionID)

}

func (s *CronogramaService) CreateActividadForProgramacion(programacionid uint, a *models.Actividad) error {
	if programacionid == 0 {
		return errores.NewBadRequestf(nil, "programacion id, invalido!")
	}
	if err := kcheck.Valid(a); err != nil {
		return errores.NewBadRequestf(nil, err.Error())
	}
	return s.res.AddActividadToProgramacion(programacionid, a)
}

func (s *CronogramaService) DeleteActivida(actividadid uint) error {
	if actividadid == 0 {
		return errores.NewBadRequestf(nil, "actividad id, invalido!")
	}
	return s.res.DeleteActividad(actividadid)
}

func (s *CronogramaService) FindCronogramas() ([]*models.Cronograma, error) {
	return s.res.FindCronograms()
}
func (s *CronogramaService) FindCronogramasOnly() ([]*models.Cronograma, error) {
	return s.res.FindCronogramsOnly()
}
func (s *CronogramaService) FindProgramaciones(cronoid uint) ([]*models.Programacion, error) {
	return s.res.FindProgramacion(cronoid)
}
func (s *CronogramaService) FindActividaes(programacionid uint) ([]*models.Actividad, error) {
	return s.res.FindActividades(programacionid)
}
