package services

import (
	"fmt"
	"time"

	"github.com/User0608/zeus_project_api/documentos"
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"github.com/User0608/zeus_project_api/repos"
	"github.com/ksaucedo002/kcheck"
)

type DocumentoService struct {
	repo *repos.DocumentoRepository
}

func NewDocumentoService(r *repos.DocumentoRepository) *DocumentoService {
	return &DocumentoService{repo: r}
}
func (s *DocumentoService) CreateMemorando(m *models.Memorando) error {
	if err := kcheck.Valid(m); err != nil {
		return errores.NewBadRequestf(nil, err.Error())
	}
	m.Fecha = time.Now()
	pdf := documentos.NewMemorando(*m)
	m.Filepath = fmt.Sprintf("%sC%s.pdf", m.Fecha.Format("2006_01_02__15_04_05"), m.Codigo)
	if err := pdf.PDF(fmt.Sprintf("files/memorandos/%s", m.Filepath)); err != nil {
		return errores.NewInternalf(err, "no se pudo completar la operacion")
	}
	return s.repo.CreateMemorando(m)
}
func (s *DocumentoService) CreateInforme(i *models.Informe) error {
	if err := kcheck.Valid(i); err != nil {
		return errores.NewBadRequestf(nil, err.Error())
	}
	i.Fecha = time.Now()
	pdf := documentos.NewInforme(*i)
	i.Filepath = fmt.Sprintf("%sC%s.pdf", i.Fecha.Format("2006_01_02__15_04_05"), i.Codigo)
	if err := pdf.PDF(fmt.Sprintf("files/informes/%s", i.Filepath)); err != nil {
		return errores.NewInternalf(err, "no se pudo completar la operacion")
	}
	return s.repo.CreateInforme(i)
}
func (s *DocumentoService) CreateOficio(o *models.Oficio) error {
	if err := kcheck.Valid(o); err != nil {
		return errores.NewBadRequestf(nil, err.Error())
	}
	o.Fecha = time.Now()
	pdf := documentos.NewOficio(*o)
	o.Filepath = fmt.Sprintf("%sC%s.pdf", o.Fecha.Format("2006_01_02__15_04_05"), o.Codigo)
	if err := pdf.PDF(fmt.Sprintf("files/oficios/%s", o.Filepath)); err != nil {
		return errores.NewInternalf(err, "no se pudo completar la operacion")
	}
	return s.repo.CreateOficio(o)
}
func (s *DocumentoService) FindMemorandos() ([]*models.Memorando, error) {
	return s.repo.FindMemorandos()
}
func (s *DocumentoService) FindInformes() ([]*models.Informe, error) {
	return s.repo.FindInformes()
}
func (s *DocumentoService) FindOficios() ([]*models.Oficio, error) {
	return s.repo.FindOficios()
}
