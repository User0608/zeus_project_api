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

func (r *CronogramaRepository) deleteWithID(id uint, model interface{}) error {
	res := r.conn.Where("id=?", id).Delete(model)
	if res.Error != nil {
		return errores.NewInternalDBf(res.Error)
	}
	if res.RowsAffected == 0 {
		return errores.NewNotFoundf(nil, errores.ErrRecordNotFaund)
	}
	return nil
}

func (r *CronogramaRepository) CrearCronograma(c *models.Cronograma) error {
	if err := r.conn.Create(c).Error; err != nil {
		return errores.NewInternalDBf(err)
	}
	return nil
}
func (r *CronogramaRepository) DeleteCronograma(cronoid uint) error {
	return r.deleteWithID(cronoid, &models.Cronograma{})
}
func (r *CronogramaRepository) AddProgramacion(crono_id uint, p *models.Programacion) error {
	p.CronogramaID = crono_id
	if err := r.conn.Create(p).Error; err != nil {
		return errores.NewInternalDBf(err)
	}
	return nil
}
func (r *CronogramaRepository) FindProgramacion(crono_id uint) ([]*models.Programacion, error) {
	programaciones := []*models.Programacion{}
	if err := r.conn.Find(&programaciones, "cronograma_id=?", crono_id).Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	return programaciones, nil
}
func (r *CronogramaRepository) DeleteProgramacion(programacionID uint) error {
	return r.deleteWithID(programacionID, &models.Programacion{})
}
func (r *CronogramaRepository) AddActividadToProgramacion(programacionID uint, a *models.Actividad) error {
	a.ProgramacionID = programacionID
	if err := r.conn.Create(a).Error; err != nil {
		return errores.NewInternalDBf(err)
	}
	return nil
}
func (r *CronogramaRepository) FindActividades(programacionID uint) ([]*models.Actividad, error) {
	programaciones := []*models.Actividad{}
	if err := r.conn.Find(&programaciones, "programacion_id=?", programacionID).Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	return programaciones, nil
}
func (r *CronogramaRepository) DeleteActividad(actividadID uint) error {
	return r.deleteWithID(actividadID, &models.Actividad{})
}
func (r *CronogramaRepository) FindCronograms() ([]*models.Cronograma, error) {
	crons := []*models.Cronograma{}
	if err := r.conn.Preload("Programas.Actividades").Find(&crons).Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	return crons, nil
}
func (r *CronogramaRepository) FindCronogramsOnly() ([]*models.Cronograma, error) {
	crons := []*models.Cronograma{}
	if err := r.conn.Find(&crons).Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	return crons, nil
}
