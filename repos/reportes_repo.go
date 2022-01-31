package repos

import (
	"time"

	"github.com/User0608/zeus_project_api/errores"
	"gorm.io/gorm"
)

type ReportesRepository struct {
	conn *gorm.DB
}

func NewReportesRepository(conn *gorm.DB) *ReportesRepository {
	return &ReportesRepository{conn}
}

func (r *ReportesRepository) Actividades(start, end time.Time) ([]map[string]interface{}, error) {
	actividades := []map[string]interface{}{}
	res := r.conn.Raw("select * from fn_actividades(?,?)",
		start.Format("2006-01-02"),
		end.Format("2006-01-02")).Scan(&actividades)
	if res.Error != nil {
		return nil, errores.NewInternalDBf(res.Error)
	}
	return actividades, nil
}
