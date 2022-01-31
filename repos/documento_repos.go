package repos

import (
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"gorm.io/gorm"
)

type DocumentoRepository struct {
	conn *gorm.DB
}

func NewDocumentoRepository(c *gorm.DB) *DocumentoRepository {
	return &DocumentoRepository{conn: c}
}
func (r *DocumentoRepository) CreateMemorando(m *models.Memorando) error {
	if err := r.conn.Create(m).Error; err != nil {
		return errores.NewInternalDBf(err)
	}
	return nil
}
func (r *DocumentoRepository) CreateInforme(i *models.Informe) error {
	if err := r.conn.Create(i).Error; err != nil {
		return errores.NewInternalDBf(err)
	}
	return nil
}
func (r *DocumentoRepository) CreateOficio(o *models.Oficio) error {
	if err := r.conn.Create(o).Error; err != nil {
		return errores.NewInternalDBf(err)
	}
	return nil
}
func (r *DocumentoRepository) FindMemorandos() ([]*models.Memorando, error) {
	memorandos := []*models.Memorando{}
	if err := r.conn.Omit("Contenido").Find(&memorandos).Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	return memorandos, nil
}
func (r *DocumentoRepository) FindInformes() ([]*models.Informe, error) {
	informe := []*models.Informe{}
	if err := r.conn.Omit("Contenido").Find(&informe).Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	return informe, nil
}
func (r *DocumentoRepository) FindOficios() ([]*models.Oficio, error) {
	oficios := []*models.Oficio{}
	if err := r.conn.Omit("Contenido").Find(&oficios).Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	return oficios, nil
}
