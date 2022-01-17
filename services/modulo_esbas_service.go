package services

import (
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/repos"
)

type ModuloEsbasService struct {
	repo *repos.ModuloEsbasRepository
}

func NewModuloEsbasService(r *repos.ModuloEsbasRepository) *ModuloEsbasService {
	return &ModuloEsbasService{repo: r}
}
func (s *ModuloEsbasService) GetAll() ([]*models.ModuloEsbas, error) {
	return s.repo.FindAll()
}
func (s *ModuloEsbasService) FindCiclos() ([]*models.Ciclo, error) {
	return s.repo.FindCiclos()
}
