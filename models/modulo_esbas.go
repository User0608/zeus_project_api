package models

type ModuloEsbas struct {
	Nombre        string `json:"nombre"`
	Descripcion   string `json:"descripcion"`
	NumeroHoras   uint   `json:"numero_horas"`
	NumeroAlumnos uint   `json:"numero_alumnos"`
	CreateBy      string `json:"-"`
}
