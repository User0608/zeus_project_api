package services

import (
	"fmt"
	"time"

	"github.com/User0608/zeus_project_api/documentos"
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/repos"
)

type DocumentoService struct {
	repo *repos.DocumentoRepository
}

func NewDocumentoService(r *repos.DocumentoRepository) *DocumentoService {
	return &DocumentoService{repo: r}
}
func (s *DocumentoService) CreateMemorando(m *models.Memorando) error {
	m.Fecha = time.Now()
	pdf := documentos.NewMemorando(*m)
	m.Filepath = fmt.Sprintf("%sC%s.pdf", m.Fecha.Format("2006_01_02__15_04_05"), m.Codigo)
	if err := pdf.PDF(fmt.Sprintf("files/memorandos/%s", m.Filepath)); err != nil {
		return errores.NewInternalf(err, "no se pudo completar la operacion")
	}
	return s.repo.CreateMemorando(m)
}
func (s *DocumentoService) FindMemorandos() ([]*models.Memorando, error) {
	return s.repo.FindMemorandos()
}
