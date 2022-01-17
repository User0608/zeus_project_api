package repos

import (
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"gorm.io/gorm"
)

type ModuloEsbasRepository struct {
	conn *gorm.DB
}

func NewModuloEsbasRepository(c *gorm.DB) *ModuloEsbasRepository {
	return &ModuloEsbasRepository{conn: c}
}

func (r *ModuloEsbasRepository) FindAll() ([]*models.ModuloEsbas, error) {
	modulos := []*models.ModuloEsbas{}
	if err := r.conn.Find(&modulos).Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	return modulos, nil
}
func (r *ModuloEsbasRepository) FindCiclos() ([]*models.Ciclo, error) {
	ciclos := []*models.Ciclo{}
	if err := r.conn.Find(&ciclos).Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	return ciclos, nil
}
