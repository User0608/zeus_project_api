package documentos

import (
	"fmt"
	"strings"

	"github.com/User0608/zeus_project_api/models"
	"github.com/go-pdf/fpdf"
	"github.com/sirupsen/logrus"
)

type Informe struct {
	models.Informe
	FechaString string
	DelCargo    string
	pdf         *fpdf.Fpdf
}

func NewInforme(m models.Informe) *Informe {
	pdf := fpdf.New("P", "in", "A4", "")
	loadFonts(pdf)
	pdf.SetMargins(1, 1, 1)
	pdf.AddPage()
	if err := pdf.Error(); err != nil {
		logrus.Warn(err)
	}
	return &Informe{
		FechaString: fmt.Sprintf("Guadalupe, %d de %s del %d", m.Fecha.Day(), meses[m.Fecha.Month().String()], m.Fecha.Year()),
		Informe: models.Informe{
			Codigo:     fmt.Sprintf("INFORME Nº %s/UBO \"Nstra. Sra. de Gpe.\"-G/C", m.Codigo),
			ParteDel:   m.ParteDel,
			DirigidoAl: m.DirigidoAl,
			Asunto:     m.Asunto,
			Fecha:      m.Fecha,
			Contenido:  m.Contenido,
			CreateBy:   m.CreateBy,
		},
		pdf: pdf,
		///falta dato
	}
}
func (m *Informe) PDF(savePath string) error {
	if err := m.pdf.Error(); err != nil {
		return err
	}
	m.writeTitle()
	m.writeDetalles()
	m.writeBody()
	m.writeATT()
	// addParagraph(m.pdf, "DEL                :hola a todos como estan")
	return m.pdf.OutputFileAndClose(savePath)
}
func (m *Informe) writeTitle() {
	m.pdf.SetFont("liberation", "B", 12)
	m.pdf.MultiCell(0, INxPOINT*16*2.5, m.Codigo, "", "C", false)
}
func (m *Informe) writeDetalles() {
	cellHight := INxPOINT * (TEXT_SIZE) * 2.5
	m.pdf.SetFont("liberation", "R", TEXT_SIZE)
	m.pdf.CellFormat(1, cellHight, "DEL", "0", 0, "L", false, 0, "")
	m.pdf.CellFormat(4, cellHight, fmt.Sprintf(":%s", m.ParteDel), "0", 2, "L", false, 0, "")
	m.pdf.CellFormat(4, INxPOINT*(TEXT_SIZE)*1.5, fmt.Sprintf("  %s", m.DelCargo), "0", 1, "L", false, 0, "")
	m.pdf.Ln(0)
	m.pdf.CellFormat(1, cellHight, "AL", "0", 0, "L", false, 0, "")
	m.pdf.CellFormat(4, cellHight, fmt.Sprintf(": %s", m.DirigidoAl), "0", 2, "L", false, 0, "")
	m.pdf.Ln(0)
	m.pdf.CellFormat(1, cellHight, "ASUNTO", "0", 0, "L", false, 0, "")
	m.pdf.CellFormat(4, cellHight, fmt.Sprintf(": %s", m.Asunto), "0", 2, "L", false, 0, "")
	m.pdf.Ln(0)
	m.pdf.CellFormat(1, cellHight, "FECHA", "0", 0, "L", false, 0, "")
	m.pdf.CellFormat(4, cellHight, fmt.Sprintf(": %s", m.FechaString), "0", 1, "L", false, 0, "")
	line := "___________________________________"
	m.pdf.SetFont("liberation", "B", 24)
	m.pdf.CellFormat(0, cellHight, line, "0", 2, "L", false, 0, "")
	m.pdf.CellFormat(0, CELL_HIGHT, "", "0", 2, "L", false, 0, "")
}
func (m *Informe) writeBody() {
	m.pdf.SetFont("liberation", "R", TEXT_SIZE)
	paragraphs := strings.Split(strings.TrimSpace(m.Contenido), "\n")
	fmt.Println(len(paragraphs))
	for _, p := range paragraphs {
		m.pdf.MultiCell(0, CELL_HIGHT, fmt.Sprintf("   %s", p), "", "J", false)
	}
}
func (m *Informe) writeATT() {
	m.pdf.SetFont("liberation", "R", TEXT_SIZE)
	m.pdf.CellFormat(0, CELL_HIGHT, "", "0", 2, "L", false, 0, "")
	m.pdf.CellFormat(0, CELL_HIGHT, "", "0", 2, "L", false, 0, "")
	m.pdf.MultiCell(0, CELL_HIGHT, "Atentamente,", "", "C", false)
}
