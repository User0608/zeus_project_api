package services

import (
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/repos"
)

type CodigoService struct {
	sign       string
	repository *repos.CodigoRepository
}

func NewCodigoService(r *repos.CodigoRepository) *CodigoService {
	return &CodigoService{repository: r, sign: "service.CodigoService"}
}

func (s *CodigoService) FindAllCodigos() ([]models.CategoriaCodigo, error) {
	return s.repository.FindAllCodigos()
}
