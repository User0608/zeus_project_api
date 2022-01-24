package repos

import (
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"gorm.io/gorm"
)

type CronogramaRepository struct {
	conn *gorm.DB
}

func NewCronogramaRepository(c *gorm.DB) *CronogramaRepository {
	return &CronogramaRepository{conn: c}
}

func (r *CronogramaRepository) CrearCronograma(c *models.Cronograma) error {
	if err := r.conn.Create(c).Error; err != nil {
		return errores.NewInternalDBf(err)
	}
	return nil
}

func (r *CronogramaRepository) FindCronograms() ([]*models.Cronograma, error) {
	crons := []*models.Cronograma{}
	if err := r.conn.Preload("Programas.Actividades").Find(&crons).Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	return crons, nil
}

func (r *CronogramaRepository) CrearProgramacion(p *models.Programacion) error {
	return nil
}
