package genpdf

import (
	"strconv"

	"github.com/mikeshimura/goreport"
)

func GenerarPdf() {
	r := goreport.CreateGoReport()
	r.SumWork["amountcum="] = 0.0
	font1 := goreport.FontMap{
		FontName: "liberation",
		FileName: "cmd/fonts/LiberationSans-Regular.ttf",
	}
	fonts := []*goreport.FontMap{&font1}
	r.SetFonts(fonts)
	d := new(S1Detail)
	r.RegisterBand(goreport.Band(*d), goreport.Detail)
	h := new(S1Header)
	r.RegisterBand(goreport.Band(*h), goreport.PageHeader)
	s := new(S1Summary)
	r.RegisterBand(goreport.Band(*s), goreport.Summary)
	r.Records = goreport.ReadTextFile("data.txt", 7)
	r.SetPage("A4", "mm", "L")
	r.SetFooterY(190)
	r.Convert(true)
	r.Execute("simple1.pdf")
}

type S1Detail struct {
}

func (h S1Detail) GetHeight(report goreport.GoReport) float64 {
	return 10
}
func (h S1Detail) Execute(report goreport.GoReport) {
	cols := report.Records[report.DataPos].([]string)
	report.Font("liberation", 12, "")
	y := 2.0
	report.Cell(15, y, cols[0])
	report.Cell(30, y, cols[1])
	report.Cell(60, y, cols[2])
	report.Cell(90, y, cols[3])
	report.Cell(120, y, cols[4])
	report.CellRight(135, y, 25, cols[5])
	report.CellRight(160, y, 20, cols[6])
	amt := ParseFloatNoError(cols[5]) * ParseFloatNoError(cols[6])
	report.SumWork["amountcum="] += amt
	report.CellRight(180, y, 30, strconv.FormatFloat(amt, 'f', 2, 64))
}
func ParseFloatNoError(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

type S1Header struct {
}

func (h S1Header) GetHeight(report goreport.GoReport) float64 {
	return 30
}
func (h S1Header) Execute(report goreport.GoReport) {
	report.Font("liberation", 14, "")
	report.Cell(50, 15, "Sales Report")
	report.Font("liberation", 12, "")
	report.Cell(240, 20, "page")
	report.Cell(260, 20, strconv.Itoa(report.Page))
	y := 23.0
	report.Cell(15, y, "D No")
	report.Cell(30, y, "Dept")
	report.Cell(60, y, "Order")
	report.Cell(90, y, "Stock")
	report.Cell(120, y, "Name")
	report.CellRight(135, y, 25, "Unit Price")
	report.CellRight(160, y, 20, "Qty")
	report.CellRight(190, y, 20, "Amount")
}

type S1Summary struct {
}

func (h S1Summary) GetHeight(report goreport.GoReport) float64 {
	return 10
}
func (h S1Summary) Execute(report goreport.GoReport) {
	report.Cell(160, 2, "Total")
	report.CellRight(180, 2, 30, strconv.FormatFloat(
		report.SumWork["amountcum="], 'f', 2, 64))
}
