package repos

import (
	"fmt"

	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"gorm.io/gorm"
)

type CodigoRepository struct {
	sign string
	conn *gorm.DB
}

func NewCodigoRepository(g *gorm.DB) *CodigoRepository {
	return &CodigoRepository{
		conn: g,
		sign: "repos.CodigoRepository",
	}
}

func (r *CodigoRepository) FindAllCodigos() ([]models.CategoriaCodigo, error) {
	categorias := []models.CategoriaCodigo{}
	if err := r.conn.Preload("CodigosItems").Find(&categorias).Error; err != nil {
		return nil, errores.NewInternalf(fmt.Errorf("repos.FindAllCodigos:%w", err), errores.ErrDatabaseRequest)
	}
	return categorias, nil
}
