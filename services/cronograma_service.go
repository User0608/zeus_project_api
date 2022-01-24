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

func (s *CronogramaService) FindCronogramas() ([]*models.Cronograma, error) {
	return s.res.FindCronograms()
}
