package models

type Cronograma struct {
	ID          uint            `json:"id"`
	Nombre      string          `json:"nombre"`
	Descripcion string          `json:"descripcion"`
	CreatedBy   string          `json:"created_by"`
	Programas   []*Programacion `json:"programas,omitempty"`
}
