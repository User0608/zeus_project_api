package repos

import (
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"gorm.io/gorm"
)

type ConvocatoriaRepository struct {
	conn *gorm.DB
}

func NewConvocatoriaRepository(c *gorm.DB) *ConvocatoriaRepository {
	return &ConvocatoriaRepository{conn: c}
}

func (r *ConvocatoriaRepository) FindAll() ([]*models.Convocatoria, error) {
	convocatorias := []*models.Convocatoria{}
	if err := r.conn.Find(&convocatorias).Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	return convocatorias, nil
}
func (r *ConvocatoriaRepository) Update(c *models.Convocatoria) error {
	if err := r.conn.Updates(c).Error; err != nil {
		return errores.NewInternalDBf(err)
	}
	return nil
}
func (r *ConvocatoriaRepository) Create(c *models.Convocatoria) error {
	if err := r.conn.Create(c).Error; err != nil {
		return errores.NewInternalDBf(err)
	}
	return nil
}
func (r *ConvocatoriaRepository) FindByName(name string) (*models.Convocatoria, error) {
	c := &models.Convocatoria{}
	if err := r.conn.Limit(1).Find(c, "nombre=?", name).Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	return c, nil
}

func (r *ConvocatoriaRepository) Delete(name string) error {
	if err := r.conn.Delete(&models.Convocatoria{}, "nombre=?", name).Error; err != nil {
		return errores.NewInternalDBf(err)
	}
	return nil
}
