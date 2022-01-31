package services

import (
	"time"

	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/repos"
)

type ReporteService struct {
	repository *repos.ReportesRepository
}

func NewReporteService(r *repos.ReportesRepository) *ReporteService {
	return &ReporteService{repository: r}
}

func (s *ReporteService) ActividadesIntervaloMeses(start, end time.Time) ([]map[string]interface{}, error) {
	if start.After(end) {
		return nil, errores.NewBadRequestf(nil, "intervalo de fechas inválido")
	}
	return s.repository.Actividades(start, end)
}
