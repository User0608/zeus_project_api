package services

import (
	"github.com/User0608/zeus_project_api/errs"
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/repos"
)

type UsuarioService struct {
	sign       string
	repository *repos.UsuarioRepository
}

func NewUsuarioRepository(r *repos.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repository: r, sign: "service.UsuarioService"}
}

func (s UsuarioService) Login(r models.LogginRequest) (*models.Usuario, error) {
	if err := r.Valid(); err != nil {
		return nil, errs.Wrap(errs.Trc(s.sign, "Logging"), err)
	}
	return s.repository.Login(r)
}
func (s UsuarioService) CreateUser(cu models.PostUsuario) (*models.Usuario, error) {
	if err := cu.Valid(); err != nil {
		return nil, errs.Wrap(errs.Trc(s.sign, "CreateUser"), err)
	}
	return s.repository.Create(cu)
}
