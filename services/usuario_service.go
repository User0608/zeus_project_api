package services

import (
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/repos"
	"github.com/ksaucedo002/kcheck"
)

type UsuarioService struct {
	sign       string
	repository *repos.UsuarioRepository
}

func NewUsuarioRepository(r *repos.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repository: r, sign: "service.UsuarioService"}
}

func (s *UsuarioService) Login(r models.LogginRequest) (*models.Usuario, error) {
	if err := kcheck.Valid(r); err != nil {
		return nil, errores.NewBadRequestf(nil, err.Error())
	}
	return s.repository.Login(r)
}
func (s *UsuarioService) CreateUser(cu models.PostUsuario) (*models.Usuario, error) {
	if err := kcheck.Valid(cu); err != nil {
		return nil, errores.NewBadRequestf(nil, err.Error())
	}
	return s.repository.Create(cu)
}
func (s *UsuarioService) FindAll() ([]models.Usuario, error) {
	return s.repository.Listar()
}
func (s *UsuarioService) FreeUsers() ([]models.Usuario, error) {
	return s.repository.FreeUsers()
}
func (s *UsuarioService) Delete(username string) error {
	if username == "" {
		const message = "el parámetro username no puede estar vacío"
		return errores.NewBadRequestf(nil, message)
	}
	return s.repository.Delete(username)
}
